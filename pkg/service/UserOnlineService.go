package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetListUSerOnlineService(fil string) ([]model.User, error) {
	users := make([]model.User, 0)
	userOnlines, err := impl.NewUserOnlineRepoImpl(database.DB).GetListUSerOnline(fil)
	if err != nil {
		return users, err
	}
	for _, userOnline := range userOnlines {
		user := model.User{
			UserID:   userOnline.UserID,
			Username: userOnline.Username,
			First:    userOnline.First,
			Last:     userOnline.Last,
		}
		users = append(users, user)
	}
	return users, nil
}
func AddUserOnlineService(online model.UserOnline) error {
	return impl.NewUserOnlineRepoImpl(database.DB).AddUserOnline(online)
}
func DeleteUserOnlineService(socketid string) error {
	return impl.NewUserOnlineRepoImpl(database.DB).DeleteUserOnline(socketid)
}
