package main

import (
	"encoding/json"
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"golang.org/x/net/websocket"
	"net/http"
	"strings"
)

func echoHandler(ws *websocket.Conn) {
	receivedtext := make([]byte, 128)

	n, err := ws.Read(receivedtext)

	if err != nil {
		fmt.Printf("Received: %d bytes\n", n)
	}

	s := string(receivedtext[:n])

	decoder := json.NewDecoder(strings.NewReader(s))

	messagePayload := model.MessagePayload{}

	_ = decoder.Decode(&messagePayload)

	if len(s) > 0 {
		fmt.Printf("Received: %d bytes: %s", n, messagePayload)
	}
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	//http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
