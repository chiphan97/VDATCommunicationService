package message

type PayLoad struct {
	ID            uint   `json:"id"`
	SubjectSender string `json:"subjectSender"`
	Content       string `json:"content"`
	IdGroup       int    `json:"idGroup"`
}

func (p *PayLoad) convertToModel() {

}
