package message

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
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
