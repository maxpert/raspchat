package rica

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type ChatService struct {
	sync.Mutex
	groupInfo    GroupInfoManager
	chatStore    *ChatLogStore
	nickRegistry *NickRegistry
	upgrader     *websocket.Upgrader
	gcmWorker    *GCMWorker
	httpMux      *http.ServeMux
}

func NewChatService() *ChatService {
	initChatHandlerTypes()
	initGifCache()
	store, e := NewChatLogStore(CurrentAppConfig.DBPath+"/chats.bolt.db", []byte("chats"))

	if e != nil {
		log.Panic(e)
	}

	return &ChatService{
		groupInfo:    NewInMemoryGroupInfo(),
		nickRegistry: NewNickRegistry(),
		gcmWorker:    NewGCMWorker(CurrentAppConfig.GCMToken),
		chatStore:    store,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// CheckOrigin:     func(_ *http.Request) bool { return true },
		},
	}
}

func (c *ChatService) WithRESTRoutes(prefix string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(prefix+"/api/", c.httpRoutes(prefix+"/api", httprouter.New()))
	c.httpMux = mux
	return c
}

func (c *ChatService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling", req.URL.Path)
	if strings.HasPrefix(req.URL.Path, "/chat/api") {
		log.Println("Handling HTTP API call")
		c.httpMux.ServeHTTP(w, req)
		return
	}

	if c.upgradeConnectionToWebSocket(w, req) {
		log.Println("Connection upgraded to WS")
	}
}

func (c *ChatService) httpRoutes(prefix string, router *httprouter.Router) http.Handler {
	router.POST(prefix+"/push", c.onPushPost)
	router.POST(prefix+"/register", c.onPushSubscribe)

	router.GET(prefix+"/channel/:id/message", c.onGetChatHistory)
	router.GET(prefix+"/channel/:id/message/:msg_id", c.onGetChatMessage)
	router.GET(prefix+"/channel", c.onGetChannels)
	return router
}

func (c *ChatService) upgradeConnectionToWebSocket(w http.ResponseWriter, req *http.Request) bool {
	log.Println("New websocket connection request")
	conn, err := c.upgrader.Upgrade(w, req, nil)
	if err == nil {
		log.Println("Socket upgraded!!!")
		transporter := NewWebsocketMessageTransport(conn)
		handler := NewChatHandler(c.nickRegistry, c.groupInfo, transporter, c.chatStore)
		go handler.Loop()
		return true
	}

	log.Println("Error upgrading connection...", err)
	return false
}

func (c *ChatService) onPushSubscribe(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	token := req.FormValue("gcm_sub_token")
	if token == "" {
		fmt.Fprintf(w, "false")
		return
	}

	transporter := NewGCMTransport(token, c.gcmWorker)
	handler := NewChatHandler(c.nickRegistry, c.groupInfo, transporter, c.chatStore)
	go handler.Loop()
	fmt.Fprintf(w, "true")
}

func (c *ChatService) onPushPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	token := req.FormValue("gcm_sub_token")
	log.Println("Received new POST message request...", req.Method, " => ", token != "")
	t := NewGCMTransport(token, c.gcmWorker)
	if msg, err := ioutil.ReadAll(req.Body); req.Method == "POST" && err == nil {
		t.PostMessage(string(msg))
		fmt.Fprintf(w, "true")
		return
	}

	fmt.Fprintf(w, "false")
}

func (c *ChatService) onGetChatHistory(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	log.Println("Get chat history...")
	groupId := p.ByName("id")

	queryParams := req.URL.Query()
	var offset uint = 0
	var limit uint = 20

	if o, err := strconv.ParseUint(queryParams.Get("offset"), 10, 32); err == nil {
		offset = uint(o)
	}

	if l, err := strconv.ParseUint(queryParams.Get("limit"), 10, 32); err == nil {
		limit = uint(l)
	}

	log.Println("Limit =", limit, "Offset =", offset)
	log, err := c.chatStore.GetMessagesFor(groupId, offset, limit)
	if err == nil {
		response := make(map[string]interface{})
		response["limit"] = limit
		response["offset"] = offset
		response["messages"] = log
		response["id"] = groupId
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorMessage{
			Error: err.Error(),
		})
	}
}

func (c *ChatService) onGetChatMessage(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
}

func (c *ChatService) onGetChannels(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}
