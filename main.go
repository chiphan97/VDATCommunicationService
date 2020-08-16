package main

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository/impl"
	"log"
)

func main() {
	database.Connect()
	//r := mux.NewRouter()

	//textDB
	chatboxs, err := impl.NewChatBoxRepoImpl(database.DB).GetChatBoxs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chatboxs)

	messages, err1 := impl.NewMessageRepoImpl(database.DB).GetMessages()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(messages)

	//err := http.ListenAndServe("8000",r)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("server listen in port 8000")
}
