package impl

import (
	"database/sql"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository"
)

type MessageRepoImpl struct {
	Db *sql.DB
}

func NewMessageRepoImpl(db *sql.DB) repository.MessageRepo {
	return &MessageRepoImpl{Db: db}
}
func (mess *MessageRepoImpl) GetMessagesByChatBox(idChatBox int) ([]model.MessageModel, error) {
	messages := make([]model.MessageModel, 0)
	statement := `SELECT * FROM messages WHERE id_chat = $1`
	rows, err := mess.Db.Query(statement,idChatBox)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		message := model.MessageModel{}
		err := rows.Scan(&message.ID,&message.IdChat, &message.Content, &message.SeenAt, &message.CreatedAt, &message.UpdatedAt, &message.DeletedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
