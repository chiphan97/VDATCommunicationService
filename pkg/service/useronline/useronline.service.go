package useronline

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
)

//func GetListUSerOnlineService(fil string) ([]Dto, error) {
//	users := make([]Dto, 0)
//	userOnlines, err := NewRepoImpl(database.DB).GetListUSerOnline(fil)
//	if err != nil {
//		return users, err
//	}
//	for _, userOnline := range userOnlines {
//		user := Dto{
//			UserID:   userOnline.UserID,
//			UserName: userOnline.UserName,
//			First:    userOnline.First,
//			Last:     userOnline.Last,
//		}
//		users = append(users, user)
//	}
//	return users, nil
//}
func AddUserOnlineService(payload Payload) error {
	check, err := NewRepoImpl(database.DB).GetUserOnlineBySocketIdAndHostId(payload.SocketID, payload.HostName)
	if err != nil {
		return err
	}
	if check == (UserOnline{}) {
		online := payload.convertToModel()

		err = NewRepoImpl(database.DB).AddUserOnline(online)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUserOnlineService(socketid string) error {
	return NewRepoImpl(database.DB).DeleteUserOnline(socketid)
}
