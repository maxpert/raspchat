package rascore

import (
    "errors"
    "sync"

    "sibte.so/rascore/consts"

    "github.com/gorilla/websocket"
)

// WebsocketMessageTransport go routine safe transport connection
type WebsocketMessageTransport struct {
    connection          *websocket.Conn
    connectionReadLock  *sync.Mutex
    connectionWriteLock *sync.Mutex
}

// NewWebsocketMessageTransport creates a new websocket connection transport
func NewWebsocketMessageTransport(conn *websocket.Conn) *WebsocketMessageTransport {
    return &WebsocketMessageTransport{
        connection:          conn,
        connectionReadLock:  &sync.Mutex{},
        connectionWriteLock: &sync.Mutex{},
    }
}

// ReadMessage from transport
func (h *WebsocketMessageTransport) ReadMessage() ([]byte, error) {
    h.connectionReadLock.Lock()
    msgType, msg, err := h.connection.ReadMessage()
    h.connectionReadLock.Unlock()

    if err != nil {
        return nil, err
    }

    if msgType != websocket.TextMessage {
        return nil, errors.New(rasconsts.ERROR_INVALID_MSGTYPE_ERR)
    }

    return msg, nil
}

// WriteMessage to transport
func (h *WebsocketMessageTransport) WriteMessage(id uint64, msg []byte) error {
    return h.writeMessageOnSocket(msg)
}

func (h *WebsocketMessageTransport) writeMessageOnSocket(msg []byte) error {
    h.connectionWriteLock.Lock()
    defer h.connectionWriteLock.Unlock()
    return h.connection.WriteMessage(websocket.TextMessage, msg)
}

// FlushBatch of messages
func (h *WebsocketMessageTransport) FlushBatch(id uint64) {
}

// BeginBatch of messages to be written on transport
func (h *WebsocketMessageTransport) BeginBatch(id uint64) {
    // Ignore since it's not buffered
}
