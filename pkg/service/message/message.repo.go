package message

type MessageRepo interface {
	//GetMessagesByChatBox(idChatBox int) ([]model.MessageModel, error)
	InsertMessage(message Messages) error
	//GetMessagesByChatBoxAndSeenAtOrderByCreatedAtLimit10(idChatBox int) ([]model.MessageModel, error)
	//UpdateMessageByChatBox(idChatBox int) error
	//DeleteMessageById(idMesssage int) error
	GetMessagesByGroupAndUser(idGroup int, subUser string) ([]Messages, error)
}
