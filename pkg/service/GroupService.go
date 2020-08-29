package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func GetGroupByOwnerAndUserService(owner string, user string) ([]model.Groups, error) {
	groups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(owner, user)
	if err != nil {
		return nil, err
	} else {
		if len(groups) <= 0 {
			group, err := impl.NewGroupRepoImpl(database.DB).AddGroupTypeONE(owner)
			if err != nil {
				return nil, err
			}
			users := []string{owner, user}
			err = impl.NewGroupUserRepoImpl(database.DB).AddGroupUser(users, int(group.ID))
			if err != nil {
				return nil, err
			}
			groups, err = impl.NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(owner, user)
			if err != nil {
				return nil, err
			}
			return groups, nil
		} else {
			return groups, nil
		}
	}
}
func GetGroupByUserService(user string) ([]model.Groups, error) {
	groups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByUser(user)
	if err != nil && len(groups) <= 0 {
		return nil, err
	}
	return groups, nil
}
