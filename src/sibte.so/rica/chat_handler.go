package rica

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/speps/go-hashids"
)

const (
	_FROM_SERVER = "SERVER"

	_JOIN_GROUP_COMMAND   = "join-group"
	_LEAVE_GROUP_COMMAND  = "leave-group"
	_SET_NICK_COMMAND     = "set-nick"
	_SEND_MSG_COMMAND     = "send-msg"
	_SEND_RAW_MSG_COMMAND = "send-raw-msg"

	_JOIN_GROUP_REPLY      = "group-join"
	_LEAVE_GROUP_REPLY     = "group-leave"
	_SET_NICK_REPLY        = "nick-set"
	_MEMBER_NICK_SET_REPLY = "member-nick-set"
	_NEW_MSG_REPLY         = "new-msg"
	_NEW_RAW_MSG_REPLY     = "new-raw-msg"
	_GROUP_MSG_REPLY       = "group-message"
	_ERROR_MSG_REPLY       = "error-msg"

	_ERROR_INVALID_MSGTYPE_ERR = "Chat handler received invalid message type"
)

type ChatHandler struct {
	sync.Mutex
	id                  string
	nick                string
	groupInfoManager    GroupInfoManager
	nickRegistry        *NickRegistry
	connectionReadLock  *sync.Mutex
	connectionWriteLock *sync.Mutex
	connection          *websocket.Conn
	incoming            chan interface{}
	groups              map[string]interface{}
}

var pEventToStructMap map[string]reflect.Type
var pHashId *hashids.HashID = hashids.New()

func initChatHandlerTypes() {
	if pEventToStructMap == nil {
		pEventToStructMap = make(map[string]reflect.Type)
		pEventToStructMap[_SEND_MSG_COMMAND] = reflect.TypeOf(ChatMessage{})
		pEventToStructMap[_JOIN_GROUP_COMMAND] = reflect.TypeOf(StringMessage{})
		pEventToStructMap[_LEAVE_GROUP_COMMAND] = reflect.TypeOf(StringMessage{})
		pEventToStructMap[_SET_NICK_COMMAND] = reflect.TypeOf(StringMessage{})
		pEventToStructMap[_NEW_RAW_MSG_REPLY] = reflect.TypeOf(RecipientContentMessage{})
	}
}

func NewChatHandler(nickReg *NickRegistry, groupInfoMan GroupInfoManager, connection *websocket.Conn) *ChatHandler {
	log.Println("New chat handler")
	initChatHandlerTypes()
	uid, _ := pHashId.Encode([]int{
		int(rand.Int31n(1000)),
		int(rand.Int31n(1000)),
		int(rand.Int31n(1000)),
	})
	return &ChatHandler{
		id:                  uid,
		nick:                uid,
		nickRegistry:        nickReg,
		groupInfoManager:    groupInfoMan,
		connectionReadLock:  &sync.Mutex{},
		connectionWriteLock: &sync.Mutex{},
		connection:          connection,
		incoming:            make(chan interface{}, 32),
		groups:              make(map[string]interface{}, 0),
	}
}

func pDecodeMessage(msg []byte) (ret interface{}, rErr error) {
	eventMsg := &BaseMessage{}
	rErr = json.Unmarshal(msg, eventMsg)
	if rErr != nil {
		log.Println("Unable to decode message type from msg", string(msg))
		ret = nil
		return
	}

	var mType reflect.Type
	var ok bool
	if mType, ok = pEventToStructMap[eventMsg.EventName]; !ok {
		rErr = errors.New("Invalid message type")
		ret = nil
		return
	}

	log.Println("Deserializing ", mType.Name())
	ret = reflect.New(mType).Interface()
	rErr = json.Unmarshal(msg, ret)
	return
}

func (h *ChatHandler) readSocket() (interface{}, error) {
	h.connectionReadLock.Lock()
	msgType, msg, err := h.connection.ReadMessage()
	h.connectionReadLock.Unlock()

	log.Println("Message Type", msgType, "Message", string(msg), "Error", err)

	if err != nil {
		return nil, err
	}

	if msgType != websocket.TextMessage {
		return nil, errors.New(_ERROR_INVALID_MSGTYPE_ERR)
	}

	if jsonMsg, e := pDecodeMessage(msg); e == nil {
		return jsonMsg, nil
	} else {
		log.Println("Invalid message", e)
	}

	return nil, err
}

