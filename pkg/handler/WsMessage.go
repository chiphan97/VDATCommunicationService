package handler

type WsMessage struct {
	From   string
	To     []string
	Body   interface{}
	Status string
}
