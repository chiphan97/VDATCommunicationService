package handler

import (
	"encoding/json"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
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

//Tạo hội thoại n- n
//API load danh sách group (public, private)
func GroupApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		user := JWTparseOwner(r.Header.Get("Authorization"))
		groups, err := service.GetGroupByUserService(user)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	case http.MethodPost:
		var group model.Groups
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			utils.ResponseErr(w, http.StatusBadRequest)
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
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}

}
