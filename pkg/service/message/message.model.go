package message

import "time"

type AbstractModel struct {
	ID        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
type Messages struct {
	AbstractModel
	SubjectSender string `json:"subject_sender"`
	Content       string `json:"content"`
	IdGroup       int    `json:"id_group"`
}
