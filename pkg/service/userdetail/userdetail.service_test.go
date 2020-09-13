package userdetail

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"testing"
)

func TestAddUserDetailService(t *testing.T) {
	database.Connect()
	p := Payload{
		ID:       "test",
		Username: "test",
		First:    "test",
		Last:     "test",
		Role:     ADMIN,
	}
	err := AddUserDetailService(p)
	if err != nil {
		t.Fatal(err)
		return
	}
}
func TestGetUserDetailByIDService(t *testing.T) {
	database.Connect()
	dto, err := GetUserDetailByIDService("test")
	if err != nil {
		t.Fatal(err)
		return
	}
	if dto != (Dto{}) {
		t.Log(dto)
	}
}
func TestGetListUserDetailService(t *testing.T) {
	database.Connect()
	dtos, err := GetListUserDetailService("a")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(dtos)
}
