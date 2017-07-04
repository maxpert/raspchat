package rascore

import (
    "bytes"
    "encoding/binary"
    "encoding/gob"
    "errors"
    "fmt"
    "path"

    "github.com/maxpert/badger"

    "sibte.so/rascore/utils"
)

// ChatLogStore represents abstraction for chat log store
type ChatLogStore struct {
    store       *badger.KV
    cMaxIDBytes []byte
}

// NewChatLogStore creates a chat log store for passed dataPath
func NewChatLogStore(dataPath string) (*ChatLogStore, error) {
    opts := badger.DefaultOptions
    opts.Dir = path.Join(dataPath, "keys")
    opts.ValueDir = path.Join(dataPath, "values")
    opts.SyncWrites = true // Messages are something we don't want to loose

    if err := rasutils.CreatePathIfMissing(opts.Dir); err != nil {
        return nil, err
    }

    if err := rasutils.CreatePathIfMissing(opts.ValueDir); err != nil {
        return nil, err
    }

    db, err := badger.NewKV(&opts)
    if err != nil {
        return nil, err
    }

    return &ChatLogStore{
        store:       db,
        cMaxIDBytes: idToBytes(^uint64(0)),
    }, nil
}

func idToBytes(id uint64) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, id)
    return b
}

// Save saves a given message to given group with given id
func (c *ChatLogStore) Save(group string, id uint64, msg IEventMessage) error {
    bytesMsg := c.serialize(msg)

    if bytesMsg == nil {
        return errors.New("Unable to serialize msg")
    }

    bytesID := idToBytes(id)
    maxIDBytes := c.cMaxIDBytes

    // <group-name><id> -> <msg>
    // <id> -> <group-name>
    // <group-name><MAX_ID> -> byte[0]
    entries := make([]*badger.Entry, 3)
    entries[0] = &badger.Entry{
        Key:   append([]byte(group), bytesID...),
        Value: bytesMsg,
    }

    entries[1] = &badger.Entry{
        Key:   bytesID,
        Value: []byte(group),
    }

    entries[2] = &badger.Entry{
        Key:   append([]byte(group), maxIDBytes...),
        Value: make([]byte, 0),
    }

    return c.store.BatchSet(entries)
}

// GetMessagesFor returns messages for given group starting at start_id
// Resultset is governed by offset and limit passed
func (c *ChatLogStore) GetMessagesFor(group, startID string, offset, limit uint) ([]IEventMessage, error) {
    var ret []IEventMessage

    opts := badger.DefaultIteratorOptions
    opts.FetchValues = true
    opts.Reverse = true
    csr := c.store.NewIterator(opts)
    if csr == nil {
        return ret, nil
    }

    defer csr.Close()

    maxIDBytes := idToBytes(^uint64(0))
    endBytesID := append([]byte(group), maxIDBytes...)
    if startID != "" {
        endBytesID = []byte(startID)
    }

    i := uint(0)
    for csr.Seek(endBytesID); csr.Valid(); csr.Next() {
        tuple := csr.Item()
        k := tuple.Key()
        v := tuple.Value()
        i++

        if k == nil || bytes.HasPrefix(k, []byte(group)) == false {
            break
        }

        if i < offset {
            continue
        }

        if i > limit {
            break
        }

        msg := c.deserialize(v)
        if msg == nil {
            continue
        }

        ret = append(ret, msg)
    }

    return ret, nil
}

// GetMessage returns message for given id
func (c *ChatLogStore) GetMessage(id uint64) (IEventMessage, error) {
    // <group-name><id> -> <msg>
    // <id> -> <group-name>
    // <group-name><MAX_ID> -> byte[0]
    kvItem := badger.KVItem{}
    err := c.store.Get(idToBytes(id), &kvItem)
    if err != nil {
        return nil, err
    }

    if kvItem.Value() == nil {
        return nil, nil
    }

    // Create copy of array since we should not modify the values returned
    groupName := append([]byte(nil), kvItem.Value()...)
    kvItem = badger.KVItem{}
    err = c.store.Get(append(groupName, idToBytes(id)...), &kvItem)
    if err != nil {
        return nil, err
    }

    if kvItem.Value() == nil {
        return nil, errors.New("Unable to locate message value")
    }

    m := c.deserialize(kvItem.Value())
    if m == nil {
        return nil, fmt.Errorf("Unable to deserialize message %v", id)
    }

    return m, nil
}

func (c *ChatLogStore) serialize(v IEventMessage) []byte {
    var buffer bytes.Buffer
    enc := gob.NewEncoder(&buffer)

    if enc.Encode(v) != nil {
        return nil
    }

    return buffer.Bytes()
}

func (c *ChatLogStore) deserialize(b []byte) IEventMessage {
    buffer := bytes.NewBuffer(b)
    dec := gob.NewDecoder(buffer)

    chM := &ChatMessage{}
    if dec.Decode(chM) == nil {
        return chM
    }

    rpCM := &RecipientContentMessage{}
    if dec.Decode(rpCM) == nil {
        return rpCM
    }

    rpM := &RecipientMessage{}
    if dec.Decode(rpM) == nil {
        return rpM
    }

    var intr IEventMessage
    if err := dec.Decode(intr); err == nil {
        return intr
    }

    return nil
}
