package groups

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"log"
	"testing"
)

func TestGetListUserOnlineAndOffByGroupService(t *testing.T) {
	database.Connect()
	dtos, err := GetListUserOnlineAndOffByGroupService(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dtos)
}
