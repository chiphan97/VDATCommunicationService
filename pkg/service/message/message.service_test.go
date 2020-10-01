package message

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"log"
	"testing"
)

func TestGetMessagesByGroupAndUserService(t *testing.T) {
	database.Connect()
	ms, err := GetMessagesByGroupAndUserService(1, "thien")
	if err != nil {
		log.Fatal(err)
	}
	println(ms)
}
func TestAddMessageService(t *testing.T) {
	database.Connect()
	payload := PayLoad{
		SubjectSender: "ffb63922-8f99-46ba-9648-d07f3ac14757",
		Content:       "mới thêm",
		IdGroup:       1,
	}
	err := AddMessageService(payload)
	if err != nil {
		t.Fatal(err)
	}

}
func TestLoadMessageHistoryService(t *testing.T) {
	database.Connect()

	dtos, err := LoadMessageHistoryService(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dtos)

}
