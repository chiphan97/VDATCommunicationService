package useronline

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
)

//func GetListUSerOnlineService(fil string) ([]Dto, error) {
//	users := make([]Dto, 0)
//	userOnlines, err := NewUserOnlineRepoImpl(database.DB).GetListUSerOnline(fil)
//	if err != nil {
//		return users, err
//	}
//	for _, userOnline := range userOnlines {
//		user := Dto{
//			UserID:   userOnline.UserID,
//			Username: userOnline.Username,
//			First:    userOnline.First,
//			Last:     userOnline.Last,
//		}
//		users = append(users, user)
//	}
//	return users, nil
//}
func AddUserOnlineService(payload Payload) error {
	online := payload.convertToModel()

	err := NewUserOnlineRepoImpl(database.DB).AddUserOnline(online)
	if err != nil {
		return err
	}
	return nil
}
func DeleteUserOnlineService(socketid string) error {
	return NewUserOnlineRepoImpl(database.DB).DeleteUserOnline(socketid)
}
