package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

// tao chat 1 1 neu chua co, neu co r tra lai
func GetGroupByOwnerAndUserService(group model.Groups) ([]model.Groups, error) {
	groups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(group.UserCreate, group.ListUser[0])
	if err != nil {
		return nil, err
	} else {
		if len(groups) <= 0 {

			group, err := impl.NewGroupRepoImpl(database.DB).AddGroupType(group)
			if err != nil {
				return nil, err
			}
			group.ListUser = append(group.ListUser, group.UserCreate)
			err = impl.NewGroupUserRepoImpl(database.DB).AddGroupUser(group.ListUser, int(group.ID))
			if err != nil {
				return nil, err
			}
			groups, err = impl.NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(group.UserCreate, group.ListUser[0])
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
	pubGroups, err := impl.NewGroupRepoImpl(database.DB).GetGroupByPrivateAndUser(false, user)
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

func AddGroupManyService(group model.Groups) (model.Groups, error) {
	group, err := impl.NewGroupRepoImpl(database.DB).AddGroupType(group)
	if err != nil {
		return group, err
	}
	err = impl.NewGroupUserRepoImpl(database.DB).AddGroupUser(group.ListUser, int(group.ID))
	if err != nil {
		return group, err
	}
	return group, nil
}

func UpdateGroupService(group model.Groups) (model.Groups, error) {
	return impl.NewGroupRepoImpl(database.DB).UpdateGroup(group)
}
func DeleteGroupService(idgroup int) error {
	return impl.NewGroupRepoImpl(database.DB).DeleteGroup(idgroup)
}
func CheckRoleOwnerInGroupService(owner string, idgroup int) (bool, error) {
	return impl.NewGroupRepoImpl(database.DB).GetOwnerByGroupAndOwner(owner, idgroup)
}
