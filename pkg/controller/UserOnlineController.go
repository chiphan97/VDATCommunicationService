package controller

import (
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func RegisterUserOnlineApi(r *mux.Router) {
	r.HandleFunc("/api/v1/users", AuthenMiddleJWT(UserOnlineApi))
	r.HandleFunc("/api/v1/users", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		writer.Header().Set("Access-Control-Max-Age", "86400")
	})

}

//API tìm kiếm người dùng filtter
func UserOnlineApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fil := r.URL.Query()["keyword"]
		userOns, err := service.GetListUSerOnlineService(fil[0])
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(userOns))
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
