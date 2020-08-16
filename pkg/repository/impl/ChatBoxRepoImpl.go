package impl

import (
	"database/sql"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository"
)

type ChatBoxRepoImpl struct {
	Db *sql.DB
}

func NewChatBoxRepoImpl(db *sql.DB) repository.ChatBoxRepo {
	return &ChatBoxRepoImpl{Db: db}
}
func (cbx *ChatBoxRepoImpl) GetChatBoxs() ([]model.ChatBoxModel, error) {
	chatboxs := make([]model.ChatBoxModel, 0)
	statement := `SELECT * FROM chatbox`
	rows, err := cbx.Db.Query(statement)
	if err != nil {
		return chatboxs, err
	}
	for rows.Next() {
		chatbox := model.ChatBoxModel{}
		err := rows.Scan(&chatbox.ID, &chatbox.Sender, &chatbox.Receiver, &chatbox.CreatedAt, &chatbox.UpdatedAt, &chatbox.DeletedAt)
		if err != nil {
			return chatboxs, err
		}
		chatboxs = append(chatboxs, chatbox)
	}
	return chatboxs, nil
}
