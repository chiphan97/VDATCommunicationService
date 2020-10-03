package groups

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"time"
)

type Broker struct {
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Inbound chan MessageEvent

	// Outbound messages that need to Send to Clients.
	Outbound chan MessageEvent

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client

	MessageRepository []*MessageEvent
}

var Wsbroker = &Broker{
	Clients:    make(map[*Client]bool),
	Inbound:    make(chan MessageEvent),
	Outbound:   make(chan MessageEvent),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

// luong nay dc khoi tao khi chay truong trinh
// Broker la noi tiep nhan client khi client do mo connect
//nhan cac message dc gui len cua client va tra message dc gui di den client nhan
func (b *Broker) Run() {
	// polling "new" message from repository
	// and Send to Outbound channel to Send to Clients
	// finally, marked message that sent to Outbound channel as "done"
	go func() {
		for {
			for idx, m := range b.MessageRepository {
				if m.Status != "done" {
					select {
					case b.Outbound <- *m:
					default:
						close(b.Outbound)
					}

					b.MessageRepository[idx].Status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-b.Register:
			b.Clients[client] = true
			fmt.Println(client.UserID + "is connect")
		case client := <-b.Unregister:
			if _, ok := b.Clients[client]; ok {
				delete(b.Clients, client)
				close(client.Send)
			}
		case message := <-b.Inbound:
			b.MessageRepository = append(b.MessageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(b.MessageRepository))
		case message := <-b.Outbound:
			fmt.Println("send")
			msg := HandLerEvent(message)

			users, err := GetListUserByGroupService(int(message.IdGroup))

			if err != nil {
				fmt.Println(err)
			}
			u, err := userdetail.GetUserDetailByIDService(message.UserId)
			if err != nil {
				fmt.Println(err)
			}
			users = append(users, u)

			for _, user := range users {
				for client := range b.Clients {
					if user.ID == client.UserID {
						select {
						case client.Send <- msg:
						default:
							close(client.Send)
							delete(b.Clients, client)
						}
					}
				}
			}

		}
	}
}
