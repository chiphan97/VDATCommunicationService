package useronline

type Repo interface {
	AddUserOnline(online UserOnline) error
	DeleteUserOnline(socketid string) error
	GetUserOnlineBySocketIdAndHostId(socketID string, hostname string) (UserOnline, error)
}
