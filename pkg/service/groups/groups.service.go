package groups

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
)

// tao chat 1 1 neu chua co, neu co r tra lai
func GetGroupByOwnerAndUserService(groupPayload GroupsPayLoad, owner string) ([]GroupsDTO, error) {
	groupdtos := make([]GroupsDTO, 0)

	group := groupPayload.ConvertToModel()
	groups, err := NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(owner, group.Users[0])
	if err != nil {
		return nil, err
	} else {
		if len(groups) <= 0 {
			group.UserCreate = owner
			group, err := NewGroupRepoImpl(database.DB).AddGroupType(group)
			if err != nil {
				return nil, err
			}
			group.Users = append(group.Users, group.UserCreate)
			err = NewGroupRepoImpl(database.DB).AddGroupUser(group.Users, int(group.ID))
			if err != nil {
				return nil, err
			}
			groups, err = NewGroupRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(group.UserCreate, group.Users[0])
			if err != nil {
				return nil, err
			}
			for _, group := range groups {
				groupdto := group.ConvertToDTO()
				groupdtos = append(groupdtos, groupdto)
			}
			return groupdtos, nil
		} else {
			for _, group := range groups {
				groupdto := group.ConvertToDTO()
				groupdtos = append(groupdtos, groupdto)
			}
			return groupdtos, nil
		}
	}
}

func GetGroupByUserService(user string) ([]GroupsDTO, error) {
	groupdtos := make([]GroupsDTO, 0)
	groups, err := NewGroupRepoImpl(database.DB).GetGroupByUser(user)
	if err != nil {
		return nil, err
	}
	pubGroups, err := NewGroupRepoImpl(database.DB).GetGroupByPrivateAndUser(false, user)
	if err != nil {
		return nil, err
	}
	if len(pubGroups) > 0 {
		for _, g := range pubGroups {
			groups = append(groups, g)
		}
	}
	for _, group := range groups {
		groupdto := group.ConvertToDTO()
		groupdtos = append(groupdtos, groupdto)
	}
	return groupdtos, nil
}

func AddGroupManyService(groupPayLoad GroupsPayLoad, owner string) (GroupsDTO, error) {
	var groupdto GroupsDTO
	group := groupPayLoad.ConvertToModel()
	group.UserCreate = owner
	group.Users = append(group.Users, owner)
	group, err := NewGroupRepoImpl(database.DB).AddGroupType(group)
	if err != nil {
		return groupdto, err
	}
	groupdto = group.ConvertToDTO()
	err = NewGroupRepoImpl(database.DB).AddGroupUser(group.Users, int(group.ID))
	if err != nil {
		return groupdto, err
	}
	return groupdto, nil
}

func UpdateGroupService(groupsPayLoad GroupsPayLoad, idGroup int) (GroupsDTO, error) {
	var groupdto GroupsDTO
	group := groupsPayLoad.ConvertToModel()
	group.ID = uint(idGroup)
	group, err := NewGroupRepoImpl(database.DB).UpdateGroup(group)
	if err != nil {
		return groupdto, err
	}
	groupdto = group.ConvertToDTO()
	return groupdto, err
}
func DeleteGroupService(idgroup int) error {
	return NewGroupRepoImpl(database.DB).DeleteGroup(idgroup)
}
func CheckRoleOwnerInGroupService(owner string, idgroup int) (bool, error) {
	return NewGroupRepoImpl(database.DB).GetOwnerByGroupAndOwner(owner, idgroup)
}
func AddUserInGroupService(userIds []string, groupId int) error {
	return NewGroupRepoImpl(database.DB).AddGroupUser(userIds, groupId)
}
func DeleteUserInGroupService(userIds []string, groupId int) error {
	return NewGroupRepoImpl(database.DB).DeleteGroupUser(userIds, groupId)
}
func GetListUserByGroupService(groupId int) ([]useronline.Dto, error) {
	users := make([]useronline.Dto, 0)
	userOnlines, err := NewGroupRepoImpl(database.DB).GetListUserByGroup(groupId)
	if err != nil {
		return users, err
	}
	for _, useronline := range userOnlines {
		user := useronline.ConvertToDto()
		users = append(users, user)
	}
	return users, nil
}
