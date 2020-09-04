package useronline

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
)

func GetListUSerOnlineService(fil string) ([]User, error) {
	users := make([]User, 0)
	userOnlines, err := NewUserOnlineRepoImpl(database.DB).GetListUSerOnline(fil)
	if err != nil {
		return users, err
	}
	for _, userOnline := range userOnlines {
		user := User{
			UserID:   userOnline.UserID,
			Username: userOnline.Username,
			First:    userOnline.First,
			Last:     userOnline.Last,
		}
		users = append(users, user)
	}
	return users, nil
}
func AddUserOnlineService(online UserOnline) error {
	return NewUserOnlineRepoImpl(database.DB).AddUserOnline(online)
}
func DeleteUserOnlineService(socketid string) error {
	return NewUserOnlineRepoImpl(database.DB).DeleteUserOnline(socketid)
}
