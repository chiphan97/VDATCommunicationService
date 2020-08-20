package handler

import (
	"github.com/gorilla/websocket"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)    // connected clients
var broadcast = make(chan model.MessagePayload) // broadcast channel
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true
	clients[ws] = true

	for {
		var messagePayload model.MessagePayload
		err := ws.ReadJSON(&messagePayload)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- messagePayload
		if len(messagePayload.Message) > 0 {
			chatbox, err := service.CreateChatBox(messagePayload.SenderID, messagePayload.ReceiverID)
			if err != nil {
				log.Fatal(err)
			}
			message := model.MessageModel{
				Content: messagePayload.Message,
				IdChat:  chatbox.ID,
			}
			err = service.InsertMessagesService(message)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Read in a new message as JSON and map it to a Message object
		//err := ws.ReadJSON(&msg)
		//if err != nil {
		//	log.Printf("error: %v", err)
		//	delete(clients, ws)
		//	break
		//}
		//// Send the newly received message to the broadcast channel
		//broadcast <- msg
	}
}

func HandleMessages() {
	for {
		messagePayload := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(messagePayload)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
