package rascore

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "strings"
    "sync"
    "time"

    "sibte.so/rascore/consts"

    "github.com/speps/go-hashids"
)

type userOutGoingInfo struct {
    channel chan interface{}
    ip      string
}

// ChatHandler to handle chat connection
type ChatHandler struct {
    sync.Mutex
    id               string
    nick             string
    groupInfoManager GroupInfoManager
    nickRegistry     *NickRegistry
    transport        IMessageTransport
    outgoingInfo     *userOutGoingInfo
    groups           map[string]interface{}
    blackList        map[string]interface{}
    chatStore        *ChatLogStore
}

var pHashID *hashids.HashID
var pSnowFlake = DefaultSnowFlake()

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

// NewChatHandler creates new ChatHandler
func NewChatHandler(
        nickReg *NickRegistry,
        groupInfoMan GroupInfoManager,
        trans IMessageTransport,
        store *ChatLogStore,
        ip string,
        blackList map[string]interface{}) *ChatHandler {
    if pHashID == nil {
        pHashID, _ = hashids.New()
    }

    uid, _ := pHashID.Encode([]int{
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
        blackList:        blackList,
        outgoingInfo:     &userOutGoingInfo{
            channel:      make(chan interface{}, 32),
            ip:           ip,
        },
        groups:           make(map[string]interface{}, 0),
    }

    return ret
}

func (h *ChatHandler) recoverFromErrors(tag string) {
    if r := recover(); r != nil {
        log.Println("!!!PANIC!!!", tag, r)
    }
}

func (h *ChatHandler) socketReaderLoop(socketChannel chan interface{}, errorChannel chan error, maxRetry uint) {
    defer h.recoverFromErrors("socketReaderLoop")
    readRetry := uint(0)

    for {
        if readRetry > maxRetry {
            break
        }
        
        msg, err := h.transport.ReadMessage()

        if err != nil {
            log.Println("Error reading socket message", err)
            readRetry = readRetry + 1
            time.Sleep(time.Duration(readRetry * 100) * time.Millisecond)
            continue
        }

        // If id blacklisted kill the channel
        if _, ok := h.blackList[h.id]; ok {
            h.blackList[h.outgoingInfo.ip] = struct {}{}
            errorChannel <- fmt.Errorf("User with %v and IP %v is banned", h.id, h.outgoingInfo.ip)
            break
        }

        // If ip is blacklisted just stop and return
        if _, ok := h.blackList[h.outgoingInfo.ip]; ok {
            errorChannel <- fmt.Errorf("User with %v and IP %v is banned", h.id, h.outgoingInfo.ip)
            break
        }

        if err != nil {
            errorChannel <- err
            break
        }

        // Decode and send messages to pipe
        var decMsg IEventMessage
        decMsg, err = transportDecodeMessage(msg)
        if err != nil && err.Error() == rasconsts.ERROR_INVALID_MSGTYPE_ERR {
            log.Println("Invalid message sent, skipping message....", err)
            continue
        }

        if err != nil {
            log.Println("Unable to decode message", err)
            continue
        }

        socketChannel <- decMsg
    }
}

func (h *ChatHandler) socketWriterLoop() {
    defer h.recoverFromErrors("socketWriterLoop")
    h.sendWelcome()

    for {
        select {
        case m, ok := <-h.outgoingInfo.channel:
            // channel read is not ok channel might be closed
            // try writing ping message to ensure channel is still open
            // if channel is closed write will panic
            if !ok {
                h.outgoingInfo.channel <- &PingMessage{
                    BaseMessage: messageOf(rasconsts.PING_COMMAND),
                    Type:        int(time.Now().Unix()),
                }
            }

            h.handleOutgoingMessage(m)
        case <-time.After(15 * time.Second):
            h.outgoingInfo.channel <- &PingMessage{
                BaseMessage: messageOf(rasconsts.PING_COMMAND),
                Type:        int(time.Now().Unix()),
            }
        }
    }
}

func (h *ChatHandler) writeMessage(m IEventMessage) error {
    mbytes, err := json.Marshal(m)
    if err == nil {
        log.Println("Writing message ", string(mbytes))
        err = h.transport.WriteMessage(m.Identity(), mbytes)
    }

    if err != nil {
        log.Println("Unable to write socket message", err)
    }

    return err
}

func (h *ChatHandler) sendWelcome() {
    msg := "# Welcome to server"
    if f, e := ioutil.ReadFile("/proc/cpuinfo"); e == nil {
        msg = msg + "\n" + string(f)
    }
    welcomeMsg := &StringMessage{
        BaseMessage: messageOf(rasconsts.FROM_SERVER),
        Message:     msg,
    }

    h.writeMessage(welcomeMsg)

    nickMsg := &NickMessage{
        BaseMessage: messageOf(rasconsts.SET_NICK_REPLY),
        OldNick:     h.id,
        NewNick:     h.id,
    }

    h.writeMessage(nickMsg)
}

func (h *ChatHandler) handleOutgoingMessage(msg interface{}) {
    baseMsg, ok := msg.(IEventMessage)

    if !ok {
        panic(fmt.Sprintf("Invalid outgoing message %v", msg))
    }

    timer := StartStopWatch("handleInternnalMessage:" + h.id)
    defer timer.LogDuration()
    if err := h.writeMessage(baseMsg); err != nil {
        h.Stop()
    }
}

func (h *ChatHandler) handleSocketMessage(msg interface{}) {
    switch v := msg.(type) {
    case *ChatMessage:
        h.onChatMessage(v)
    case *StringMessage:
        h.handleStringMessage(v)
    case *RecipientContentMessage:
        h.onRecipientContentMessage(v)
    }
}

func (h *ChatHandler) handleStringMessage(msg *StringMessage) {
    switch msg.EventName {
    case rasconsts.JOIN_GROUP_COMMAND:
        h.onJoinGroup(msg)
    case rasconsts.LEAVE_GROUP_COMMAND:
        h.onLeaveGroup(msg)
    case rasconsts.SET_NICK_COMMAND:
        h.onSetNick(msg)
    case rasconsts.LIST_MEMBERS_COMMAND:
        h.onListMembers(msg)
    }
}

func (h *ChatHandler) onRecipientContentMessage(msg *RecipientContentMessage) {
    switch msg.EventName {
    case rasconsts.SEND_RAW_MSG_COMMAND:
        h.sendTo(rasconsts.FROM_SERVER, msg.To, msg.Message)
    }
}

func (h *ChatHandler) onChatMessage(msg *ChatMessage) {
    strMsg := strings.TrimSpace(msg.Message)
    if len(strMsg) <= 0 || len(strMsg) > 512 {
        return
    }

    if _, ok := h.groups[msg.To]; ok {
        h.publish(msg.To, &ChatMessage{
            RecipientMessage: RecipientMessage{
                BaseMessage: messageOf(rasconsts.GROUP_MSG_REPLY),
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
        groupName = rasconsts.FROM_SERVER
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

    h.outgoingInfo.channel <- &RecipientContentMessage{
        RecipientMessage: RecipientMessage{
            BaseMessage: messageOf(rasconsts.LIST_MEMBERS_REPLY),
            To:          groupName,
            From:        rasconsts.FROM_SERVER,
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
    h.groupInfoManager.AddUser(msg.Message, h.id, h.outgoingInfo)

    h.publish(msg.Message, &RecipientMessage{
        BaseMessage: messageOf(rasconsts.JOIN_GROUP_REPLY),
        To:          msg.Message,
        From:        h.nick,
    })
}

func (h *ChatHandler) onLeaveGroup(msg *StringMessage) {
    timer := StartStopWatch("onLeaveGroup:" + msg.Message)
    defer timer.LogDuration()

    h.publish(msg.Message, &RecipientMessage{
        BaseMessage: messageOf(rasconsts.LEAVE_GROUP_REPLY),
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

    oldNick := h.nick
    newNick, err := h.nickRegistry.SetBestPossibleNick(h.id, msg.Message)

    if err == nil {
        h.nick = newNick
        nickMsg := &NickMessage{
            BaseMessage: messageOf(rasconsts.SET_NICK_REPLY),
            OldNick:     oldNick,
            NewNick:     newNick,
        }
        err = h.writeMessage(nickMsg)

        if err == nil {
            h.publishOnJoinedChannels(nickMsg.EventName, nickMsg)
            return
        }
    }

    log.Println("Unable to change nick", err)
}

func (h *ChatHandler) publishOnJoinedChannels(eventName string, msg interface{}) {
    timer := StartStopWatch("publishOnJoinedChannels:" + h.id)
    defer timer.LogDuration()
    joinedGroups := make([]string, 0, len(h.groups))

    h.Lock()
    for g := range h.groups {
        joinedGroups = append(joinedGroups, g)
    }
    h.Unlock()

    for _, g := range joinedGroups {
        h.publish(g, &RecipientContentMessage{
            RecipientMessage: RecipientMessage{
                BaseMessage: messageOf(rasconsts.MEMBER_NICK_SET_REPLY),
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

    msg.Stamp()
    h.chatStore.Save(groupName, msg.Identity(), msg)

    groupMembers := h.groupInfoManager.GetUsers(groupName)
    for _, id := range groupMembers {
        h.sendTo(groupName, id, msg)
    }
}

func (h *ChatHandler) sendTo(groupName, name string, msg interface{}) {
    defer h.recoverFromErrors("sendTo")
    tmp := h.groupInfoManager.GetUserInfoObject(groupName, name)
    if tmp == nil {
        return
    }

    if inf, ok := tmp.(*userOutGoingInfo); ok {
        select {
        case inf.channel <- msg:
            break
        case <-time.After(200 * time.Millisecond):
            log.Println("Publishing on", name, "timed out")
        }
    } else {
        log.Println("Invalid channel type skipping publish to", name)
    }
}

// Loop over incoming and out going socket channels
func (h *ChatHandler) Loop() {
    defer h.recoverFromErrors("Loop")
    h.nickRegistry.Register(h.id, h.nick)
    h.groups[rasconsts.FROM_SERVER] = struct{}{}
    h.groupInfoManager.AddUser(rasconsts.FROM_SERVER, h.id, h.outgoingInfo)

    readErrorChannel := make(chan error)
    sockChannel := make(chan interface{}, 32)

    go h.socketReaderLoop(sockChannel, readErrorChannel, 3)
    go h.socketWriterLoop()
    defer func() {
        close(readErrorChannel)
        close(sockChannel)
    }()

selectLoop:
    for {
        select {
        case m := <-sockChannel:
            h.handleSocketMessage(m)
        case e := <-readErrorChannel:
            log.Println("Error received", e)
            break selectLoop
        }
    }

    h.Stop()
}

// Stop a client connection and perform cleanup
func (h *ChatHandler) Stop() {
    close(h.outgoingInfo.channel)
    currentGroupsMap := h.groups
    h.groups = make(map[string]interface{})
    joinedGroups := make([]string, 0, len(h.groups))

    for g := range currentGroupsMap {
        joinedGroups = append(joinedGroups, g)
        h.groupInfoManager.RemoveUser(g, h.id)
    }

    h.nickRegistry.Unregister(h.id)
    for _, groupName := range joinedGroups {
        h.publish(groupName, &RecipientMessage{
            BaseMessage: messageOf(rasconsts.LEAVE_GROUP_REPLY),
            To:          groupName,
            From:        h.nick,
        })
    }
}
