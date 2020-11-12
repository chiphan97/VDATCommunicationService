package dchat

import (
	"encoding/json"
	"fmt"
	message_service "gitlab.com/vdat/mcsvc/chat/pkg/service/message"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
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

// Chat WebSocket godoc
// @Summary Chat websocket
// @Tags dchat
// @Param socketId path string true "socketId to know client"
// @Param token query string true "token to be join chat"
// @Param message body dchat.Message true "Works based on field event type (read NOTE)"
// @Accept  json
// @Produce  json
// @Success 200 {object} dchat.Message
// @Router /message/{socketId} [post]
// @Description NOTE
// @Description Event For Send Message
// @Description
// @Description "type":"subcribe_group" - to open the group the person has joined
// @Description
// @Description "type":"send_text" - to send text from current client to users in that group
// @Description
// @Description "type":"load_old_mess" - to load continues history message in group
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
						//close(b.Outbound)
					}

					b.MessageRepository[idx].Data.Status = "done"
				}
			}
			time.Sleep(1 * time.Millisecond)
			//time.Sleep(200 * time.Millisecond)
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
			switch message.TypeEvent {
			case SEND:
				userOn, err := useronline.GetListUSerOnlineByGroupService(message.Data.GroupId)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(userOn)
				payload := message_service.PayLoad{
					SubjectSender: message.Client,
					Content:       message.Data.Body,
					IdGroup:       message.Data.GroupId,
				}
				newMess, err := message_service.AddMessageService(payload)
				if err != nil {
					log.Fatal(err)
				}
				for client := range b.Clients {
					for _, u := range userOn {
						if u.UserID == client.UserId && u.SocketID == client.SocketId {
							message.Data.Id = int(newMess.ID)
							message.Data.CreatedAt = newMess.CreatedAt
							message.Data.UpdatedAt = newMess.UpdatedAt
							message.Data.Sender = newMess.SubjectSender
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
			case RELY:
				fmt.Println(message.Data.Id)
				fmt.Println(message.Data.Body)

				userOn, err := useronline.GetListUSerOnlineByGroupService(message.Data.GroupId)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(userOn)
				payload := message_service.PayLoad{
					SubjectSender: message.Client,
					Content:       message.Data.Body,
					IdGroup:       message.Data.GroupId,
					ID:            message.Data.Id,
				}
				newMess, err := message_service.AddRelyService(payload)
				if err != nil {
					log.Fatal(err)
				}
				for client := range b.Clients {
					for _, u := range userOn {
						if u.UserID == client.UserId && u.SocketID == client.SocketId {
							message.Data.Id = int(newMess.ID)
							message.Data.CreatedAt = newMess.CreatedAt
							message.Data.UpdatedAt = newMess.UpdatedAt
							message.Data.Sender = newMess.SubjectSender
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

			case LOADCHILDMESS:
				fmt.Println(message.Data.Id)
			case SUBCRIBE:
				historys, err := message_service.LoadMessageHistoryService(message.Data.GroupId)

				if err != nil {
					log.Println(err)
				}
				var msg []byte
				for _, h := range historys {
					for client := range b.Clients {
						if client.UserId == message.Client && client.SocketId == message.Data.SocketID {

							mess := Message{
								TypeEvent: SUBCRIBE,
								Data: Data{
									Id:        int(h.ID),
									GroupId:   message.Data.GroupId,
									Body:      h.Content,
									Sender:    h.SubjectSender,
									CreatedAt: h.CreatedAt,
									UpdatedAt: h.UpdatedAt,
								},
							}

							msg, _ = json.Marshal(mess)
							select {
							case client.Send <- msg:
							default:
								close(client.Send)
								delete(b.Clients, client)
							}
						}

					}
				}
			case LOADOLDMESS:
				continueHistory, err := message_service.LoadContinueMessageHistoryService(message.Data.IdContinueOldMess, message.Data.GroupId)
				if err != nil {
					log.Println(err)
				}
				var msg []byte
				fmt.Println(continueHistory)
				for _, h := range continueHistory {
					for client := range b.Clients {
						if client.UserId == message.Client && client.SocketId == message.Data.SocketID {
							mess := Message{
								TypeEvent: LOADOLDMESS,
								Data: Data{
									Id:        int(h.ID),
									GroupId:   message.Data.GroupId,
									Body:      h.Content,
									Sender:    h.SubjectSender,
									CreatedAt: h.CreatedAt,
									UpdatedAt: h.UpdatedAt,
								},
							}
							msg, _ = json.Marshal(mess)
							select {
							case client.Send <- msg:
							default:
								close(client.Send)
								delete(b.Clients, client)
							}
						}

					}
				}
			default:

			}

		}

	}
}
