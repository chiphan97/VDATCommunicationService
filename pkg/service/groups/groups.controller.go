package groups

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"

	"log"
	"net/http"
	"strconv"
)

func GroupManageWs(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	// authenticate
	param := r.URL.Query()["token"][0]
	fmt.Println(param)
	owner, err := auth.JWTparseOwnerGroupWs(param)
	if err != nil {
		utils.ResponseErr(w, http.StatusUnauthorized)
		return
	}
	fmt.Println(owner)
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{UserID: owner, Broker: Wsbroker, Conn: conn, Send: make(chan []byte, 256)}
	client.Broker.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.WritePump()
	go client.ReadPump()
}
func RegisterGroupApi(r *mux.Router) {
	r.HandleFunc("/api/v1/groups", auth.AuthenMiddleJWT(GetListGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups", auth.AuthenMiddleJWT(CreateGroupApi)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", auth.AuthenMiddleJWT(UpdateInfoGroupApi)).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}", auth.AuthenMiddleJWT(DeleteGroupApi)).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(GetListUserByGroupApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(AddUserInGroupApi)).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members", auth.AuthenMiddleJWT(UserOutGroupApi)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/groups/{idGroup}/members/{userId}", auth.AuthenMiddleJWT(DeleteGroupUserApi)).Methods(http.MethodDelete, http.MethodOptions)

}

//API load danh sách groups theo patient hoac theo doctor
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
func CreateGroupApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	var groupPayLoad PayLoad
	err := json.NewDecoder(r.Body).Decode(&groupPayLoad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))

	if groupPayLoad.Type == ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
		groupsDto, err := GetGroupByOwnerAndUserService(groupPayLoad, owner)
		if err != nil {
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
		utils.ResponseOk(w, true)
	}
}

//API xoa thành viên khoi 1 nhóm va chi owner moi dc xoa
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
		utils.ResponseOk(w, true)
	}
}

//API user outgroup nhung owner ko dc out
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
		utils.ResponseOk(w, true)
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
