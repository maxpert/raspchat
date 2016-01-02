package rica

import (
	"log"
	"net/http"
	"sync"

	"github.com/googollee/go-socket.io"
)

type iClientRelayDirectory interface {
	GetRelay(id string) *Relay
	AddRelay(id string, relay *Relay)
	RemoveRelay(id string)
}

type RelayService struct {
	sync.Mutex
	socketio     *socketio.Server
	relayMap     map[string]*Relay
	groupInfo    GroupInfoManager
	nickRegistry *NickRegistry
}

func NewRelayService(server *socketio.Server, groupInfo GroupInfoManager, nickReg *NickRegistry) *RelayService {
	me := &RelayService{
		socketio:     server,
		relayMap:     make(map[string]*Relay),
		groupInfo:    groupInfo,
		nickRegistry: nickReg,
	}

	server.On("connection", me.onNewConnection)
	return me
}

func (me *RelayService) onNewConnection(so socketio.Socket) {
	sockid := so.Id()
	log.Println("New connection", sockid)
	if oldRelay := me.GetRelay(sockid); oldRelay != nil {
		log.Println("!!!!!!!!!!!!!!!!! Using existing connection ", sockid)
		return
	}

	me.createNewRelay(so)
}

func (me *RelayService) createNewRelay(so socketio.Socket) {
	sockid := so.Id()

	r := NewRelay(so, me, me.groupInfo, me.nickRegistry)
	me.AddRelay(sockid, r)

	so.On("disconnection", func() {
		me.destroyRelay(sockid)
	})

	so.On("error", func(so socketio.Socket, err error) {
		log.Println("Error", err)
		me.destroyRelay(sockid)
	})

	r.Start()
}

func (me *RelayService) destroyRelay(sockid string) {
	if r := me.GetRelay(sockid); r != nil {
		log.Println("Removing connection id", sockid)
		r.Stop()
		delete(me.relayMap, sockid)
		log.Println("Removed connection id", sockid)
	}
}

func (me *RelayService) AddRelay(id string, relay *Relay) {
	me.Lock()
	defer me.Unlock()

	me.relayMap[id] = relay
}

func (me *RelayService) GetRelay(id string) *Relay {
	me.Lock()
	defer me.Unlock()

	if r, ok := me.relayMap[id]; ok {
		return r
	}

	return nil
}

func (me *RelayService) RemoveRelay(id string) {
	me.Lock()
	defer me.Unlock()

	delete(me.relayMap, id)
}

func (me *RelayService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	me.socketio.ServeHTTP(w, req)
}
