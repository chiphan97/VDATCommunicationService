package useronline

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"testing"
)

func TestAddUserOnlineService(t *testing.T) {
	database.Connect()
	p := Payload{
		HostName: "test host",
		SocketID: "test socket",
		UserID:   "test",
	}
	dp := userdetail.Payload{
		ID:       "test",
		Username: "ko co thi them",
		First:    "ko co thi them",
		Last:     "ko co thi them",
		Role:     userdetail.ADMIN,
	}
	err := AddUserOnlineService(p, dp)
	if err != nil {
		t.Log(err)
	}
}
func TestDeleteUserOnlineService(t *testing.T) {
	database.Connect()
	err := DeleteUserOnlineService("test socket")
	if err != nil {
		t.Log(err)
	}
}
