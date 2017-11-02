package rascore

import (
    "bytes"
    "encoding/binary"
    "errors"
    "path"

    "github.com/dgraph-io/badger"

    "sibte.so/rascore/utils"
)

// ChatLogStore represents abstraction for chat log store
type ChatLogStore struct {
    store       *badger.DB
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

    db, err := badger.Open(opts)
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
    bytesMsg, err := serializeMessage(msg)

    if err != nil {
        return err
    }

    bytesID := idToBytes(id)
    maxIDBytes := c.cMaxIDBytes

    // <group-name><id> -> <msg>
    // <id> -> <group-name>
    // <group-name><MAX_ID> -> byte[0]
    tnx := c.store.NewTransaction(true)
    tnx.Set(append([]byte(group), bytesID...), bytesMsg)
    tnx.Set(bytesID, []byte(group))
    tnx.Set(append([]byte(group), maxIDBytes...), make([]byte, 0))

    return tnx.Commit(nil)
}

// GetMessagesFor returns messages for given group starting at start_id
// result-set shape is governed by offset and limit passed
func (c *ChatLogStore) GetMessagesFor(group, startID string, offset, limit uint) ([]IEventMessage, error) {
    var ret []IEventMessage

    opts := badger.DefaultIteratorOptions
    opts.Reverse = true
    tnx := c.store.NewTransaction(false)
    defer tnx.Discard()

    csr := tnx.NewIterator(opts)
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
        v, err := tuple.Value()
        if err != nil {
            return nil, err
        }

        i++

        if k == nil || bytes.HasPrefix(k, []byte(group)) == false {
            break
        }

        if i < offset || bytes.Equal(endBytesID, k) {
            continue
        }

        if i > limit {
            break
        }

        msg, err := deserializeMessage(v)
        if err != nil {
            return nil, err
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
    tnx := c.store.NewTransaction(false)
    defer tnx.Discard()

    kvItem, err := tnx.Get(idToBytes(id))
    if err != nil {
        return nil, err
    }

    val, err := kvItem.Value()
    if err == nil {
        return nil, nil
    }

    // Create copy of array since we should not modify the values returned
    groupName := append([]byte(nil), val...)
    kvItem, err = tnx.Get(append(groupName, idToBytes(id)...))
    if err != nil {
        return nil, err
    }

    val, err = kvItem.Value()
    if err == nil {
        return nil, errors.New("Unable to locate message value")
    }

    m, err := deserializeMessage(val)
    if err != nil {
        return nil, err
    }

    return m, nil
}

// Close and flush store values
func (c *ChatLogStore) Close() error {
    return c.store.Close()
}

func serializeMessage(valueMsg IEventMessage) ([]byte, error) {
    msg := NewCompositeMessage(valueMsg)
    buffer := make([]byte, 0)
    return msg.MarshalMsg(buffer)
}

func deserializeMessage(bytes []byte) (IEventMessage, error) {
    msg := NewCompositeMessage(nil)
    if _, err := msg.UnmarshalMsg(bytes); err != nil {
        return nil, err
    }

    return msg.Message(), nil
}
