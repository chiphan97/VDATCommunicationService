package useronline

type Repo interface {
	GetListUSerOnline() ([]UserOnline, error)
	AddUserOnline(online UserOnline) error
	DeleteUserOnline(socketid string) error
	GetUserOnlineBySocketIdAndHostId(socketID string, hostname string) (UserOnline, error)
}
