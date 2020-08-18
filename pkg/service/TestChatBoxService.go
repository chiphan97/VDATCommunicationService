package service

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"testing"
)

func TestCreateChatBox(t * testing.T)  {
	database.Connect()

	receiverId := "anonymousUser"
	chatBox, err := CreateChatBox(receiverId)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(chatBox)
	}
}

func TestFindChatBoxBySender(t * testing.T)  {
	database.Connect()

	senderId := "anonymousUser"
	chatBox, err := FindChatBoxBySender(senderId)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(chatBox)
	}
}

func TestFindChatBoxById(t * testing.T)  {
	database.Connect()

	senderId := 1
	chatBox, err := FindChatBoxById(uint(senderId))
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(chatBox)
	}
}

func TestDeleteChatBoxById(t * testing.T)  {
	database.Connect()

	senderId := 1
	chatBox, err := DeleteChatBox(uint(senderId))
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(chatBox)
	}
}

