package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type GroupRepo interface {
	//chat one - one
	GetGroupByOwnerAndUserAndTypeOne(owner string, user string) ([]model.Groups, error)
	GetGroupByUser(user string) ([]model.Groups, error)
	AddGroupTypeONE(owner string) (model.Groups, error)
}
