package userdetail

import (
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func RegisterUserApi(r *mux.Router) {
	r.HandleFunc("/api/v1/users", auth.AuthenMiddleJWT(GetUserApi)).Methods(http.MethodGet, http.MethodOptions)
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
