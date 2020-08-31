package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type MessageRepo interface {
	//GetMessagesByChatBox(idChatBox int) ([]model.MessageModel, error)
	//InsertMessage(messageModel model.MessageModel) error
	//GetMessagesByChatBoxAndSeenAtOrderByCreatedAtLimit10(idChatBox int) ([]model.MessageModel, error)
	//UpdateMessageByChatBox(idChatBox int) error
	//DeleteMessageById(idMesssage int) error
	GetMessagesByGroupAndUser(idGroup int, subUser string) ([]model.Messages, error)
}
