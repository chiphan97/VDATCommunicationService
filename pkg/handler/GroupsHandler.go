package handler

import (
	"encoding/json"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
)

func RegisterGroupApi() {
	http.HandleFunc("/create-group-one", AuthenMiddleJWT(CreateGroupTypeOneApi))
	http.HandleFunc("/groups", AuthenMiddleJWT(GroupApi))
}

// api Tạo hội thoại 1 - 1 (nhóm bí mật)
func CreateGroupTypeOneApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user := r.URL.Query()["user"][0]
		owner := JWTparseOwner(r.Header.Get("Authorization"))
		groups, err := service.GetGroupByOwnerAndUserService(owner, user)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	} else {
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}

func GroupApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: //API load danh sách group (public, private)
		user := JWTparseOwner(r.Header.Get("Authorization"))
		groups, err := service.GetGroupByUserService(user)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	case http.MethodPost: //Tạo hội thoại n- n
		var group model.Groups
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			utils.ResponseErr(w, http.StatusForbidden)
			return
		}
		owner := JWTparseOwner(r.Header.Get("Authorization"))
		group.ListUser = append(group.ListUser, owner)
		group, err = service.AddGroupManyService(owner, group.NameGroup, group.Private, group.ListUser)
		if err != nil {
			utils.ResponseErr(w, http.StatusBadRequest)
			return
		}
		w.Write(utils.ResponseWithByte(group))
	case http.MethodPut: //API update group info
		var group model.Groups
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			utils.ResponseErr(w, http.StatusForbidden)
			return
		}
		newgroup, err := service.UpdateGroupService(group)
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
			return
		}
		w.Write(utils.ResponseWithByte(newgroup))
	case http.MethodDelete: //API xóa hội thoại
		idgroupstr := r.URL.Query()["group_id"][0]
		idgroup, _ := strconv.Atoi(idgroupstr)
		err := service.DeleteGroupService(idgroup)
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
			return
		}
		utils.ResponseOk(w, "delete success")
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}

}
