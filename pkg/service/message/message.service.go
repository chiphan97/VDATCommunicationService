package message

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
)

func GetMessagesByGroupAndUserService(idGroup int, subUser string) ([]Messages, error) {
	return RepoImpl(database.DB).GetMessagesByGroupAndUser(idGroup, subUser)
}
func AddMessageService(payload PayLoad) error {
	model := payload.convertToModel()
	return RepoImpl(database.DB).InsertMessage(model)
}
func LoadMessageHistoryService(idGroup int) ([]Dto, error) {
	dtos := make([]Dto, 0)
	model, err := RepoImpl(database.DB).GetMessagesByGroup(idGroup)
	if err != nil {
		return dtos, err
	}
	for _, m := range model {
		dto := m.convertToDTO()
		dtos = append(dtos, dto)
	}
	return dtos, nil

}
