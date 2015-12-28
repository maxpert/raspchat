package rica

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/googollee/go-socket.io"
)

var message_delimeter string = "~~~~>"
var from_server_prefix string = "SERVER" + message_delimeter

type Relay struct {
	sync.Mutex
	sock      socketio.Socket
	clientid  string
	nick      string
	groupInfo GroupInfoManager
}

type RelayService struct {
	sync.Mutex
	socketio  *socketio.Server
	relayMap  map[string]*Relay
	groupInfo GroupInfoManager
}

var _nickRegistry *NickRegistry = NewNickRegistry()

func NewRelay(sock socketio.Socket, infoMan GroupInfoManager) *Relay {
	log.Println("Creating new relay server")
	return &Relay{
		sock:      sock,
		clientid:  sock.Id(),
		nick:      sock.Id(),
		groupInfo: infoMan,
	}
}

func NewRelayService(server *socketio.Server) *RelayService {
	me := &RelayService{
		socketio:  server,
		relayMap:  make(map[string]*Relay),
		groupInfo: NewInMemoryGroupInfo(),
	}

	server.On("connection", func(so socketio.Socket) {
		sockid := so.Id()
		log.Println("New connection", sockid)

		if _, ok := me.relayMap[sockid]; ok {
			log.Println("Using existing connection", sockid)
			return
		}

		me.createNewRelay(so)
	})

	return me
}

func (me *RelayService) createNewRelay(so socketio.Socket) {
	sockid := so.Id()

	me.Lock()
	r, ok := me.relayMap[sockid]
	if !ok {
		r = NewRelay(so, me.groupInfo)
		me.relayMap[sockid] = r
	}
	me.Unlock()

	so.On("disconnection", func() {
		me.destroyRelay(sockid)
	})

	so.On("error", func(so socketio.Socket, err error) {
		log.Println("Error", err)
		me.destroyRelay(so.Id())
	})

	r.Start()
}

func (me *RelayService) destroyRelay(sockid string) {
	r, ok := me.relayMap[sockid]
	if ok {
		go r.Stop()
		log.Println("Removing connection id", sockid)
		delete(me.relayMap, sockid)
	}
}

func (me *RelayService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	me.socketio.ServeHTTP(w, req)
}

func (me *Relay) Start() {
	me.nick = me.sock.Id()
	_nickRegistry.Register(me.clientid, me.nick)

	me.sock.On("send-msg", me.onClientSend)
	me.sock.On("set-nick", me.onClientSetNick)
	me.sock.On("join-group", me.onClientJoin)
	me.sock.On("leave-group", me.onClientLeave)
	me.sock.On("ping", me.onPing)
	me.sock.Emit("new-msg", from_server_prefix+getWelcomeMessage())
}

func (me *Relay) Stop() {
	_nickRegistry.Unregister(me.clientid)

	for _, grp := range me.sock.Rooms() {
		me.groupInfo.RemoveUser(grp, me.clientid)

		me.sock.Leave(grp)
		me.sock.BroadcastTo(grp, "group-leave", me.nick+"@"+grp)
	}
	me.sock = nil
	log.Println("Stopping socket client id", me.clientid)
}

func (me *Relay) onClientSetNick(msg string) {
	invalidAliasRegex, _ := regexp.Compile("[^A-Za-z0-9]")
	if invalidAliasRegex.MatchString(msg) || len(msg) > 42 {
		me.sock.Emit("new-msg", from_server_prefix+"A nick can only have alpha-numeric values")
		return
	}

	if _nickRegistry.Register(me.clientid, msg) == false {
		me.sock.Emit("new-msg", from_server_prefix+"Nick already registered")
		return
	}

	oldnick := me.nick
	me.nick = msg

	for _, name := range me.sock.Rooms() {
		me.sock.BroadcastTo(name, "member-nick-changed", oldnick+"->"+me.nick)
	}

	me.sock.Emit("nick-set", me.nick)
}

func (me *Relay) onClientJoin(msg string) {
	log.Println("command.join ---->", msg)
	me.groupInfo.AddUser(msg, me.clientid, true)

	me.sock.BroadcastTo(msg, "group-join", me.nick+"@"+msg)
	me.sock.Join(msg)

	me.sock.Emit("group-join", me.nick+"@"+msg)
}

func (me *Relay) onClientLeave(msg string) {
	log.Println("command.leave ---->", msg)
	me.groupInfo.RemoveUser(msg, me.clientid)

	me.sock.Leave(msg)
	me.sock.BroadcastTo(msg, "group-leave", me.nick+"@"+msg)

	me.sock.Emit("group-leave", me.nick+"@"+msg)
}

func (me *Relay) onPing() {
	me.sock.Emit("pong")
}

func (me *Relay) onClientSend(msg string) {
	log.Println("command.send ----> ", msg)

	// Split message [channel]~~~~>[msg]
	info := strings.Split(msg, message_delimeter)
	if len(info) < 2 {
		return
	}

	message := me.nick + "@" + info[0] + message_delimeter + info[1]
	log.Println("Sending message", msg, "to", info[0])
	me.sock.BroadcastTo(info[0], "group-message", message)
	me.sock.Emit("new-msg", message)
}

func getWelcomeMessage() string {
	info, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "Unable to query cpu info"
	}

	return fmt.Sprintf(
		"CPU INFO: \n --- \n %v \n",
		strings.Replace(string(info), "\n", "\n\n", 999),
	)
}

func toPrettyJson(obj interface{}) string {
	v_json, err := json.Marshal(obj)
	if err != nil {
		return "{'err': 'Unable to serialze'}"
	}

	var pretty_json bytes.Buffer
	json.Indent(&pretty_json, v_json, "", "  ")
	return string(pretty_json.Bytes())
}
