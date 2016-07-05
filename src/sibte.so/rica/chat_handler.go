package rica

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"io/ioutil"
	"log"
	"math/rand"
	"sibte.so/rica/consts"
	"strings"
	"sync"
	"time"

	"github.com/speps/go-hashids"
)

type ChatHandler struct {
	sync.Mutex
	id               string
	nick             string
	groupInfoManager GroupInfoManager
	nickRegistry     *NickRegistry
	transport        IMessageTransport
	incoming         chan interface{}
	groups           map[string]interface{}
	chatStore        *ChatLogStore
}

var pHashId *hashids.HashID = hashids.New()
var pSnowFlake *SnowFlake = DefaultSnowFlake()

func NewChatHandler(nickReg *NickRegistry, groupInfoMan GroupInfoManager, trans IMessageTransport, store *ChatLogStore) *ChatHandler {
	log.Println("New chat handler")
	uid, _ := pHashId.Encode([]int{
		int(rand.Int31n(1000)),
		int(rand.Int31n(1000)),
		int(rand.Int31n(1000)),
	})

	ret := &ChatHandler{
		id:               uid,
		nick:             uid,
		nickRegistry:     nickReg,
		groupInfoManager: groupInfoMan,
		transport:        trans,
		chatStore:        store,
		incoming:         make(chan interface{}, 32),
		groups:           make(map[string]interface{}, 0),
	}

	ret.groups[ricaEvents.FROM_SERVER] = struct{}{}
	ret.groupInfoManager.AddUser(ricaEvents.FROM_SERVER, ret.id, ret.incoming)
	return ret
}

func messageOf(event string) BaseMessage {
	id, err := pSnowFlake.Next()

	// If we get an error delay for 1ms and try again
	if err != nil {
		time.Sleep(1 * time.Millisecond)
		id, _ = pSnowFlake.Next()
	}

	return BaseMessage{
		EventName: event,
		Id:        id,
	}
}

func (h *ChatHandler) socketLoop(socketChannel chan interface{}, errorChannel chan error) {
	for {
		msg, err := h.transport.ReadMessage()

		// If message type was invalid
		if err != nil && err.Error() == ricaEvents.ERROR_INVALID_MSGTYPE_ERR {
			log.Println("Skipping message....")
			continue
		}

		if err != nil {
			errorChannel <- err
			break
		}

		socketChannel <- msg
	}
}

func (h *ChatHandler) recoverFromErrors() {
	if r := recover(); r != nil {
		log.Println("!!!PANIC!!!", r)
	}
}

func (h *ChatHandler) sendWelcome() {
	msg := "# Welcome to server"
	if f, e := ioutil.ReadFile("/proc/cpuinfo"); e == nil {
		msg = msg + "\n" + string(f)
	}
	welcomeMsg := &StringMessage{
		BaseMessage: messageOf(ricaEvents.FROM_SERVER),
		Message:     msg,
	}
	h.transport.WriteMessage(welcomeMsg.Id, welcomeMsg)

	nickMsg := &NickMessage{
		BaseMessage: messageOf(ricaEvents.SET_NICK_REPLY),
		OldNick:     h.id,
		NewNick:     h.id,
	}

	h.transport.WriteMessage(nickMsg.Id, nickMsg)
}

func (h *ChatHandler) handleInternalMessage(msg interface{}) {
	baseMsg, ok := msg.(IEventMessage)

	if !ok {
		return
	}

	timer := StartStopWatch("handleInternnalMessage:" + h.id)
	defer timer.LogDuration()
	if err := h.transport.WriteMessage(baseMsg.Identity(), baseMsg); err != nil {
		log.Println("Unable to write socket message", err)
	}
}

func (h *ChatHandler) handleMessage(msg interface{}) {
	switch v := msg.(type) {
	case *ChatMessage:
		log.Println("Chat message", v.To)
		h.onChatMessage(v)
	case *StringMessage:
		log.Println("String message", v.Message)
		h.handleStringMessage(v)
	case *RecipientContentMessage:
		log.Println("Recipient message", v)
		h.onRecipientContentMessage(v)
	default:
		log.Println("Unknown", v)
	}
}

func (h *ChatHandler) handleStringMessage(msg *StringMessage) {
	switch msg.EventName {
	case ricaEvents.JOIN_GROUP_COMMAND:
		h.onJoinGroup(msg)
	case ricaEvents.LEAVE_GROUP_COMMAND:
		h.onLeaveGroup(msg)
	case ricaEvents.SET_NICK_COMMAND:
		h.onSetNick(msg)
	case ricaEvents.LIST_MEMBERS_COMMAND:
		h.onListMembers(msg)
	}
}

func (h *ChatHandler) onRecipientContentMessage(msg *RecipientContentMessage) {
	switch msg.EventName {
	case ricaEvents.SEND_RAW_MSG_COMMAND:
		log.Println("On send raw message")
		h.sendTo(ricaEvents.FROM_SERVER, msg.To, msg.Message)
	}
}

func (h *ChatHandler) onChatMessage(msg *ChatMessage) {
	strMsg := strings.TrimSpace(msg.Message)
	if len(strMsg) <= 0 || len(strMsg) > 512 {
		log.Println("Ignoring message... due to length volation")
		return
	}

	if _, ok := h.groups[msg.To]; ok {
		h.publish(msg.To, &ChatMessage{
			RecipientMessage: RecipientMessage{
				BaseMessage: messageOf(ricaEvents.GROUP_MSG_REPLY),
				To:          msg.To,
				From:        h.nick,
			},
			Message: msg.Message,
		})
	}
}

