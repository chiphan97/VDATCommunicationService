package userdetail

import (
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"net/http"
)

func RegisterUserApi(r *mux.Router) {
	//r.HandleFunc("/api/v1/user", auth.AuthenMiddleJWT(GetUserApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/user/info", auth.AuthenMiddleJWT(CheckUserDetailApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/user/info", auth.AuthenMiddleJWT(UserLogOutApi)).Methods(http.MethodDelete, http.MethodOptions)
}

//API tìm kiếm người dùng filtter
//func GetUserApi(w http.ResponseWriter, r *http.Request) {
//	cors.SetupResponse(&w, r)
//
//	fil := r.URL.Query()["keyword"]
//	users, err := GetListUserDetailService()
//	if err != nil {
//		utils.ResponseErr(w, http.StatusNotFound)
//	}
//	w.Write(utils.ResponseWithByte(users))
//}
func CheckUserDetailApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	payload, err := JWTparseUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto, err := CheckUserDetailService(payload)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dto.HostName = utils.GetLocalIP()
	dto.SocketID = utils.ArraySocketId[0]
	utils.ArraySocketId = utils.DeleteItemInArray(utils.ArraySocketId)
	utils.WriteLines(utils.ArraySocketId, "socketid.data")

	uo := useronline.Payload{
		HostName: dto.HostName,
		SocketID: dto.SocketID,
		UserID:   payload.ID,
	}
	err = useronline.AddUserOnlineService(uo)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(utils.ResponseWithByte(dto))

	// check user he thong neu login chua ton tai thong tin trong he thong thi ghi vao database

}
func UserLogOutApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	hostname := r.URL.Query()["hostName"][0]
	socketID := r.URL.Query()["socketId"][0]
	err := useronline.DeleteUserOnlineService(socketID, hostname)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ArraySocketId = utils.RestoreItemArray(utils.ArraySocketId, socketID)
	utils.WriteLines(utils.ArraySocketId, "socketid.data")
	w.Write(utils.ResponseWithByte(true))
}
