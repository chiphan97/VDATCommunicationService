package model

import "time"

type MessageModel struct {
	AbstractModel
	Content string `json:"content"`
	SeenAt *time.Time `json:"seen_at"`
	IdChat uint `json:"id_chat"`
}

type StatusType string

const (
	SEEN StatusType = "SEEN"
	SENT StatusType = "SENT"
	RECEIVED StatusType = "RECEIVED"
)