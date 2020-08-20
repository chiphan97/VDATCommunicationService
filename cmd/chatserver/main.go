package main

import (
	"encoding/json"
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {
	receivedtext := make([]byte, 128)

	n, err := ws.Read(receivedtext)

	if err != nil {
		fmt.Printf("Received: %d bytes\n", n)
	}

	s := string(receivedtext[:n])

	var messagePayload model.MessagePayload

	err = json.Unmarshal([]byte(s), &messagePayload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messagePayload)

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
func chatHandler() {

}
func main() {
	database.Connect()
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.HandleFunc("/test", service.TestHandler)
	//http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
