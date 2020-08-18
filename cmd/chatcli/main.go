package main

import (
	"bufio"
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
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

		messagePayload := model.MessagePayload{
			IdChat:  1,
			Message: messageInput,
		}

		if _, err := ws.Write(utils.ResponseWithByte(messagePayload)); err != nil {
			log.Fatal(err)
		}
	}

}
