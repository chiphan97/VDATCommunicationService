package message

import (
	"context"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
)

func GetMessagesByGroupAndUserService(idGroup int, subUser string) ([]Messages, error) {
	return NewRepoImpl(database.DB).GetMessagesByGroupAndUser(idGroup, subUser)
}
func AddMessageService(payload PayLoad) (Dto, error) {
	var dto Dto
	model := payload.convertToModel()
	m, err := NewRepoImpl(database.DB).InsertMessage(model)
	if err != nil {
		return dto, err
	}
	dto = m.convertToDTO()
	return dto, nil
}

func AddRelyService(payload PayLoad) (Dto, error) {
	var dto Dto
	model := payload.convertToModel()
	m, err := NewRepoImpl(database.DB).InsertRely(model)
	if err != nil {
		return dto, err
	}
	dto = m.convertToDTO()
	return dto, nil
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

func LoadChildMessageService(idGroup int, parentId int) ([]Dto, error) {
	dtos := make([]Dto, 0)
	model, err := NewRepoImpl(database.DB).GetChilMessByParentId(idGroup, parentId)
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
func DeleteMessageService(idGroup int) error {
	return RepoImpl(database.DB).DeleteMessageByGroup(idGroup, context.Background())
}
