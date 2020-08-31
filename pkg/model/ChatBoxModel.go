package model

type ChatBoxModel struct {
	AbstractModel
	Sender   string `json:"sender_id"`
	Receiver string `json:"receiver_id"`
}
