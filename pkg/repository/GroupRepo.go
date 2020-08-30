package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type GroupRepo interface {
	//chat one - one
	GetGroupByOwnerAndUserAndTypeOne(owner string, user string) ([]model.Groups, error)
	GetGroupByUser(user string) ([]model.Groups, error)
	GetGroupByPrivate(private bool) ([]model.Groups, error)
	AddGroupType(owner string, name string, typ string, private bool) (model.Groups, error)
}
