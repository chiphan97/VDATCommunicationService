package handler

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func RegisterGroupApi() {
	http.HandleFunc("/create-group-one", AuthenMiddleJWT(CreateGroupTypeOneApi))
	http.HandleFunc("/groups", AuthenMiddleJWT(GroupApi))
}
func CreateGroupTypeOneApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		owner := r.URL.Query()["owner"][0]
		user := r.URL.Query()["user"][0]
		groups, err := service.GetGroupByOwnerAndUserService(owner, user)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	} else {
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
func GroupApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user := r.URL.Query()["user"][0]
		groups, err := service.GetGroupByUserService(user)
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(groups))
	} else {
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
