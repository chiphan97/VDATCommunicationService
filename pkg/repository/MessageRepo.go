package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type MessageRepo interface {
	GetMessagesByChatBox(idChatBox int) ([]model.MessageModel, error)
	InsertMessage(messageModel model.MessageModel) error
	//UpdateMessageById(messageModel model.MessageModel) error
	//DeleteMessageById(idMesssage int) error
}
