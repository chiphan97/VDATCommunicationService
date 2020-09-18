package message

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
)

func GetMessagesByGroupAndUserService(idGroup int, subUser string) ([]Messages, error) {
	return NewMessageRepoImpl(database.DB).GetMessagesByGroupAndUser(idGroup, subUser)
}
