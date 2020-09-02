package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
)

func RegisterGroupApi(r *mux.Router) {
	r.HandleFunc("/api/v1/groups", AuthenMiddleJWT(GroupApi))
	r.HandleFunc("/api/v1/groups", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		writer.Header().Set("Access-Control-Max-Age", "86400")
	})
	r.HandleFunc("/api/v1/groups/{idGroup}", AuthenMiddleJWT(GroupApi))
	r.HandleFunc("/api/v1/groups/{idGroup}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		writer.Header().Set("Access-Control-Max-Age", "86400")
	})
	r.HandleFunc("/api/v1/groups/{idGroup}/members", AuthenMiddleJWT(GroupUserApi))
	r.HandleFunc("/api/v1/groups/{idGroup}/members", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		writer.Header().Set("Access-Control-Max-Age", "86400")
	})
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", AuthenMiddleJWT(GroupUserApi))
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		writer.Header().Set("Access-Control-Max-Age", "86400")
	})
}
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

	case http.MethodPut: //API update group info
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
	case http.MethodDelete: //API xóa hội thoại
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
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}

}
func GroupUserApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//API thêm thành viên vào 1 nhóm
	case http.MethodPatch:
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
	case http.MethodDelete:
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
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
