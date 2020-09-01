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

func RegisterGroupUsersApi() {
	http.HandleFunc("/api/v1/groups/:groupId/members/:userid", AuthenMiddleJWT(GroupUserApi))

}

func GroupUserApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//API thêm thành viên vào 1 nhóm
	case http.MethodPatch:
		idgroupstr := strings.TrimPrefix(r.URL.Path, "/api/v1/groups/")
		id, _ := strconv.Atoi(strings.TrimRight(idgroupstr, "/members/"))
		var group model.Groups
		err := json.NewDecoder(r.Body).Decode(&group)
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
		userid := r.URL.Query()["user_id"]
		groupstr := r.URL.Query()["group_id"]
		groupID, _ := strconv.Atoi(groupstr[0])
		owner := JWTparseOwner(r.Header.Get("Authorization"))
		check, err := service.CheckRoleOwnerInGroupService(owner, groupID)
		if err != nil {
			utils.ResponseErr(w, http.StatusForbidden)
			return
		}
		if len(userid) > 0 {
			if !check {
				utils.ResponseErr(w, http.StatusUnauthorized)
				return
			} else {
				//xoa thanh vien trong nhom
				err := service.DeleteUserInGroup(userid, groupID)
				if err != nil {
					utils.ResponseErr(w, http.StatusRequestTimeout)
					return
				}
				utils.ResponseOk(w, "Success")
			}
		} else { //API thành viên rời khỏi nhóm
			if check {
				utils.ResponseErr(w, http.StatusNotAcceptable)
				return
			} else {
				users := []string{owner}
				err := service.DeleteUserInGroup(users, groupID)
				if err != nil {
					utils.ResponseErr(w, http.StatusRequestTimeout)
					return
				}
				utils.ResponseOk(w, "Success")
			}
		}

	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