func (h *ChatHandler) onListMembers(msg *StringMessage) {
	groupName := msg.Message
	if groupName == "" {
		groupName = ricaEvents.FROM_SERVER
	}

	membersIds := h.groupInfoManager.GetUsers(groupName)
	members := make([]string, len(membersIds))
	i := 0
	for _, id := range membersIds {
		var foundNick bool
		members[i], foundNick = h.nickRegistry.NickOf(id)
		if !foundNick {
			members[i] = id
		}
		i++
	}

	h.incoming <- &RecipientContentMessage{
		RecipientMessage: RecipientMessage{
			BaseMessage: messageOf(ricaEvents.LIST_MEMBERS_REPLY),
			To:          groupName,
			From:        ricaEvents.FROM_SERVER,
		},
		Message: members,
	}
}

func (h *ChatHandler) onJoinGroup(msg *StringMessage) {
	timer := StartStopWatch("onJoinGroup:" + msg.Message)
	defer timer.LogDuration()

	h.Lock()
	h.groups[msg.Message] = struct{}{}
	h.Unlock()
	h.groupInfoManager.AddUser(msg.Message, h.id, h.incoming)

	h.publish(msg.Message, &RecipientMessage{
		BaseMessage: messageOf(ricaEvents.JOIN_GROUP_REPLY),
		To:          msg.Message,
		From:        h.nick,
	})
}

func (h *ChatHandler) onLeaveGroup(msg *StringMessage) {
	timer := StartStopWatch("onLeaveGroup:" + msg.Message)
	defer timer.LogDuration()

	h.publish(msg.Message, &RecipientMessage{
		BaseMessage: messageOf(ricaEvents.LEAVE_GROUP_REPLY),
		To:          msg.Message,
		From:        h.nick,
	})

	h.groupInfoManager.RemoveUser(msg.Message, h.id)
	h.Lock()
	delete(h.groups, msg.Message)
	h.Unlock()
}

func (h *ChatHandler) onSetNick(msg *StringMessage) {
	timer := StartStopWatch("onSetNick")
	defer timer.LogDuration()

	log.Println("Setting nick")
	old_nick := h.nick
	new_nick, err := h.nickRegistry.SetNick(h.id, msg.Message)

	if err == nil {
		h.nick = new_nick
		nickMsg := &NickMessage{
			BaseMessage: messageOf(ricaEvents.SET_NICK_REPLY),
			OldNick:     old_nick,
			NewNick:     new_nick,
		}
		err = h.transport.WriteMessage(nickMsg.Id, nickMsg)

		if err == nil {
			h.publishOnJoinedChannels(nickMsg.EventName, nickMsg)
			log.Println("On set nick", err)
			return
		}
	}

	log.Println("Unable to change nick", err)
}

func (h *ChatHandler) publishOnJoinedChannels(eventName string, msg interface{}) {
	timer := StartStopWatch("publishOnJoinedChannels:" + h.id)
	defer timer.LogDuration()

	h.Lock()
	defer h.Unlock()

	for g, _ := range h.groups {
		h.publish(g, &RecipientContentMessage{
			RecipientMessage: RecipientMessage{
				BaseMessage: messageOf(ricaEvents.MEMBER_NICK_SET_REPLY),
				To:          g,
				From:        h.id,
			},
			Message: msg,
		})
	}
}

func (h *ChatHandler) publish(groupName string, msg IEventMessage) {
	timer := StartStopWatch("publish:" + groupName)
	defer timer.LogDuration()

	h.transport.BeginBatch(msg.Identity(), msg)
	h.chatStore.Save(groupName, msg.Identity(), msg)

	groupMembers := h.groupInfoManager.GetUsers(groupName)
	log.Println("Publish for member count:", len(groupMembers))
	for _, id := range groupMembers {
		h.sendTo(groupName, id, msg)
	}

	h.transport.FlushBatch(msg.Identity())
}

func (h *ChatHandler) sendTo(groupName, name string, msg interface{}) {
	log.Println("sendTo", groupName, name)
	tmp := h.groupInfoManager.GetUserInfoObject(groupName, name)
	if tmp == nil {
		log.Println("Skipping publish to", name)
		return
	}

	if ch, ok := tmp.(chan interface{}); ok {
		select {
		case ch <- msg:
			log.Println("Published message to", name)
		case <-time.After(5 * time.Millisecond):
			log.Println("Publishing on", name, "timed out")
		}
	} else {
		log.Println("Invalid channel type skipping publish to", name)
	}
}

func (h *ChatHandler) Loop() {
	h.nickRegistry.Register(h.id, h.nick)
	h.sendWelcome()
	errorChannel := make(chan error)
	sockChannel := make(chan interface{}, 32)

	go h.socketLoop(sockChannel, errorChannel)
	defer h.recoverFromErrors()
	defer func() {
		close(errorChannel)
		close(sockChannel)
	}()

selectLoop:
	for {
		select {
		case <-time.After(30 * time.Second):
			h.incoming <- &PingMessage{
				BaseMessage: messageOf(ricaEvents.PING_COMMAND),
				Type:        int(time.Now().Unix()),
			}
		case m := <-h.incoming:
			h.handleInternalMessage(m)
		case m := <-sockChannel:
			h.handleMessage(m)
		case e := <-errorChannel:
			log.Println("Error received", e)
			break selectLoop
		}
	}

	log.Println("Finishing client select loop")
	h.Stop()
}

func (h *ChatHandler) Stop() {
	for g, _ := range h.groups {
		h.groupInfoManager.RemoveUser(g, h.id)
	}

	h.groups = make(map[string]interface{})
	h.nickRegistry.Unregister(h.id)
	close(h.incoming)
}
