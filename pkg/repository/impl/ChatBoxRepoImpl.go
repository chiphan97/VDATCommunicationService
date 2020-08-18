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

func (cbx *ChatBoxRepoImpl) CreateChatBox(receiverId string) (model.ChatBoxModel, error) {
	chatBox := model.ChatBoxModel{}

	// TODO kiểm tra receiverId có tồn tại không
	receiverId = "receiver"

	// TODO lấy user hiện tại
	senderId := "anonymousUser"

	// kiểm tra chatBox đã tồn tại chưa
	statement := `SELECT * FROM chatboxs WHERE sender = $1 AND receiver = $2`
	rows, err := cbx.Db.Query(statement, senderId, receiverId)
	if err != nil {
		panic(err)
	} else if rows.Next() {
		panic("Chat box is existed")
	}

	// lưu hội thoại vào db
	statement = `INSERT INTO chatboxs(sender, receiver) VALUE($1, $2)`
	result, err := cbx.Db.Exec(statement, senderId, receiverId)
	if err != nil {
		panic(err)
	} else if rows.Next() {
		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		return cbx.FindChatBoxById(uint(lastId))
	}

	return chatBox, nil
}

func (cbx *ChatBoxRepoImpl) DeleteChatBox(id uint) (bool, error) {
	statement := `DELETE FROM chatboxs WHERE id = $1`
	result, err := cbx.Db.Exec(statement, id)

	if err != nil {
		panic(err)
	} else if result != nil {
		return true, nil
	}

	return false, nil
}

func (cbx *ChatBoxRepoImpl) FindChatBoxById(id uint) (model.ChatBoxModel, error) {
	chatBox := model.ChatBoxModel{}
	statement := `SELECT * FROM chatboxs WHERE id = $1`

	rows, err := cbx.Db.Query(statement, id)
	if err != nil {
		panic(err)
	} else if rows.Next() {
		err := rows.Scan(&chatBox.ID, &chatBox.Sender, &chatBox.Receiver, &chatBox.CreatedAt, &chatBox.UpdatedAt, &chatBox.DeletedAt)
		if err != nil {
			panic(err)
		}
	}

	return chatBox, nil
}

func (cbx *ChatBoxRepoImpl) FindChatBoxBySender(senderId string) ([]model.ChatBoxModel, error) {
	listChatBox := make([]model.ChatBoxModel, 0)

	statement := `SELECT * FROM chatboxs WHERE sender = $1`
	rows, err := cbx.Db.Query(statement, senderId)
	if err != nil {
		return listChatBox, nil
	}

	for rows.Next() {
		chatBox := model.ChatBoxModel{}

		err := rows.Scan(&chatBox.ID, &chatBox.Sender, &chatBox.Receiver, &chatBox.CreatedAt, &chatBox.UpdatedAt, &chatBox.DeletedAt)

		if err != nil {
			return listChatBox, err
		}

		listChatBox = append(listChatBox, chatBox)
	}

	return listChatBox, nil
}
