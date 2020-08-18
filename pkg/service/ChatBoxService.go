package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func FindChatBoxBySender(senderId string) ([]model.ChatBoxModel, error) {
	return impl.NewChatBoxRepoImpl(database.DB).FindChatBoxBySender(senderId)
}

func FindChatBoxById(id uint) (model.ChatBoxModel, error) {
	return impl.NewChatBoxRepoImpl(database.DB).FindChatBoxById(id)
}

func CreateChatBox(receiverId string, senderId string) (model.ChatBoxModel, error) {
	return impl.NewChatBoxRepoImpl(database.DB).CreateChatBox(receiverId, senderId)
}

func DeleteChatBox(id uint) (bool, error) {
	return impl.NewChatBoxRepoImpl(database.DB).DeleteChatBox(id)
}
