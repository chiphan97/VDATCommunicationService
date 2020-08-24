package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/handler"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"golang.org/x/net/websocket"
	"log"
	"net"
	"net/http"
	"os"
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

// https://github.com/gorilla/websocket/blob/master/examples/chat/main.go
func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	go metrics()

	fmt.Println("starting")

	database.Connect()
	http.HandleFunc("/", serveHome)
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.HandleFunc("/test", service.TestHandler)
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", handler.HandleConnections)
	go handler.HandleMessages()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func metrics() {
	// The debug listener mounts the http.DefaultServeMux, and serves up
	// stuff like the Prometheus metrics route, the Go debug and profiling
	// routes, and so on.
	debugListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}

	defer debugListener.Close()

	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	fmt.Println(http.Serve(debugListener, http.DefaultServeMux))
}
