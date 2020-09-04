package groups

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/controller"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
)

func RegisterGroupApi(r *mux.Router) {
	r.HandleFunc("/api/v1/groups", controller.AuthenMiddleJWT(GetGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups", controller.AuthenMiddleJWT(PostGroupApi)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", controller.AuthenMiddleJWT(PutGroupApi)).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", controller.AuthenMiddleJWT(DeleteGroupApi)).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/groups/{idGroup}/members", controller.AuthenMiddleJWT(GetListUserByGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", controller.AuthenMiddleJWT(PatchGroupUserApi)).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", controller.AuthenMiddleJWT(UserOutGroupApi)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", controller.AuthenMiddleJWT(DeleteGroupUserApi)).Methods(http.MethodDelete, http.MethodOptions)

}

func GetGroupApi(w http.ResponseWriter, r *http.Request) {
	//API load danh sách groups (public, private)
	controller.SetupResponse(&w, r)

	user := controller.JWTparseOwner(r.Header.Get("Authorization"))
	groups, err := GetGroupByUserService(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	w.Write(utils.ResponseWithByte(groups))
}
func PostGroupApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	var groupPayLoad GroupsPayLoad
	err := json.NewDecoder(r.Body).Decode(&groupPayLoad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	owner := controller.JWTparseOwner(r.Header.Get("Authorization"))

	if groupPayLoad.Type == ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
		groupsDto, err := GetGroupByOwnerAndUserService(groupPayLoad, owner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groupsDto))
	} else { //Tạo hội thoại n- n

		groupDto, err := AddGroupManyService(groupPayLoad, owner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groupDto))
	}
}
func PutGroupApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	var groupPayLoad GroupsPayLoad
	err = json.NewDecoder(r.Body).Decode(&groupPayLoad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	newgroup, err := UpdateGroupService(groupPayLoad, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	w.Write(utils.ResponseWithByte(newgroup))
}
func DeleteGroupApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := controller.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := CheckRoleOwnerInGroupService(owner, id)
	if !check {
		w.WriteHeader(http.StatusForbidden)
		utils.ResponseErr(w, http.StatusForbidden)
		return
	} else {
		err = DeleteGroupService(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		utils.ResponseOk(w, true)
	}
}

//API thêm thành viên vào 1 nhóm
func PatchGroupUserApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	var groupPayload GroupsPayLoad
	err = json.NewDecoder(r.Body).Decode(&groupPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	err = AddUserInGroupService(groupPayload.Users, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
	}
	utils.ResponseOk(w, true)
}
func DeleteGroupUserApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	userid := params["userId"]
	owner := controller.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := CheckRoleOwnerInGroupService(owner, groupID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	if !check {
		w.WriteHeader(http.StatusForbidden)
		utils.ResponseErr(w, http.StatusForbidden)
		return
	} else {
		//xoa thanh vien trong nhom
		users := []string{userid}
		err := DeleteUserInGroupService(users, groupID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		utils.ResponseOk(w, true)
	}
}
func UserOutGroupApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := controller.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := CheckRoleOwnerInGroupService(owner, groupID)

	if check {
		w.WriteHeader(http.StatusForbidden)
		utils.ResponseErr(w, http.StatusForbidden)
		return
	} else {
		//
		users := []string{owner}
		err := DeleteUserInGroupService(users, groupID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		utils.ResponseOk(w, true)
	}

}
func GetListUserByGroupApi(w http.ResponseWriter, r *http.Request) {
	controller.SetupResponse(&w, r)
	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	users, err := GetListUserByGroupService(groupID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	w.Write(utils.ResponseWithByte(users))
}
