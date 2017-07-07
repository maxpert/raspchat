//go:generate msgp -unexported

package rascore

import (
    "time"

    "github.com/tinylib/msgp/msgp"
)

type IMarshalableMessage interface {
    msgp.Marshaler
    msgp.Unmarshaler
    msgp.Sizer
}

type IEventMessage interface {
    IMarshalableMessage
    Identity() uint64
    Event() string
    Stamp()
}

type ICompositeMessage interface {
    IEventMessage
    Message() IEventMessage
    Set(IEventMessage)
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

type compositeMessage struct {
    Base             *BaseMessage                `json:"base,omitempty"`
    Ping             *PingMessage                `json:"ping,omitempty"`
    Handshake        *HandshakeMessage           `json:"handshake,omitempty"`
    Recipient        *RecipientMessage           `json:"recipient,omitempty"`
    RecipientContent *RecipientContentMessage    `json:"recipientcontent,omitempty"`
    Nick             *NickMessage                `json:"nick,omitempty"`
    String           *StringMessage              `json:"str,omitempty"`
    Error            *ErrorMessage               `json:"error,omitempty"`
}

// NewCompositeMessage returns a wrapper serializable message that
// exclusively holds one of child messages in a corresponding wrapped field
func NewCompositeMessage(m IEventMessage) ICompositeMessage {
    r := compositeMessage{}
    r.Set(m)
    return &r
}

func (r *compositeMessage) Identity() uint64 {
    return r.Message().Identity()
}

func (r *compositeMessage) Event() string {
    return r.Message().Event()
}

func (r *compositeMessage) Stamp() {
    r.Message().Stamp()
}

func (r *compositeMessage) Set(m IEventMessage)  {
    r.Base = nil
    r.Ping = nil
    r.Handshake = nil
    r.Recipient = nil
    r.RecipientContent = nil
    r.Nick = nil
    r.String = nil
    r.Error = nil

    if m == nil {
        return
    }

    switch v := m.(type) {
    case *BaseMessage:
        r.Base = v
    case *PingMessage:
        r.Ping = v
    case *HandshakeMessage:
        r.Handshake = v
    case *RecipientMessage:
        r.Recipient = v
    case *RecipientContentMessage:
        r.RecipientContent = v
    case *NickMessage:
        r.Nick = v
    case *StringMessage:
        r.String = v
    case *ErrorMessage:
        r.Error = v
    }
}

func (r *compositeMessage) Message() IEventMessage {
    if r.Base != nil {
        return r.Base
    }

    if r.Ping != nil {
        return r.Ping
    }

    if r.Handshake != nil {
        return r.Handshake
    }

    if r.Recipient != nil {
        return r.Recipient
    }

    if r.RecipientContent != nil {
        return r.RecipientContent
    }

    if r.Nick != nil {
        return r.Nick
    }

    if r.String != nil {
        return r.String
    }

    if r.Error != nil {
        return r.Error
    }

    return nil
}
