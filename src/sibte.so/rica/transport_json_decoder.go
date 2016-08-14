package rica

import (
    "encoding/json"
    "errors"
    "reflect"

    "sibte.so/rica/consts"
)

var pEventToStructMap map[string]reflect.Type

func initChatHandlerTypes() {
    if pEventToStructMap == nil {
        pEventToStructMap = make(map[string]reflect.Type)
        pEventToStructMap[ricaEvents.SEND_MSG_COMMAND] = reflect.TypeOf(ChatMessage{})
        pEventToStructMap[ricaEvents.JOIN_GROUP_COMMAND] = reflect.TypeOf(StringMessage{})
        pEventToStructMap[ricaEvents.LEAVE_GROUP_COMMAND] = reflect.TypeOf(StringMessage{})
        pEventToStructMap[ricaEvents.SET_NICK_COMMAND] = reflect.TypeOf(StringMessage{})
        pEventToStructMap[ricaEvents.LIST_MEMBERS_COMMAND] = reflect.TypeOf(StringMessage{})
        pEventToStructMap[ricaEvents.NEW_RAW_MSG_REPLY] = reflect.TypeOf(RecipientContentMessage{})
        pEventToStructMap[ricaEvents.PING_REPLY] = reflect.TypeOf(BaseMessage{})
    }
}

func transportDecodeMessage(msg []byte) (ret IEventMessage, rErr error) {
    eventMsg := &BaseMessage{}
    rErr = json.Unmarshal(msg, eventMsg)
    if rErr != nil {
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

    ret = reflect.New(mType).Interface().(IEventMessage)
    rErr = json.Unmarshal(msg, ret)
    ret.Stamp()
    return
}
