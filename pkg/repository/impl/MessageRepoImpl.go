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
func (mess *MessageRepoImpl) InsertMessages(messageModel model.MessageModel) error{
	statement := `INSERT INTO messages (id_chat,content,seen_at,create_at,update_at,delete_at) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := mess.Db.Exec(statement,
							messageModel.IdChat,
							messageModel.Content,
							messageModel.SeenAt,
							messageModel.CreatedAt,
							messageModel.UpdatedAt,
							messageModel.DeletedAt)
	return err
}