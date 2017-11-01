package rascore

// IMessageTransport the message transport interface
type IMessageTransport interface {
    WriteMessage(id uint64, message []byte) error
    ReadMessage() ([]byte, error)
    BeginBatch(id uint64)
    FlushBatch(id uint64)
}
