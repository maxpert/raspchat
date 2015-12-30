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
	"strings"
	"sync"

	"github.com/googollee/go-socket.io"
)

var message_delimeter string = "~~~~>"
var _FROM_SERVER string = "SERVER"
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
		clientid:  "",
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

		if oldSock, ok := me.relayMap[sockid]; ok {
			log.Println("!!!!!!!!!!!!!!!!! Stopping existing connection ", sockid)
			oldSock.Stop()
			delete(me.relayMap, sockid)
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
	stopWatch := StartStopWatch("Start:" + me.clientid)
	defer stopWatch.LogDuration()

	me.nick = me.sock.Id()
	_nickRegistry.Register(me.clientid, me.nick)

	me.sock.On("init-client", me.onInitClient)
	me.sock.On("send-msg", me.onClientSend)
	me.sock.On("set-nick", me.onClientSetNick)
	me.sock.On("join-group", me.onClientJoin)
	me.sock.On("leave-group", me.onClientLeave)
	log.Println("--------Socket started and hooked!--------")
}

func (me *Relay) Stop() {
	stopWatch := StartStopWatch("Stop:" + me.clientid)
	defer stopWatch.LogDuration()

	if !_nickRegistry.Unregister(me.clientid) {
		log.Println("Unable to unregister client id", me.clientid)
	}

	for _, grp := range me.sock.Rooms() {
		me.groupInfo.RemoveUser(grp, me.clientid)

		me.sock.Leave(grp)
		me.sock.BroadcastTo(grp, "group-leave", me.nick+"@"+grp)
	}
	me.sock = nil
	log.Println("Stopping socket client id", me.clientid)
}

func (me *Relay) onInitClient(m *HandshakeMessage) {
	log.Println("On client initialization", m)
	stopWatch := StartStopWatch("onInitClient:" + me.sock.Id())
	defer stopWatch.LogDuration()

	var err error
	var nickMsg *NickMessage = &NickMessage{OldNick: me.nick}
	if me.nick, err = _nickRegistry.SetNick(me.clientid, m.Nick); err != nil {
		nickMsg.NewNick = me.nick
		defer me.sock.Emit("nick-set", &ErrorMessage{
			Error: err.Error(),
			Body:  nickMsg,
		})
	} else {
		nickMsg.NewNick = me.nick
		defer me.sock.Emit("nick-set", nickMsg)
	}

	defer me.sock.Emit("new-msg", &ChatMessage{
		From:    _FROM_SERVER,
		To:      _FROM_SERVER,
		Message: getWelcomeMessage(),
	})

	me.sock.Emit("client-init")
}

func (me *Relay) onClientSetNick(msg string) {
	log.Println("command.set-nick ---->", msg)
	stopWatch := StartStopWatch("onClientSetNick")
	defer stopWatch.LogDuration()

	var err error
	oldnick := me.nick
	if me.nick, err = _nickRegistry.SetNick(me.clientid, msg); err != nil {
		me.sock.Emit("new-msg", err.Error())
	}

	if oldnick == me.nick {
		return
	}

	nickMsg := &NickMessage{
		OldNick: oldnick,
		NewNick: me.nick,
	}

	for _, name := range me.sock.Rooms() {
		log.Println("Member nick change broadcasting to", name)
		me.sock.BroadcastTo(name, "member-nick-set", name, nickMsg)
	}
	me.sock.Emit("nick-set", nickMsg)
}

func (me *Relay) onClientJoin(groupName string) {
	log.Println("command.join ---->", groupName)
	stopWatch := StartStopWatch("onClientJoin")
	defer stopWatch.LogDuration()

	me.groupInfo.AddUser(groupName, me.clientid, true)

	m := &RecipientMessage{
		To:   groupName,
		From: me.nick,
	}
	me.sock.BroadcastTo(groupName, "group-join", m)
	me.sock.Join(groupName)

	me.sock.Emit("group-join", m)
}

func (me *Relay) onClientLeave(groupName string) {
	log.Println("command.leave ---->", groupName)
	stopWatch := StartStopWatch("onClientLeave")
	defer stopWatch.LogDuration()

	me.groupInfo.RemoveUser(groupName, me.clientid)

	me.sock.Leave(groupName)
	msg := &RecipientMessage{
		From: me.nick,
		To:   groupName,
	}
	me.sock.BroadcastTo(groupName, "group-leave", msg)
	me.sock.Emit("group-leave", msg)
}

func (me *Relay) onClientSend(msg *ChatMessage) {
	log.Println("command.send ----> ", msg)
	stopWatch := StartStopWatch("onClientSend")
	defer stopWatch.LogDuration()

	// If user is not a member of group ignore his message
	if me.groupInfo.GetUserInfoObject(msg.To, me.clientid) == nil {
		log.Println("Message ignored by user", me.clientid, "due to membership")
		return
	}

	msg.From = me.nick
	log.Println("Sending message", msg.Message, "to", msg.To)
	me.sock.BroadcastTo(msg.To, "group-message", msg)
	me.sock.Emit("new-msg", msg)
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
