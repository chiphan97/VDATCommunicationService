package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetListUSerOnlineService(userHide string) ([]model.UserOnline, error) {
	return impl.NewUserOnlineRepoImpl(database.DB).GetListUSerOnline(userHide)
}
func AddUserOnlineService(online model.UserOnline) error {
	return impl.NewUserOnlineRepoImpl(database.DB).AddUserOnline(online)
}
func DeleteUserOnlineService(socketid string) error {
	return impl.NewUserOnlineRepoImpl(database.DB).DeleteUserOnline(socketid)
}
