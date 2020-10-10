package groups

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"

	"net/http"
	"strconv"
)

//func GroupManageWs(w http.ResponseWriter, r *http.Request) {
//	cors.SetupResponse(&w, r)
//	// authenticate
//	param := r.URL.Query()["token"][0]
//	fmt.Println(param)
//	owner, err := auth.JWTparseOwnerGroupWs(param)
//	if err != nil {
//		utils.ResponseErr(w, http.StatusUnauthorized)
//		return
//	}
//	fmt.Println(owner)
//	conn, err := Upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	client := &Client{UserID: owner, Broker: Wsbroker, Conn: conn, Send: make(chan []byte, 256)}
//	client.Broker.Register <- client
//
//	// Allow collection of memory referenced by the caller by doing all work in
//	// new goroutines.
//
//	go client.WritePump()
//	go client.ReadPump()
//}
func RegisterGroupApi(r *mux.Router) {
	r.HandleFunc("/api/v1/groups", auth.AuthenMiddleJWT(GetListGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups", auth.AuthenMiddleJWT(CreateGroupApi)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", auth.AuthenMiddleJWT(UpdateInfoGroupApi)).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", auth.AuthenMiddleJWT(DeleteGroupApi)).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(GetListUserOnlineByGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(AddUserInGroupApi)).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(UserOutGroupApi)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", auth.AuthenMiddleJWT(DeleteGroupUserApi)).Methods(http.MethodDelete, http.MethodOptions)

}

//API load danh sách groups theo patient hoac theo doctor

// GetListGroupApi godoc
// @Summary Get all groups
// @Description Get all groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {array} Dto
// @Router /api/v1/groups [get]
func GetListGroupApi(w http.ResponseWriter, r *http.Request) {

	cors.SetupResponse(&w, r)

	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := userdetail.GetUserDetailByIDService(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	if check.Role == userdetail.PATIENT {
		groups, err := GetGroupByPatientService(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groups))
	} else {
		groups, err := GetGroupByDoctorService(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groups))
	}

}

// api tao group n - n chi doctor dc tao va tao chat 1 1

// CreateOrder godoc
// @Summary Create a new groups
// @Description create a new groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Param groupPayLoad body PayLoad true "Create groups"
// @Success 200 {object} Dto
// @Router /api/v1/groups [post]
func CreateGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	var groupPayLoad PayLoad
	err := json.NewDecoder(r.Body).Decode(&groupPayLoad)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))

	if groupPayLoad.Type == ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
		groupsDto, err := GetGroupByOwnerAndUserService(groupPayLoad, owner)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groupsDto))
	} else { //Tạo hội thoại n- n
		check, err := userdetail.GetUserDetailByIDService(owner)
		if check.Role == userdetail.PATIENT {
			w.WriteHeader(http.StatusForbidden)
			utils.ResponseErr(w, http.StatusForbidden)
			return
		}
		groupDto, err := AddGroupManyService(groupPayLoad, owner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.ResponseErr(w, http.StatusInternalServerError)
			return
		}
		w.Write(utils.ResponseWithByte(groupDto))
	}

}

// api update ten nhom

// Updategroups godoc
// @Summary Update group by groupId
// @Description Update the group corresponding to the input groupId
// @Tags groups
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID of the group to be updated"
// @Param groupPayLoad body PayLoad true "update groups"
// @Success 200 {object} Dto
// @Router /api/v1/groups/{idGroup} [put]
func UpdateInfoGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := CheckRoleOwnerInGroupService(owner, id)
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
		var groupPayLoad PayLoad
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

}

// DeleteOrder godoc
// @Summary Delete group identified by the given idGroup
// @Description Delete the group corresponding to the input idGroup
// @Tags groups
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID of the group to be updated"
// @Success 204 "No Content"
// @Router /api/v1/groups/{idGroup} [delete]
func DeleteGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
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

//API thêm thành viên vào 1 nhóm va chi owner moi dc them

// add user to group godoc
// @Summary add user to group
// @Description add user to group
// @Tags groupUser
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID of the group to be updated"
// @Param groupPayLoad body PayLoad true "add user to group"
// @Success 200 {object} boolean
// @Router /api/v1/groups/{idGroup}/members [patch]
func AddUserInGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
	check, err := CheckRoleOwnerInGroupService(owner, id)
	if !check {
		w.WriteHeader(http.StatusForbidden)
		utils.ResponseErr(w, http.StatusForbidden)
		return
	} else {
		var groupPayload PayLoad
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
		result := utils.ResponseBool{Result: true}
		w.Write(utils.ResponseWithByte(result))
	}
}

//API xoa thành viên khoi 1 nhóm va chi owner moi dc xoa

// delete user to group by admin godoc
// @Summary delete user to group by admin
// @Description delete user to group by admin
// @Tags groupUser
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID group"
// @Param userId path int true "ID user want delete"
// @Success 200
// @Router /api/v1/groups/{idGroup}/members/{userId} [delete]
func DeleteGroupUserApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	userid := params["userId"]
	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
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
		result := utils.ResponseBool{Result: true}
		w.Write(utils.ResponseWithByte(result))
	}
}

//API user outgroup nhung owner ko dc out

// delete user to group godoc
// @Summary delete user to group
// @Description delete user to group
// @Tags groupUser
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID of the group to be add user"
// @Success 200
// @Router /api/v1/groups/{idGroup}/members [delete]
func UserOutGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
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
		result := utils.ResponseBool{Result: true}
		w.Write(utils.ResponseWithByte(result))
	}

}
func GetListUserByGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
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

// GetListMemberGroupApi godoc
// @Summary Get all member groups
// @Description Get all member groups
// @Tags groupUser
// @Accept  json
// @Produce  json
// @Param idGroup path int true "ID of the group to be updated"
// @Success 200 {array} []userdetail.Dto
// @Router /api/v1/groups/{idGroup}/members [get]
func GetListUserOnlineByGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["idGroup"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	users, err := GetListUserOnlineAndOffByGroupService(groupID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	w.Write(utils.ResponseWithByte(users))
}
