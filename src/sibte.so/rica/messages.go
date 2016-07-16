package rica

import "time"

type IEventMessage interface {
	Identity() uint64
	Event() string
	Stamp()
}

type BaseMessage struct {
	EventName    string `json:"@"`
	Id           uint64 `json:"!id,omitempty"`
	UTCTimestamp int64  `json:"utc_timestamp"`
}

func (b *BaseMessage) Identity() uint64 {
	return b.Id
}

func (b *BaseMessage) Event() string {
	return b.EventName
}

func (b *BaseMessage) Stamp() {
	b.UTCTimestamp = time.Now().Unix()
}

type PingMessage struct {
	BaseMessage
	Type int `json:"t"`
}

type HandshakeMessage struct {
	BaseMessage
	Nick  string   `json:"nick"`
	Rooms []string `json:"rooms"`
}

type RecipientMessage struct {
	BaseMessage
	To   string `json:"to"`
	From string `json:"from"`
}

type ChatMessage struct {
	RecipientMessage
	Message string `json:"msg"`
}

type RecipientContentMessage struct {
	RecipientMessage
	Message interface{} `json:"pack_msg"`
}

type NickMessage struct {
	BaseMessage
	OldNick string `json:"oldNick"`
	NewNick string `json:"newNick"`
}

type StringMessage struct {
	BaseMessage
	Message string `json:"msg"`
}

type ErrorMessage struct {
	BaseMessage
	Type  string      `json:"error_type"`
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}
