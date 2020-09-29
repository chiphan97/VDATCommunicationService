package dchat

type Message struct {
	TypeEvent string `json:"type"`
	Data      Data   `json:"data"`
}
type Data struct {
	GroupId int    `json:"groupId"`
	Body    string `json:"body"`
	Status  string
}
