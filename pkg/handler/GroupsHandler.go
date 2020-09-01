package handler

import (
	"encoding/json"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

func RegisterGroupApi() {
	//http.HandleFunc("/create-group-one", AuthenMiddleJWT(CreateGroupTypeOneApi))
	http.HandleFunc("/api/v1/groups/", AuthenMiddleJWT(GroupApi))

}

// api Tạo hội thoại 1 - 1 (nhóm bí mật)
//func CreateGroupTypeOneApi(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		var user model.UserOnline
//		err := json.NewDecoder(r.Body).Decode(&user)
//		if err != nil {
//			utils.ResponseErr(w, http.StatusForbidden)
//			return
//		}
//		owner := JWTparseOwner(r.Header.Get("Authorization"))
//		groups, err := service.GetGroupByOwnerAndUserService(owner, user.UserID)
//		if err != nil {
//			utils.ResponseErr(w, http.StatusNotFound)
//		}
//		w.Write(utils.ResponseWithByte(groups))
//	} else {
//		utils.ResponseErr(w, http.StatusBadRequest)
//	}
//}

func GroupApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//API load danh sách group (public, private)
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
		group.UserCreate = owner
		if group.TypeGroup == model.ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
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

	case http.MethodPut: //API update group info
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/groups/"))
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
	case http.MethodDelete: //API xóa hội thoại
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/groups/"))
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
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}

}
