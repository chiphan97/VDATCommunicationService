package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type ChatBoxRepo interface {
	FindChatBoxBySender(senderId string) ([]model.ChatBoxModel, error)
	FindChatBoxById(id uint) (model.ChatBoxModel, error)
	CreateChatBox(receiverId string, senderId string) (model.ChatBoxModel, error)
	DeleteChatBox(id uint) (bool, error)
}
