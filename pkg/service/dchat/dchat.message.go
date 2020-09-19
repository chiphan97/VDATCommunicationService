package dchat

type Message struct {
	From    string `json:"from"`
	GroupId int    `json:"group_id"`
	Body    string `json:"body"`
	status  string
}
