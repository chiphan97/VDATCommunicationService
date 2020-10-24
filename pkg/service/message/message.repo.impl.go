package message

import (
	"database/sql"
)

type MessageRepoImpl struct {
	Db *sql.DB
}

func RepoImpl(db *sql.DB) Repo {
	return &MessageRepoImpl{Db: db}
}

func (mess *MessageRepoImpl) GetMessagesByGroup(idChatBox int) ([]Messages, error) {
	messages := make([]Messages, 0)
	statement := `SELECT * FROM messages WHERE id_group = $1 ORDER BY created_at DESC LIMIT 20`
	rows, err := mess.Db.Query(statement, idChatBox)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		m := Messages{}
		err := rows.Scan(&m.ID,
			&m.SubjectSender,
			&m.Content,
			&m.IdGroup,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, m)
	}
	defer rows.Close()
	return messages, nil
}
func (mess *MessageRepoImpl) InsertMessage(message Messages) error {
	statement := `INSERT INTO messages (user_sender,content,id_group) VALUES ($1,$2,$3)`
	_, err := mess.Db.Exec(statement,
		message.SubjectSender,
		message.Content,
		message.IdGroup)
	return err
}

//func (mess *MessageRepoImpl) GetMessagesByChatBoxAndSeenAtOrderByCreatedAtLimit10(idChatBox int) ([]model.MessageModel, error) {
//	message := make([]model.MessageModel, 0)
//	statement := `SELECT * FROM message WHERE id_chat = $1 AND seen_at IS NULL ORDER BY created_at LIMIT 10`
//	rows, err := mess.Db.Query(statement, idChatBox)
//	if err != nil {
//		return message, err
//	}
//	for rows.Next() {
//		message := model.MessageModel{}
//		err := rows.Scan(&message.ID, &message.IdChat, &message.Content, &message.SeenAt, &message.CreatedAt, &message.UpdatedAt, &message.DeletedAt)
//		if err != nil {
//			return message, err
//		}
//		message = append(message, message)
//	}
//	return message, nil
//}
//func (mess *MessageRepoImpl) UpdateMessageByChatBox(idChatBox int) error {
//	statement := `UPDATE message SET seen_at=now() WHERE id_chat = $1`
//	_, err := mess.Db.Exec(statement, idChatBox)
//	return err
//}
func (mess *MessageRepoImpl) GetMessagesByGroupAndUser(idGroup int, subUser string) ([]Messages, error) {
	messages := make([]Messages, 0)
	statement := `SELECT m.id_mess,m.subject_sender,m.content,m.created_at,m.updated_at,m.deleted_at,m.id_group
					FROM (SELECT * FROM dchat.public.message WHERE id_group = $1) AS m
					LEFT JOIN Messages_Delete AS md
					ON m.id_mess = md.id_mess
					WHERE  m.created_at >= (SELECT g_u.last_deleted_messages
											FROM Groups AS g
											INNER JOIN Groups_Users AS g_u
											ON g.id_group = $1
											AND g.id_group = g_u.id_group
											WHERE g_u.sub_user_join = $2
											ORDER BY g_u.last_deleted_messages DESC
											LIMIT 1)
					AND (md.sub_user_deleted != $2 OR md.sub_user_deleted IS NULL)
					`
	rows, err := mess.Db.Query(statement, idGroup, subUser)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		message := Messages{}
		err := rows.Scan(&message.ID,
			&message.SubjectSender,
			&message.Content,
			&message.CreatedAt,
			&message.UpdatedAt,
			&message.DeletedAt,
			&message.IdGroup)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

//func (mess *MessageRepoImpl) DeleteMessageById(idMesssage int) error {
//	statement := `DELETE FROM message WHERE id_mess = $1`
//	_, err := mess.Db.Exec(statement, idMesssage)
//	return err
//}
