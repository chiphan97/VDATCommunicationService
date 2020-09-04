package controller

import (
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func RegisterUserOnlineApi(r *mux.Router) {
	r.HandleFunc("/api/v1/users", AuthenMiddleJWT(GetUserOnlineApi)).Methods(http.MethodGet, http.MethodOptions)

}

//API tìm kiếm người dùng filtter
func GetUserOnlineApi(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	fil := r.URL.Query()["keyword"]
	users, err := service.GetListUSerOnlineService(fil[0])
	if err != nil {
		utils.ResponseErr(w, http.StatusNotFound)
	}
	w.Write(utils.ResponseWithByte(users))
	//switch r.Method {
	//case http.MethodGet:
	//
	//default:
	//	utils.ResponseErr(w, http.StatusBadRequest)
	//}
}
