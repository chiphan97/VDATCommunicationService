package service

import (
	"fmt"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"testing"
)

func TestGetMessageByChatBoxService(t *testing.T) {
	database.Connect()
	messages, err := GetMessageByChatBoxService(1)
	if err != nil || len(messages) < 0 {
		t.Error(err)
		return
	}
	fmt.Println(messages)
}
func BenchmarkGetMessageByChatBoxService(b *testing.B) {
	database.Connect()

	messages, err := GetMessageByChatBoxService(2)
	if err != nil || len(messages) < 0 {
		b.Error(err)
	}

}
func TestInsertMessagesService(t *testing.T) {
	database.Connect()

	messageModel := model.MessageModel{
		Content: "ethsohasod",
		IdChat:  1,
	}
	err := InsertMessagesService(messageModel)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("insert success")
}
func BenchmarkInsertMessagesService(b *testing.B) {
	database.Connect()
	messageModel := model.MessageModel{
		Content: "ethsohasod",
		IdChat:  1,
	}
	err := InsertMessagesService(messageModel)
	if err != nil {
		b.Error(err)
	}
	//fmt.Println("insert success")
}

//func TestUpdateMessageService(t *testing.T) {
//	database.Connect()
//	time := time.Now()
//	abstract := model.AbstractModel{
//		ID:        6,
//		UpdatedAt: &time,
//	}
//	messageModel := model.MessageModel{
//		AbstractModel: abstract,
//		Content:       "moi up date",
//	}
//	err := UpdateMessageService(messageModel)
//	if err != nil {
//		t.Error(err)
//	}
//}
//func BenchmarkUpdateMessageService(b *testing.B) {
//	database.Connect()
//	time := time.Now()
//	abstract := model.AbstractModel{
//		ID:        7,
//		UpdatedAt: &time,
//	}
//	messageModel := model.MessageModel{
//		AbstractModel: abstract,
//		Content:       "moi up date",
//	}
//	err := UpdateMessageService(messageModel)
//	if err != nil {
//		b.Error(err)
//	}
//
//}
//func TestDeleteMessageService(t *testing.T) {
//	database.Connect()
//	err := DeleteMessageService(10)
//	if err != nil {
//		t.Error(err)
//	}
//}
//func BenchmarkDeleteMessageService(b *testing.B) {
//	database.Connect()
//	err := DeleteMessageService(15)
//	if err != nil {
//		b.Error(err)
//	}
//}
