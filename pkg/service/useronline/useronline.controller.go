package useronline

import (
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func RegisterUserOnlineApi(r *mux.Router) {
	r.HandleFunc("/api/v1/users", auth.AuthenMiddleJWT(GetUserOnlineApi)).Methods(http.MethodGet, http.MethodOptions)
}

//API tìm kiếm người dùng filtter
func GetUserOnlineApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	fil := r.URL.Query()["keyword"]
	users, err := GetListUSerOnlineService(fil[0])
	if err != nil {
		utils.ResponseErr(w, http.StatusNotFound)
	}
	w.Write(utils.ResponseWithByte(users))
}
