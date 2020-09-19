package dchat

import (
	_ "bytes"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
	"log"
	"net/http"
	"strconv"
)

func ChatHandlr(w http.ResponseWriter, r *http.Request) {
	// authenticate

	conn, err := WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	owner := auth.JWTparseOwner(r.Header.Get("Authorization"))
	idgroupstr := mux.Vars(r)["idgroup"]
	idGroup, _ := strconv.Atoi(idgroupstr)

	user := useronline.UserOnline{
		HostName: r.URL.Hostname(),
		SocketID: owner,
		UserID:   owner,
	}

	client := &Client{User: user, Broker: Wsbroker, Conn: conn, Send: make(chan []byte, 256), GroupID: idGroup}
	client.Broker.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.WritePump()
	go client.ReadPump()

}
