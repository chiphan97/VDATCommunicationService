package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetMessageByChatBoxService(id int) ([]model.MessageModel, error) {
	messages, err := impl.NewMessageRepoImpl(database.DB).GetMessagesByChatBox(id)
	return messages, err
}
func InsertMessagesService(messageModel model.MessageModel) error {
	err := impl.NewMessageRepoImpl(database.DB).InsertMessage(messageModel)
	return err
}

//func UpdateMessageService(messageModel model.MessageModel) error {
//	err := impl.NewMessageRepoImpl(database.DB).UpdateMessageById(messageModel)
//	return err
//}
//func DeleteMessageService(idMessage int) error {
//	err := impl.NewMessageRepoImpl(database.DB).DeleteMessageById(idMessage)
//	return err
//}
