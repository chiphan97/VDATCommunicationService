package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {
	receivedtext := make([]byte, 128)

	n, err := ws.Read(receivedtext)

	if err != nil {
		fmt.Printf("Received: %d bytes\n", n)
	}

	s := string(receivedtext[:n])
	if len(s) > 0 {
		fmt.Printf("Received: %d bytes: %s", n, s)
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
