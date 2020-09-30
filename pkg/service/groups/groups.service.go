package groups

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
)

// tao chat 1 1 neu chua co, neu co r tra lai
func GetGroupByOwnerAndUserService(groupPayload PayLoad, owner string) ([]Dto, error) {
	groupdtos := make([]Dto, 0)

	group := groupPayload.ConvertToModel()
	groups, err := NewRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(owner, group.Users[0])
	if err != nil {
		return nil, err
	} else {
		if len(groups) <= 0 {
			group.UserCreate = owner
			group, err := NewRepoImpl(database.DB).AddGroupType(group)
			if err != nil {
				return nil, err
			}
			group.Users = append(group.Users, group.UserCreate)
			err = NewRepoImpl(database.DB).AddGroupUser(group.Users, int(group.ID))
			if err != nil {
				return nil, err
			}
			groups, err = NewRepoImpl(database.DB).GetGroupByOwnerAndUserAndTypeOne(group.UserCreate, group.Users[0])
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

func GetGroupByPatientService(user string) ([]Dto, error) {
	groupdtos := make([]Dto, 0)
	groups, err := NewRepoImpl(database.DB).GetGroupByUser(user)
	if err != nil {
		return nil, err
	}
	pubGroups, err := NewRepoImpl(database.DB).GetGroupByPrivateAndUser(false, user)
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
func GetGroupByDoctorService(user string) ([]Dto, error) {
	groupdtos := make([]Dto, 0)
	groups, err := NewRepoImpl(database.DB).GetGroupByUser(user)
	if err != nil {
		return nil, err
	}
	pubGroups, err := NewRepoImpl(database.DB).GetGroupByType(MANY, user)
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
func AddGroupManyService(groupPayLoad PayLoad, owner string) (Dto, error) {
	var groupdto Dto
	group := groupPayLoad.ConvertToModel()
	group.UserCreate = owner
	group.Users = append(group.Users, owner)
	group, err := NewRepoImpl(database.DB).AddGroupType(group)
	if err != nil {
		return groupdto, err
	}
	groupdto = group.ConvertToDTO()
	err = NewRepoImpl(database.DB).AddGroupUser(group.Users, int(group.ID))
	if err != nil {
		return groupdto, err
	}
	return groupdto, nil
}

func UpdateGroupService(groupsPayLoad PayLoad, idGroup int) (Dto, error) {
	var groupdto Dto
	group := groupsPayLoad.ConvertToModel()
	group.ID = uint(idGroup)
	group, err := NewRepoImpl(database.DB).UpdateGroup(group)
	if err != nil {
		return groupdto, err
	}
	groupdto = group.ConvertToDTO()
	return groupdto, err
}
func DeleteGroupService(idgroup int) error {
	return NewRepoImpl(database.DB).DeleteGroup(idgroup)
}
func CheckRoleOwnerInGroupService(owner string, idgroup int) (bool, error) {
	return NewRepoImpl(database.DB).GetOwnerByGroupAndOwner(owner, idgroup)
}
func AddUserInGroupService(userIds []string, groupId int) error {
	return NewRepoImpl(database.DB).AddGroupUser(userIds, groupId)
}
func DeleteUserInGroupService(userIds []string, groupId int) error {
	return NewRepoImpl(database.DB).DeleteGroupUser(userIds, groupId)
}
func GetListUserByGroupService(groupId int) ([]userdetail.Dto, error) {
	dtos := make([]userdetail.Dto, 0)
	users, err := NewRepoImpl(database.DB).GetListUserByGroup(groupId)
	if err != nil {
		return dtos, err
	}
	for _, u := range users {
		user := u.ConvertToDto()
		dtos = append(dtos, user)
	}
	return dtos, nil
}