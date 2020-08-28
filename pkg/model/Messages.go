package model

type Messages struct {
	AbstractModel
	SubjectSender string `json:"subject_sender"`
	Content       string `json:"content"`
	IdGroup       int    `json:"id_group"`
}
