package rica

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type ChatService struct {
	sync.Mutex
	groupInfo    GroupInfoManager
	nickRegistry *NickRegistry
	upgrader     *websocket.Upgrader
}

func NewChatService() *ChatService {
	return &ChatService{
		groupInfo:    NewInMemoryGroupInfo(),
		nickRegistry: NewNickRegistry(),
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (c *ChatService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if conn, err := c.upgrader.Upgrade(w, req, nil); err == nil {
		log.Println("Socket upgraded!!!")
		handler := NewChatHandler(c.nickRegistry, c.groupInfo, conn)
		go handler.Loop()
	}
}