func (h *ChatHandler) writeSocket(msg interface{}) error {
	h.connectionWriteLock.Lock()
	defer h.connectionWriteLock.Unlock()
	return h.connection.WriteJSON(msg)
}

func (h *ChatHandler) socketLoop(socketChannel chan interface{}, errorChannel chan error) {
	for {
		msg, err := h.readSocket()

		// If message type was invalid
		if err != nil && err.Error() == _ERROR_INVALID_MSGTYPE_ERR {
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
	welcomeMsg := &StringMessage{
		BaseMessage: BaseMessage{_FROM_SERVER},
		Message:     "Welcome",
	}
	h.writeSocket(welcomeMsg)

	h.writeSocket(&NickMessage{
		BaseMessage: BaseMessage{_SET_NICK_REPLY},
		OldNick:     h.id,
		NewNick:     h.id,
	})
}

func (h *ChatHandler) handleInternalMessage(msg interface{}) {
	timer := StartStopWatch("handleInternnalMessage:" + h.id)
	defer timer.LogDuration()
	if err := h.writeSocket(msg); err != nil {
		log.Println("Unable to write socket message", err)
	}
}

func (h *ChatHandler) handleMessage(msg interface{}) {
	switch v := msg.(type) {
	case *ChatMessage:
		log.Println("Chat message", v.To, "=>", v.Message)
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
	case _JOIN_GROUP_COMMAND:
		h.onJoinGroup(msg)
	case _LEAVE_GROUP_COMMAND:
		h.onLeaveGroup(msg)
	case _SET_NICK_COMMAND:
		h.onSetNick(msg)
	}
}

func (h *ChatHandler) onRecipientContentMessage(msg *RecipientContentMessage) {
	switch msg.EventName {
	case _SEND_RAW_MSG_COMMAND:
		log.Println("On send raw message")
	}
}

func (h *ChatHandler) onChatMessage(msg *ChatMessage) {
	if _, ok := h.groups[msg.To]; ok {
		h.publish(msg.To, &ChatMessage{
			BaseMessage: BaseMessage{_GROUP_MSG_REPLY},
			To:          msg.To,
			From:        h.nick,
			Message:     msg.Message,
		})
	}
}

func (h *ChatHandler) onJoinGroup(msg *StringMessage) {
	timer := StartStopWatch("onJoinGroup")
	defer timer.LogDuration()

	h.Lock()
	h.groups[msg.Message] = struct{}{}
	h.Unlock()
	h.groupInfoManager.AddUser(msg.Message, h.id, h.incoming)

	h.publish(msg.Message, &RecipientMessage{
		BaseMessage: BaseMessage{_JOIN_GROUP_REPLY},
		To:          msg.Message,
		From:        h.nick,
	})
}

func (h *ChatHandler) onLeaveGroup(msg *StringMessage) {
	timer := StartStopWatch("onLeaveGroup")
	defer timer.LogDuration()

	h.Lock()
	delete(h.groups, msg.Message)
	h.Unlock()
	h.groupInfoManager.RemoveUser(msg.Message, h.id)

	h.publish(msg.Message, &RecipientMessage{
		BaseMessage: BaseMessage{_LEAVE_GROUP_REPLY},
		To:          msg.Message,
		From:        h.nick,
	})
}

func (h *ChatHandler) onSetNick(msg *StringMessage) {
	timer := StartStopWatch("onSetNick")
	defer timer.LogDuration()

	log.Println("Setting nick")
	old_nick := h.nick
	new_nick, err := h.nickRegistry.SetNick(h.id, msg.Message)

	if err == nil {
		nickMsg := &NickMessage{
			BaseMessage: BaseMessage{_SET_NICK_REPLY},
			OldNick:     old_nick,
			NewNick:     new_nick,
		}
		err = h.writeSocket(nickMsg)

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
			BaseMessage: BaseMessage{_MEMBER_NICK_SET_REPLY},
			To:          g,
			From:        h.id,
			Message:     msg,
		})
	}
}

func (h *ChatHandler) publish(groupName string, msg interface{}) {
	timer := StartStopWatch("publish:" + groupName)
	defer timer.LogDuration()

	groupMembers := h.groupInfoManager.GetUsers(groupName)
	log.Println("Publish for member count:", len(groupMembers))
	for _, name := range groupMembers {
		tmp := h.groupInfoManager.GetUserInfoObject(groupName, name)
		if tmp == nil {
			log.Println("Skipping publish to", name)
			continue
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
