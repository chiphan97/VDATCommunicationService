package impl

import (
	"database/sql"
	"golangproject/model"
	"golangproject/repository"
)
type MessageRepoImpl struct {
	Db *sql.DB
}
func NewMessageRepoImpl(db *sql.DB) repository.MessageRepo{
	return &MessageRepoImpl{Db:db}
}
func (mess *MessageRepoImpl) GetMessages() ([]model.MessageModel, error){
	messages := make([]model.MessageModel,0)
	statement := `SELECT * FROM message`
	rows, err := mess.Db.Query(statement)
	if err!=nil{
		return messages, err
	}
	for rows.Next(){
		message := model.MessageModel{}
		err := rows.Scan(&message.ID,&message.Content,&message.SeenAt,&message.CreatedAt,&message.UpdatedAt,&message.DeletedAt,&message.IdChat,&message.Status)
		if err!=nil{
			return messages, err
		}
		messages = append(messages,message)
	}

	return messages, nil
}