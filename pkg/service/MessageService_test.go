package service

import (

	"testing"
)

func TestGetMessageByChatBoxService(t *testing.T) {
	messages,err := GetMessageByChatBoxService(2)
	if err!= nil || len(messages) <0{
		t.Error(err)
	}
	//t.Log(messages)
}