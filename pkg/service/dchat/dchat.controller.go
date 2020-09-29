package dchat

import (
	_ "bytes"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"log"
	"net/http"
)

func ChatHandlr(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	// authenticate

	conn, err := WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	owner := auth.JWTparseOwner2(r.URL.Query()["token"][0])

	client := &Client{UserId: owner, Broker: Wsbroker, Conn: conn, Send: make(chan []byte, 256)}
	client.Broker.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.WritePump()
	go client.ReadPump()

}
