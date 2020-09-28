package dchat

import (
	"encoding/json"
	"fmt"
	"time"
)

type Broker struct {
	// 1 group va nhieu client connect toi
	Groups map[int]map[*Client]bool

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
	Groups:     make(map[int]map[*Client]bool),
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
				if m.status != "done" {
					select {
					case b.Outbound <- *m:
					default:
						close(b.Outbound)
					}

					b.MessageRepository[idx].status = "done"
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case client := <-b.Register:
			// khi client dang nhap thi broker se lay group dua tren idgroup cua client, neu chua co thi tao group vao broker
			groupConnect := b.Groups[client.GroupID]
			if groupConnect == nil {
				groupConnect = make(map[*Client]bool)
				b.Groups[client.GroupID] = groupConnect
			}
			b.Groups[client.GroupID][client] = true
			fmt.Print("123")
			for group := range b.Groups {
				fmt.Print(group)
			}

			fmt.Println("client " + client.User.UserID + " is connected")
		case client := <-b.Unregister:
			// khi client dang xuat khoi group, delete client khoi group
			groupConnect := b.Groups[client.GroupID]
			if groupConnect != nil {
				if _, ok := groupConnect[client]; ok {
					delete(groupConnect, client)
					close(client.Send)
					//neu group do ko con client nao thi xoa luon group khoi he thong
					if len(groupConnect) == 0 {
						delete(b.Groups, client.GroupID)
					}
				}
			}
			//khi co tin nhan dc gui di , message se duoc day vao inbound va day vao MessageRepository
		case message := <-b.Inbound:

			b.MessageRepository = append(b.MessageRepository, &message)
			fmt.Printf("%+v, %d\n", message, len(b.MessageRepository))

		case message := <-b.Outbound:
			fmt.Println("Send")
			clients := b.Groups[message.GroupId]
			for client := range clients {
				msg, _ := json.Marshal(message)
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(b.Groups[message.GroupId], client)
				}
			}
		}

	}
}
