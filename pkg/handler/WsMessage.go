package handler

type WsMessage struct {
	Type string `json:"type"`
	//From   string      `json:"from"`
	//To     []string    `json:"to"`
	Body   interface{} `json:"data"`
	Status string
}
