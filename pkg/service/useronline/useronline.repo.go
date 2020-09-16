package useronline

type Repo interface {
	AddUserOnline(online UserOnline) error
	DeleteUserOnline(socketid string) error
	GetUserOnlineById(userId string) (UserOnline, error)
}
