package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type MessageRepo interface {
	GetMessagesByChatBox(idChatBox int) ([]model.MessageModel, error)
	InsertMessages(messageModel model.MessageModel) error
}
