package rica

type HandshakeMessage struct {
	Nick  string   `json:"nick"`
	Rooms []string `json:"rooms"`
}

type RecipientMessage struct {
	To   string `json:"to"`
	From string `json:"from"`
}
type ChatMessage struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Message string `json:"msg"`
}

type ErrorMessage struct {
	Type  string      `json:"error_type"`
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

type NickMessage struct {
	OldNick string `json:"oldNick"`
	NewNick string `json:"newNick"`
}

type EventMessage struct {
	Name string      `json:"_en"`
	Body interface{} `json:"event"`
}
