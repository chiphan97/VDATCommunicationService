package message

import "time"

type Dto struct {
	ID            uint       `json:"id"`
	SubjectSender string     `json:"subjectSender"`
	Content       string     `json:"content"`
	IdGroup       int        `json:"idGroup"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
