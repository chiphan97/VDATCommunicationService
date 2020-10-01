package dchat

type Message struct {
	TypeEvent string `json:"type"`
	Data      Data   `json:"data"`
	Client    string
}
type Data struct {
	GroupId  int    `json:"groupId"`
	Body     string `json:"body"`
	Sender   string
	SocketID string `json:"socketId"`
	Status   string
}
