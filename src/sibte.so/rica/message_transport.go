package rica

type IMessageTransport interface {
    WriteMessage(id uint64, message IEventMessage) error
    ReadMessage() (IEventMessage, error)
    BeginBatch(id uint64, message IEventMessage)
    FlushBatch(id uint64)
}
