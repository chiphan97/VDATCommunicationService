package model

import "time"

type MessageModel struct {
	AbstractModel
	Content string `json:"content"`
	Status StatusType `json:"status"`
	SeenAt *time.Time `json:"seen_at"`
	IdChat uint `json:"id_chat"`
}

type StatusType string

const (
	SEEN StatusType = "SEEN"
	SENT StatusType = "SENT"
	RECEIVED StatusType = "RECEIVED"
)