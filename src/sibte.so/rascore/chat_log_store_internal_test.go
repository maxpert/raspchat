package rascore

import (
    "testing"
    "os"
)

var testPath = os.TempDir()

func TestChatLogStoreOpenDatabase(t *testing.T) {
    logStore, err := NewChatLogStore(testPath)
    defer logStore.Close()
    if err != nil {
        t.Error(err)
    }
}

func TestChatLogStoreSaveMessage(t *testing.T) {
    logStore, _ := NewChatLogStore(testPath)
    defer logStore.Close()

    err := logStore.Save("foo", 1, &PingMessage{
        Type: 0,
    })

    if err != nil {
        t.Error(err)
    }
}

func TestChatLogStoreGetMessage(t *testing.T) {
    logStore, _ := NewChatLogStore(testPath)
    defer logStore.Close()

    origMsg := RecipientMessage{
        To: "foo@bar.com",
        From: "bar@foo.com",
    }
    origMsg.Id = 1
    origMsg.EventName = "sample"
    origMsg.UTCTimestamp = 1
    logStore.Save("foo", 1, &origMsg)

    msg, err := logStore.GetMessage(1)

    if err != nil {
        t.Error(err)
    }

    if msg == nil {
        t.Error("Invalid message", msg)
    }

    if _, ok := msg.(*RecipientMessage); !ok {
        t.Error("Invalid message type deserialized")
    }
}

func BenchmarkChatLogStoreDeserialize(b *testing.B) {
    origMsg := RecipientMessage{
        To: "foo@bar.com",
        From: "bar@foo.com",
    }
    origMsg.Id = 1
    origMsg.EventName = "sample"
    origMsg.UTCTimestamp = 1
    serializedMsg, _ := serializeMessage(&origMsg)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        deserializeMessage(serializedMsg)
    }
}
