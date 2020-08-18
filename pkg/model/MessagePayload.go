package model

type MessagePayload struct {
	ReceiverID string `json:"receiver_id"`
	SenderID   string `json:"sender_id"`
	Message    string `json:"message"`
}
