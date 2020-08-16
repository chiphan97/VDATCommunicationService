package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetMessageByChatBoxService(id int) ([]model.MessageModel,error){

	database.Connect()
	messages, err := impl.NewMessageRepoImpl(database.DB).GetMessagesByChatBox(id)
	return messages, err

}