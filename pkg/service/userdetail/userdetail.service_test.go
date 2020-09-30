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
	dto, err := GetUserDetailByIDService("893a4692-63bb-4919-80d9-aece678c0422")
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
func TestUpdateUserDetailservice(t *testing.T) {
	database.Connect()

	payload := Payload{
		ID:       "893a4692-63bb-4919-80d9-aece678c0422",
		Username: "test",
		First:    "test",
		Last:     "test",
		Role:     DOCTOR,
	}
	err := UpdateUserDetailservice(payload)
	if err != nil {
		t.Fatal(err)
		return
	}

}

//func TestCheckUserDetailService(t *testing.T) {
//	database.Connect()
//	//dto,err := CheckUserDetailService("893a4692-63bb-4919-80d9-aece678c0422")
//	dto,err := CheckUserDetailService("893a46")
//	if err!= nil{
//		t.Fatal(err)
//	}
//	t.Log(dto)
//}
