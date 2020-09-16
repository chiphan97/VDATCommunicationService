package useronline

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
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
func AddUserOnlineService(payload Payload, deltailPayload userdetail.Payload) error {
	online := payload.convertToModel()
	check, err := userdetail.GetUserDetailByIDService(payload.UserID)
	if err != nil {
		return err
	}
	if check == (userdetail.Dto{}) {
		err := userdetail.AddUserDetailService(deltailPayload)
		if err != nil {
			return err
		}
	}
	err = NewUserOnlineRepoImpl(database.DB).AddUserOnline(online)
	if err != nil {
		return err
	}
	return nil
}
func DeleteUserOnlineService(socketid string) error {
	return NewUserOnlineRepoImpl(database.DB).DeleteUserOnline(socketid)
}
