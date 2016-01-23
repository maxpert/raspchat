package rica

import (
	"errors"
	"log"
	"sync"

	"sibte.so/rica/consts"

	"github.com/gorilla/websocket"
)

type WebsocketMessageTransport struct {
	connection          *websocket.Conn
	connectionReadLock  *sync.Mutex
	connectionWriteLock *sync.Mutex
}

func NewWebsocketMessageTransport(conn *websocket.Conn) *WebsocketMessageTransport {
	return &WebsocketMessageTransport{
		connection:          conn,
		connectionReadLock:  &sync.Mutex{},
		connectionWriteLock: &sync.Mutex{},
	}
}

func (h *WebsocketMessageTransport) ReadMessage() (IEventMessage, error) {
	h.connectionReadLock.Lock()
	msgType, msg, err := h.connection.ReadMessage()
	h.connectionReadLock.Unlock()

	log.Println("Message Type", msgType, "Message", string(msg), "Error", err)

	if err != nil {
		return nil, err
	}

	if msgType != websocket.TextMessage {
		return nil, errors.New(ricaEvents.ERROR_INVALID_MSGTYPE_ERR)
	}

	if jsonMsg, e := pDecodeMessage(msg); e == nil {
		return jsonMsg, nil
	} else {
		log.Println("Invalid message", e)
	}

	return nil, err
}

func (h *WebsocketMessageTransport) WriteMessage(id uint64, msg IEventMessage) error {
	return h.writeMessageOnSocket(msg)
}

func (h *WebsocketMessageTransport) writeMessageOnSocket(msg IEventMessage) error {
	h.connectionWriteLock.Lock()
	defer h.connectionWriteLock.Unlock()
	return h.connection.WriteJSON(msg)
}

func (h *WebsocketMessageTransport) FlushBatch(id uint64) {
	log.Println("Websocket published message id", id)
}

func (h *WebsocketMessageTransport) BeginBatch(id uint64, msg IEventMessage) {
	log.Println("Websocket published message id", id)
}
