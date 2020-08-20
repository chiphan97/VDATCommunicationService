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
func (mess *MessageRepoImpl) GetMessagesByChatBoxAndSeenAtOrderByCreatedAtLimit10(idChatBox int) ([]model.MessageModel, error) {
	messages := make([]model.MessageModel, 0)
	statement := `SELECT * FROM messages WHERE id_chat = $1 AND seen_at IS NULL ORDER BY created_at LIMIT 10`
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
func (mess *MessageRepoImpl) UpdateMessageByChatBox(idChatBox int) error {
	statement := `UPDATE messages SET seen_at=now() WHERE id_chat = $1`
	_, err := mess.Db.Exec(statement, idChatBox)
	return err
}

//func (mess *MessageRepoImpl) DeleteMessageById(idMesssage int) error {
//	statement := `DELETE FROM messages WHERE id_mess = $1`
//	_, err := mess.Db.Exec(statement, idMesssage)
//	return err
//}
