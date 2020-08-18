package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

func main() {
	for {
		origin := "http://localhost/"
		url := "ws://localhost:8080/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("message:")
		messageReader := bufio.NewReader(os.Stdin)
		messageInput, _ := messageReader.ReadString('\n')

		messageInput = messageInput[:len(messageInput)-1]
		if _, err := ws.Write([]byte(messageInput + "\n")); err != nil {
			log.Fatal(err)
		}
	}

}
