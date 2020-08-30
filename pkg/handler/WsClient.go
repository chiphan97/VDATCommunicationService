package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"log"
	"time"
)

type WsClient struct {
	Broker *WsBroker

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	User model.UserOnline
}

func (client *WsClient) ReadPump() {
	defer func() {
		client.Broker.Unregister <- client
		_ = client.Conn.Close()
	}()
	client.Conn.SetReadLimit(MaxMessageSize)
	_ = client.Conn.SetReadDeadline(time.Now().Add(PongWait))
	client.Conn.SetPongHandler(func(string) error { _ = client.Conn.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, Newline, Space, -1))

		var messageJSON WsMessage
		_ = json.Unmarshal(message, &messageJSON)
		messageJSON.From = client.User.UserID

		client.Broker.Inbound <- messageJSON
	}
}

func (client *WsClient) CheckUserOnlinePump(userHide string) {
	defer func() {
		client.Broker.Unregister <- client
		_ = client.Conn.Close()
	}()

	for {
		usersOnline, _ := service.GetListUSerOnlineService()
		message := WsMessage{
			From:   "VDAT-SERVICE",
			To:     nil,
			Body:   usersOnline,
			Status: "",
		}

		client.Broker.Inbound <- message
		time.Sleep(10000 * time.Millisecond)
	}
}

func (client *WsClient) WritePump() {
	ticker := time.NewTicker(PingPeriod)

	defer func() {
		ticker.Stop()
		_ = client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				// The broker closed the channel.
				_ = client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(Newline)
				_, _ = w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
