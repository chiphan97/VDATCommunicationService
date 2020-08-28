package handler

import (
	"github.com/gorilla/websocket"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
)

type WsClient struct {
	Broker *WsBroker

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	User model.UserOnline
}
