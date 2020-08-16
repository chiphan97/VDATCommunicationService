package service

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"testing"
)

func TestGetMessageByChatBoxService(t *testing.T) {
	database.Connect()
	messages,err := GetMessageByChatBoxService(2)
	if err!= nil || len(messages) <0{
		t.Error(err)
		return
	}
	fmt.Println(messages)
}
func TestInsertMessagesService(t *testing.T) {
	database.Connect()
	err := InsertMessagesService()
	if err!= nil{
		t.Error(err)
		return
	}
	fmt.Println("insert success")
}