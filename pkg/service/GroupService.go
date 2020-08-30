package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

// tao chat 1 1 neu chua co, neu co r tra lai
func GetGroupByOwnerAndUserService(owner string, user string) ([]model.Groups, error) {
	groups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(owner, user)
	if err != nil {
		return nil, err
	} else {
		if len(groups) <= 0 {
			userOnline, err := impl.NewUserOnlineRepoImpl(database.DB).GetUserOnline(user)
			if err != nil {
				return nil, err
			}
			group, err := impl.NewGroupRepoImpl(database.DB).AddGroupType(owner, userOnline.Username, model.ONE, true)
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
	if err != nil {
		return nil, err
	}
	pubGroups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByPrivate(false)
	if err != nil {
		return nil, err
	}
	if len(pubGroups) > 0 {
		for _, g := range pubGroups {
			groups = append(groups, g)
		}
	}
	return groups, nil
}

func AddGroupManyService(owner string, nameGroup string, private bool, listUser []string) (model.Groups, error) {
	group, err := impl.NewGroupRepoImpl(database.DB).AddGroupType(owner, nameGroup, model.MANY, private)
	if err != nil {
		return group, err
	}
	err = impl.NewGroupUserRepoImpl(database.DB).AddGroupUser(listUser, int(group.ID))
	if err != nil {
		return group, err
	}
	return group, nil
}
