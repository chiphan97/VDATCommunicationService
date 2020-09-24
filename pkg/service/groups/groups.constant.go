package groups

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	CREATE = "group:create_group"
	UPDATE = "group:update_group"
	DELETE = "group:delete_group"
	LIST   = "group:list_group"

	LISTMEMBER   = "group:member:list_member"
	ADDMEMBER    = "group:member:add_member"
	MEMBEROUT    = "group:member:member_out_group"
	DELETEMEMBER = "group:member:delete_member"
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// Maximum message size allowed from peer.
	MaxMessageSize = 512
)

var (
	Newline = []byte{'\n'}
	Space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
