package handler

import (
	"encoding/json"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
	"strconv"
)

func RegisterGroupUsersApi() {
	http.HandleFunc("/group-user", AuthenMiddleJWT(GroupUserApi))
}

func GroupUserApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//API thêm thành viên vào 1 nhóm
	case http.MethodPost:
		var group model.Groups
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			utils.ResponseErr(w, http.StatusForbidden)
			return
		}
		err = service.AddUserInGroup(group.ListUser, int(group.ID))
		if err != nil {
			utils.ResponseErr(w, http.StatusRequestTimeout)
		}
		utils.ResponseOk(w, "Success")
	case http.MethodDelete: //API thành viên rời khỏi nhóm
		userid := r.URL.Query()["user_id"]
		groupstr := r.URL.Query()["group_id"]
		if len(userid) > 0 && len(groupstr) > 0 {
			//API thành viên rời khỏi nhóm
			users := []string{userid[0]}
			groupid, _ := strconv.Atoi(groupstr[0])
			err := service.DeleteUserInGroup(users, groupid)
			if err != nil {
				utils.ResponseErr(w, http.StatusRequestTimeout)
				return
			}
			utils.ResponseOk(w, "Success")
		} else {
			//API xóa 1 hoac nhieu thành viên trogn nhóm
			var group model.Groups
			err := json.NewDecoder(r.Body).Decode(&group)
			if err != nil {
				utils.ResponseErr(w, http.StatusForbidden)
				return
			}
			err = service.DeleteUserInGroup(group.ListUser, int(group.ID))
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
