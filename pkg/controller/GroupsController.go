package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
)

func RegisterGroupApi(r *mux.Router) {
	r.HandleFunc("/api/v1/groups", AuthenMiddleJWT(GetGroupApi)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/groups", AuthenMiddleJWT(PostGroupApi)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/groups/{idGroup}", AuthenMiddleJWT(PutGroupApi)).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/groups/{idGroup}", AuthenMiddleJWT(DeleteGroupApi)).Methods(http.MethodPut)

	r.HandleFunc("/api/v1/groups/{idGroup}/members", AuthenMiddleJWT(PatchGroupUserApi)).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", AuthenMiddleJWT(DeleteGroupUserApi)).Methods(http.MethodDelete)

}

func GetGroupApi(w http.ResponseWriter, r *http.Request) {
	//API load danh sách group (public, private)
	SetupResponse(&w, r)

	user := JWTparseOwner(r.Header.Get("Authorization"))
	groups, err := service.GetGroupByUserService(user)
	if err != nil {
		utils.ResponseErr(w, http.StatusNotFound)
	}
	w.Write(utils.ResponseWithByte(groups))
}
func PostGroupApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	var group model.Groups
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := JWTparseOwner(r.Header.Get("Authorization"))
	group.UserCreate = owner
	if group.Type == model.ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
		groups, err := service.GetGroupByOwnerAndUserService(group)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	} else { //Tạo hội thoại n- n
		group.ListUser = append(group.ListUser, owner)
		group, err = service.AddGroupManyService(group)
		if err != nil {
			utils.ResponseErr(w, http.StatusBadRequest)
			return
		}
		w.Write(utils.ResponseWithByte(group))
	}
}
func PutGroupApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	var group model.Groups
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	group.ID = uint(id)
	newgroup, err := service.UpdateGroupService(group)
	if err != nil {
		utils.ResponseErr(w, http.StatusRequestTimeout)
		return
	}
	w.Write(utils.ResponseWithByte(newgroup))
}
func DeleteGroupApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := JWTparseOwner(r.Header.Get("Authorization"))
	check, err := service.CheckRoleOwnerInGroupService(owner, id)
	if !check {
		utils.ResponseErr(w, http.StatusMethodNotAllowed)
		return
	} else {
		err = service.DeleteGroupService(id)
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
			return
		}
		utils.ResponseOk(w, "delete success")
	}
}

//func GroupApi(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//
//	case http.MethodPost:
//
//
//	case http.MethodPut: //API update group info
//
//	case http.MethodDelete: //API xóa hội thoại
//
//	default:
//		utils.ResponseErr(w, http.StatusBadRequest)
//	}
//
//}
//API thêm thành viên vào 1 nhóm
func PatchGroupUserApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	var group model.Groups
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	group.ID = uint(id)
	err = service.AddUserInGroup(group.ListUser, int(group.ID))
	if err != nil {
		utils.ResponseErr(w, http.StatusRequestTimeout)
	}
	utils.ResponseOk(w, "Success")
}
func DeleteGroupUserApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	userid := params["userId"]
	owner := JWTparseOwner(r.Header.Get("Authorization"))
	check, err := service.CheckRoleOwnerInGroupService(owner, groupID)
	if err != nil {
		utils.ResponseErr(w, http.StatusForbidden)
		return
	}
	if !check {
		users := []string{owner}
		err := service.DeleteUserInGroup(users, groupID)
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
			return
		}
		utils.ResponseOk(w, "Success")
	} else {
		//xoa thanh vien trong nhom
		users := []string{userid}
		err := service.DeleteUserInGroup(users, groupID)
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
			return
		}
		utils.ResponseOk(w, "Success")
	}
}

//func GroupUserApi(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//
//	case http.MethodPatch:
//
//	case http.MethodDelete:
//
//	default:
//		utils.ResponseErr(w, http.StatusBadRequest)
//	}
//}
