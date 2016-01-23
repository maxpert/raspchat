package rica

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type ChatLogStore struct {
	store  *bolt.DB
	bucket []byte
}

func NewChatLogStore(path string, bucket []byte) (*ChatLogStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &ChatLogStore{
		store:  db,
		bucket: bucket,
	}, nil
}

func (c *ChatLogStore) Save(group string, id uint64, msg IEventMessage) error {
	bytesMsg := c.serialize(msg)

	if bytesMsg == nil {
		return errors.New("Unable to serialize msg")
	}

	tx, err := c.store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, err := tx.CreateBucketIfNotExists(c.bucket)
	if err != nil {
		return err
	}

	bytesId := c.idToBytes(id)
	maxIdBytes := c.idToBytes(^uint64(0))

	// <group-name><id> -> <msg>
	// <id> -> <group-name>
	// <group-name><MAXID> -> byte[0]
	b.Put(append([]byte(group), bytesId...), bytesMsg)
	b.Put(bytesId, []byte(group))
	b.Put(append([]byte(group), maxIdBytes...), make([]byte, 0))
	tx.Commit()

	return nil
}

func (c *ChatLogStore) GetMessagesFor(group string, offset uint, limit uint) ([]IEventMessage, error) {
	tx, err := c.store.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	ret := make([]IEventMessage, 0)
	bkt := tx.Bucket(c.bucket)

	if bkt == nil {
		return ret, nil
	}

	csr := bkt.Cursor()
	if csr == nil {
		return ret, nil
	}

	maxIdBytes := c.idToBytes(^uint64(0))
	endBytesId := append([]byte(group), maxIdBytes...)
	i := uint(0)
	for k, v := csr.Seek(endBytesId); true; k, v = csr.Prev() {
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

	log.Println("Chat log store conversations ->", len(ret))
	return ret, nil
}

func (c *ChatLogStore) GetMessage(id uint64) (IEventMessage, error) {
	tx, err := c.store.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	b := tx.Bucket(c.bucket)
	if b == nil {
		return nil, nil
	}

	group := b.Get(c.idToBytes(id))

	if group == nil {
		return nil, nil
	}

	bytesMsg := b.Get(append(group, c.idToBytes(id)...))
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

func (c *ChatLogStore) idToBytes(id uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	return b
}
