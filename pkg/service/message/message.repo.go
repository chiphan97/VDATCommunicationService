package message

type Repo interface {
	GetMessagesByGroup(idChatBox int) ([]Messages, error)
	InsertMessage(message Messages) (int, error)
	//GetMessagesByChatBoxAndSeenAtOrderByCreatedAtLimit10(idChatBox int) ([]model.MessageModel, error)
	//UpdateMessageByChatBox(idChatBox int) error
	//DeleteMessageById(idMesssage int) error
	GetMessagesByGroupAndUser(idGroup int, subUser string) ([]Messages, error)
	GetContinueMessageByIdAndGroup(idMessage int, idGroup int) ([]Messages, error)
}
