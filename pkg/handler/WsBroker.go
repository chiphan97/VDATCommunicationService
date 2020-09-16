package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type WsBroker struct {
	// Registered clients.
	Clients map[*WsClient]bool

	// Inbound message from the clients.
	Inbound chan WsMessage

	// Outbound message that need to send to clients.
	Outbound chan WsMessage

	// Register requests from the clients.
	Register chan *WsClient

	// Unregister requests from clients.
	Unregister chan *WsClient

	MessageRepository []*WsMessage

	Conn *websocket.Conn
}

func (broker *WsBroker) run() {
	// polling "new" message from repository
	// and send to outbound channel to send to clients
	// finally, marked message that sent to outbound channel as "done"
	go func() {
		for {
			for idx, message := range broker.MessageRepository {
				if message.Status != "done" {
					select {
					case broker.Outbound <- *message:
					default:
						close(broker.Outbound)
					}

					broker.MessageRepository[idx].Status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-broker.Register:
			broker.Clients[client] = true
			//add in database when client on
			//_ = useronline.AddUserOnlineService(client.User)

		case client := <-broker.Unregister:
			if _, ok := broker.Clients[client]; ok {
				//delete in database when client off
				//_ = useronline.DeleteUserOnlineService(client.User.SocketID)
				delete(broker.Clients, client)
				close(client.Send)
			}
		case message := <-broker.Inbound:
			broker.MessageRepository = append(broker.MessageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(broker.MessageRepository))
		case message := <-broker.Outbound:
			fmt.Println("send")
			for client := range broker.Clients {
				msg, _ := json.Marshal(message)
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(broker.Clients, client)
				}
			}
		}

	}
}
