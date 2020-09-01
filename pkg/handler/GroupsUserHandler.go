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
	http.HandleFunc("/user-out-group", AuthenMiddleJWT(UserOutGroupApi))
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

func UserOutGroupApi(w http.ResponseWriter, r *http.Request) {

	//groupstr := r.URL.Query()["group_id"][0]
	//
	//groupid, _ := strconv.Atoi(groupstr)
	//check,err := service.CheckRoleOwnerInGroupService(userid[0],groupid)
	//if err!=nil{
	//	utils.ResponseErr(w, http.StatusForbidden)
	//	return
	//}
	//if check{
	//	utils.ResponseErr(w, http.StatusNotAcceptable)
	//	return
	//}else{
	//
	//}

}
