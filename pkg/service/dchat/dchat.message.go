package dchat

type Message struct {
	TypeEvent string `json:"type" example:"send_text"`
	Data      Data   `json:"data" `
	Client    string
}
type Data struct {
	GroupId  int    `json:"groupId" example:1`
	Body     string `json:"body" example:"tin nhan moi"`
	Sender   string `example:"null"`
	SocketID string `json:"socketId" example:"9999"`
	Status   string `example:"null"`
}
