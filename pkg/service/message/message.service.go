package message

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
)

func GetMessagesByGroupAndUserService(idGroup int, subUser string) ([]Messages, error) {
	return NewRepoImpl(database.DB).GetMessagesByGroupAndUser(idGroup, subUser)
}
func AddMessageService(payload PayLoad) (int, error) {
	model := payload.convertToModel()
	return NewRepoImpl(database.DB).InsertMessage(model)
}
func LoadMessageHistoryService(idGroup int) ([]Dto, error) {
	dtos := make([]Dto, 0)
	model, err := NewRepoImpl(database.DB).GetMessagesByGroup(idGroup)
	if err != nil {
		return dtos, err
	}
	for _, m := range model {
		dto := m.convertToDTO()
		dtos = append(dtos, dto)
	}
	return dtos, nil

}
func LoadContinueMessageHistoryService(idMessage int, idGroup int) ([]Dto, error) {
	dtos := make([]Dto, 0)
	model, err := NewRepoImpl(database.DB).GetContinueMessageByIdAndGroup(idMessage, idGroup)
	if err != nil {
		return dtos, err
	}
	for _, m := range model {
		dto := m.convertToDTO()
		dtos = append(dtos, dto)
	}
	return dtos, nil

}
