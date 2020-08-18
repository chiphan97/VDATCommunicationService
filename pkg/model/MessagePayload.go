package model

type MessagePayload struct {
	IdChat  int    `json:"id_chat"`
	Message string `json:"message"`
}
