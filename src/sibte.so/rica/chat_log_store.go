package rica

import (
    "bytes"
    "encoding/binary"
    "encoding/gob"
    "errors"
    "fmt"

    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
)

type ChatLogStore struct {
    store       *leveldb.DB
    cMaxIDBytes []byte
}

func NewChatLogStore(path string) (*ChatLogStore, error) {
    db, err := leveldb.OpenFile(path, nil)
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

func (c *ChatLogStore) Save(group string, id uint64, msg IEventMessage) error {
    bytesMsg := c.serialize(msg)

    if bytesMsg == nil {
        return errors.New("Unable to serialize msg")
    }

    bytesId := idToBytes(id)
    maxIdBytes := c.cMaxIDBytes

    // <group-name><id> -> <msg>
    // <id> -> <group-name>
    // <group-name><MAXID> -> byte[0]
    b := &leveldb.Batch{}
    b.Put(append([]byte(group), bytesId...), bytesMsg)
    b.Put(bytesId, []byte(group))
    b.Put(append([]byte(group), maxIdBytes...), make([]byte, 0))
    return c.store.Write(b, &opt.WriteOptions{
        Sync: false,
    })
}

func (c *ChatLogStore) GetMessagesFor(group string, start_id string, offset uint, limit uint) ([]IEventMessage, error) {
    var ret []IEventMessage

    csr := c.store.NewIterator(nil, nil)
    if csr == nil {
        return ret, nil
    }

    maxIDBytes := idToBytes(^uint64(0))
    endBytesID := append([]byte(group), maxIDBytes...)
    if start_id != "" {
        endBytesID = []byte(start_id)
    }

    i := uint(0)
    for csr.Seek(endBytesID); true; csr.Prev() {
        // Make sure we don't modify k & v
        k := csr.Key()
        v := csr.Value()
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

func (c *ChatLogStore) GetMessage(id uint64) (IEventMessage, error) {
    group, err := c.store.Get(idToBytes(id), nil)
    if err != nil {
        return nil, err
    }

    if group == nil {
        return nil, nil
    }

    // Create copy of array since we should not modify the values returned
    group = append([]byte(nil), group...)

    bytesMsg, err := c.store.Get(append(group, idToBytes(id)...), nil)
    if err != nil {
        return nil, err
    }

    if bytesMsg == nil {
        return nil, errors.New("Unable to locate message value")
    }

    m := c.deserialize(bytesMsg)
    if m == nil {
        return nil, errors.New(fmt.Sprintf("Unable to deserialize message %v %v", group, id))
    }

    return m, nil
}

func (c *ChatLogStore) Cleanup(group string) {
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

    stM := &StringMessage{}
    if dec.Decode(stM) == nil {
        return stM
    }

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
