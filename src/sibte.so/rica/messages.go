package rica

type BaseMessage struct {
	EventName string `json:"@"`
}

type HandshakeMessage struct {
	BaseMessage
	Nick  string   `json:"nick"`
	Rooms []string `json:"rooms"`
}

type RecipientMessage struct {
	BaseMessage
	To   string `json:"to"`
	From string `json:"from"`
}

type ChatMessage struct {
	BaseMessage
	To      string `json:"to"`
	From    string `json:"from"`
	Message string `json:"msg"`
}

type ErrorMessage struct {
	BaseMessage
	Type  string      `json:"error_type"`
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

type NickMessage struct {
	BaseMessage
	OldNick string `json:"oldNick"`
	NewNick string `json:"newNick"`
}

type RecipientContentMessage struct {
	BaseMessage
	To      string      `json:"to"`
	From    string      `json:"from"`
	Message interface{} `json:"pack_msg"`
}

type EventMessage struct {
	BaseMessage
	Name string      `json:"_en"`
	Body interface{} `json:"event"`
}

type StringMessage struct {
	BaseMessage
	Message string `json:"msg"`
}
