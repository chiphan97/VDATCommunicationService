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
	rows, err := mess.Db.Query(statement, idChatBox)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		message := model.MessageModel{}
		err := rows.Scan(&message.ID, &message.IdChat, &message.Content, &message.SeenAt, &message.CreatedAt, &message.UpdatedAt, &message.DeletedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
func (mess *MessageRepoImpl) InsertMessage(messageModel model.MessageModel) error {
	statement := `INSERT INTO messages (id_chat,content) VALUES ($1,$2)`
	_, err := mess.Db.Exec(statement,
		messageModel.IdChat,
		messageModel.Content)
	return err
}
func (mess *MessageRepoImpl) UpdateMessageById(messageModel model.MessageModel) error {
	statement := `UPDATE messages SET content = $1, update_at = $2 WHERE id_mess = $3`
	_, err := mess.Db.Exec(statement, messageModel.Content, messageModel.UpdatedAt, messageModel.ID)
	return err
}
func (mess *MessageRepoImpl) DeleteMessageById(idMesssage int) error {
	statement := `DELETE FROM messages WHERE id_mess = $1`
	_, err := mess.Db.Exec(statement, idMesssage)
	return err
}
