package model
type ChatBoxModel struct {
	AbstractModel
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Messages []MessageModel
}