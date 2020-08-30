package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type UserOnlineRepo interface {
	GetListUSerOnline(userHide string) ([]model.UserOnline, error)
	AddUserOnline(online model.UserOnline) error
	DeleteUserOnline(socketid string) error
}
