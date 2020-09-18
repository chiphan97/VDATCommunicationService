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
	r.HandleFunc("/api/v1/user", auth.AuthenMiddleJWT(GetUserApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/user/info", auth.AuthenMiddleJWT(CheckUserDetailApi)).Methods(http.MethodGet, http.MethodOptions)
}

//API tìm kiếm người dùng filtter
func GetUserApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	fil := r.URL.Query()["keyword"]
	users, err := GetListUserDetailService(fil[0])
	if err != nil {
		utils.ResponseErr(w, http.StatusNotFound)
	}
	w.Write(utils.ResponseWithByte(users))
}
func CheckUserDetailApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	payload, err := JWTparseUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload.Role = DOCTOR

	dto, err := CheckUserDetailService(payload)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uo := useronline.Payload{
		HostName: r.URL.RawPath,
		SocketID: payload.ID,
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
