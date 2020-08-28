package handler

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"net/http"
)

func UsersOnlineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		usersOnline, err := service.GetListUSerOnlineService()
		if err != nil && len(usersOnline) <= 0 {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(usersOnline))
	}
}
