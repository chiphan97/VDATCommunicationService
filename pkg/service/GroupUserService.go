package service

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
)

func AddUserInGroup(userIds []string, groupId int) error {
	return impl.NewGroupUserRepoImpl(database.DB).AddGroupUser(userIds, groupId)
}
func DeleteUserInGroup(userIds []string, groupId int) error {
	return impl.NewGroupUserRepoImpl(database.DB).DeleteGroupUser(userIds, groupId)
}
