package useronline

type UserOnlineRepo interface {
	GetListUSerOnline(filter string) ([]UserOnline, error)
	AddUserOnline(online UserOnline) error
	DeleteUserOnline(socketid string) error
	GetUserOnline(userId string) (UserOnline, error)
}
