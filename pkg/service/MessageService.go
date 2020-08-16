package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
	"time"
)

func GetMessageByChatBoxService(id int) ([]model.MessageModel,error){
	messages, err := impl.NewMessageRepoImpl(database.DB).GetMessagesByChatBox(id)
	return messages, err
}
func InsertMessagesService() error{
	time := time.Now()
	abstract := model.AbstractModel{
		CreatedAt: &time,
		UpdatedAt: &time,
		DeletedAt: &time,
	}
	messageModel := model.MessageModel{
		abstract,
		"asndasjdasd",
		&time,
		5,
	}
	err := impl.NewMessageRepoImpl(database.DB).InsertMessages(messageModel)
	return err
}