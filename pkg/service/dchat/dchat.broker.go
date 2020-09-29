package dchat

import (
	"encoding/json"
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/groups"

	"log"
	"time"
)

type Broker struct {
	// 1 group va nhieu client connect toi
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Inbound chan Message

	// Outbound messages that need to Send to Clients.
	Outbound chan Message

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client

	MessageRepository []*Message
}

var Wsbroker = &Broker{
	Clients:    make(map[*Client]bool),
	Inbound:    make(chan Message),
	Outbound:   make(chan Message),
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
				if m.Data.Status != "done" {
					select {
					case b.Outbound <- *m:
					default:
						close(b.Outbound)
					}

					b.MessageRepository[idx].Data.Status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-b.Register:
			// khi client dang nhap thi broker se lay group dua tren idgroup cua client, neu chua co thi tao group vao broker

			b.Clients[client] = true

			fmt.Println("client " + client.UserId + " is connected")
		case client := <-b.Unregister:
			// khi client dang xuat khoi group, delete client khoi group
			if _, ok := b.Clients[client]; ok {
				//delete in database when client off
				//_ = useronline.DeleteUserOnlineService(client.User.SocketID)
				delete(b.Clients, client)
				close(client.Send)
			}
			//khi co tin nhan dc gui di , message se duoc day vao inbound va day vao MessageRepository
		case message := <-b.Inbound:

			b.MessageRepository = append(b.MessageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(b.MessageRepository))

		case message := <-b.Outbound:
			fmt.Println("Send")
			switch message.TypeEvent {
			case SEND:
				dtos, err := groups.GetListUserOnlineAndOffByGroupService(message.Data.GroupId)
				if err != nil {
					log.Fatal(err)
				}
				for client := range b.Clients {
					for _, u := range dtos {
						if u.ID == client.UserId && u.Status == groups.USERON {
							msg, _ := json.Marshal(message)
							select {
							case client.Send <- msg:
							default:
								close(client.Send)
								delete(b.Clients, client)
							}
						}
					}

				}
			case SUBCRIBE:
			default:

			}

		}

	}
}
