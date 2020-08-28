package handler

type WsMessage struct {
	From   string      `json:"from"`
	To     []string    `json:"to"`
	Body   interface{} `json:"body"`
	Status string
}
