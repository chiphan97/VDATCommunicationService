package message

import "time"

type AbstractModel struct {
	ID        uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
type Messages struct {
	AbstractModel
	SubjectSender string
	Content       string
	IdGroup       int
	ParentId      int
}

func (m *Messages) convertToDTO() Dto {
	message := Dto{
		ID:            m.ID,
		SubjectSender: m.SubjectSender,
		Content:       m.Content,
		IdGroup:       m.IdGroup,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
	return message
}
