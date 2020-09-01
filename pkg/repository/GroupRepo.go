package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type GroupRepo interface {
	//chat one - one
	GetGroupByOwnerAndUserAndTypeOne(owner string, user string) ([]model.Groups, error)
	GetGroupByUser(user string) ([]model.Groups, error)
	GetGroupByPrivateAndUser(private bool, user string) ([]model.Groups, error)
	AddGroupType(group model.Groups) (model.Groups, error)
	UpdateGroup(group model.Groups) (model.Groups, error)
	DeleteGroup(idGourp int) error
	GetOwnerByGroupAndOwner(owner string, groupId int) (bool, error)
}
