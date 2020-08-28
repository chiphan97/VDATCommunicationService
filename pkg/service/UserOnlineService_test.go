package service

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"testing"
)

func TestAddUserOnlineService(t *testing.T) {
	database.Connect()
	user := model.UserOnline{
		HostName: "test",
		SocketID: "test",
		UserID:   "test",
		Username: "test",
		First:    "test",
		Last:     "test",
	}
	err := AddUserOnlineService(user)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("success")
	}
}
