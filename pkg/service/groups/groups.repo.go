package groups

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type GroupRepo interface {
	//chat one - one
	GetGroupByOwnerAndUserAndTypeOne(owner string, user string) ([]Groups, error)
	GetGroupByUser(user string) ([]Groups, error)
	GetGroupByPrivateAndUser(private bool, user string) ([]Groups, error)
	AddGroupType(group Groups) (Groups, error)
	UpdateGroup(group Groups) (Groups, error)
	DeleteGroup(idGourp int) error
	GetOwnerByGroupAndOwner(owner string, groupId int) (bool, error)
	GetListUserByGroup(idGourp int) ([]model.UserOnline, error)
	AddGroupUser(users []string, idgroup int) error
	DeleteGroupUser(users []string, idgroup int) error
}
