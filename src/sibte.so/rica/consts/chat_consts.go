package ricaEvents

const (
    FROM_SERVER = "SERVER"

    PING_COMMAND         = "ping"
    JOIN_GROUP_COMMAND   = "join-group"
    LEAVE_GROUP_COMMAND  = "leave-group"
    SET_NICK_COMMAND     = "set-nick"
    SEND_MSG_COMMAND     = "send-msg"
    LIST_MEMBERS_COMMAND = "list-group"
    SEND_RAW_MSG_COMMAND = "send-raw-msg"

    PING_REPLY            = "pong"
    JOIN_GROUP_REPLY      = "group-join"
    LEAVE_GROUP_REPLY     = "group-leave"
    SET_NICK_REPLY        = "nick-set"
    MEMBER_NICK_SET_REPLY = "member-nick-set"
    NEW_MSG_REPLY         = "new-msg"
    NEW_RAW_MSG_REPLY     = "new-raw-msg"
    LIST_MEMBERS_REPLY    = "group-list"
    GROUP_MSG_REPLY       = "group-message"
    ERROR_MSG_REPLY       = "error-msg"

    ERROR_INVALID_MSGTYPE_ERR = "Chat handler received invalid message type"
)
