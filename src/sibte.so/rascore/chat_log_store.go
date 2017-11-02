package rascore

import (
    "bytes"
    "encoding/binary"
    "fmt"

    "github.com/boltdb/bolt"
)

var _DefaultBucket []byte = []byte("messages")

// ChatLogStore represents abstraction for chat log store
type ChatLogStore struct {
    store       *bolt.DB
    cMaxIDBytes []byte
}

// NewChatLogStore creates a chat log store for passed dataPath
func NewChatLogStore(dataPath string) (*ChatLogStore, error) {
    db, err := bolt.Open(dataPath, 0660, nil)
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

    tnx, err := c.store.Begin(true)
    if err != nil {
        return err
    }

    buck, err := tnx.CreateBucketIfNotExists(_DefaultBucket)
    if err != nil {
        return err
    }

    // <group-name><id> -> <msg>
    // <id> -> <group-name>
    // <group-name><MAX_ID> -> byte[0]
    buck.Put(append([]byte(group), bytesID...), bytesMsg)
    buck.Put(bytesID, []byte(group))
    buck.Put(append([]byte(group), maxIDBytes...), make([]byte, 0))

    return tnx.Commit()
}

// GetMessagesFor returns messages for given group starting at start_id
// result-set shape is governed by offset and limit passed
func (c *ChatLogStore) GetMessagesFor(group, startID string, offset, limit uint) ([]IEventMessage, error) {
    var ret []IEventMessage

    tnx, err := c.store.Begin(false)
    if err != nil {
        return nil, err
    }

    defer tnx.Rollback()
    buck := tnx.Bucket(_DefaultBucket)
    if buck == nil {
        return make([]IEventMessage, 0), nil
    }

    csr := buck.Cursor()
    if csr == nil {
        return ret, nil
    }

    maxIDBytes := idToBytes(^uint64(0))
    endBytesID := append([]byte(group), maxIDBytes...)
    if startID != "" {
        endBytesID = []byte(startID)
    }

    i := uint(0)

    k, v := csr.Seek(endBytesID)
    if k == nil || !bytes.Equal(endBytesID, k) {
        return make([]IEventMessage, 0), nil
    }

    for true {
        k, v = csr.Prev()
        i++

        if k == nil || bytes.HasPrefix(k, []byte(group)) == false {
            break
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
    tnx, err := c.store.Begin(false)
    if err != nil {
        return nil, err
    }

    defer tnx.Rollback()

    buck := tnx.Bucket(_DefaultBucket)
    if buck == nil {
        return nil, nil
    }

    val := buck.Get(idToBytes(id))
    if val == nil {
        return nil, nil
    }

    // Create copy of array since we should not modify the values returned
    groupName := append([]byte(nil), val...)
    val = buck.Get(append(groupName, idToBytes(id)...))
    if val != nil {
        return nil, fmt.Errorf("Unable to group message entry for id %v", id)
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
