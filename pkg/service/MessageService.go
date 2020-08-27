package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetMessagesByGroupAndUserService(idGroup int, subUser string) ([]model.Messages, error) {
	return impl.NewMessageRepoImpl(database.DB).GetMessagesByGroupAndUser(idGroup, subUser)
}
