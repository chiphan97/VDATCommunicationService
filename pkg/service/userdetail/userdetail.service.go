package userdetail

import "gitlab.com/vdat/mcsvc/chat/pkg/database"

func AddUserDetailService(payload Payload) error {
	detail := payload.convertToModel()
	err := NewRepoImpl(database.DB).AddUserDetail(detail)
	if err != nil {
		return err
	}
	return nil
}
func GetUserDetailByIDService(id string) (Dto, error) {
	var dto Dto
	detail, err := NewRepoImpl(database.DB).GetUserDetailById(id)
	if err != nil {
		return dto, err
	}
	dto = detail.ConvertToDto()
	return dto, nil
}
func GetListUserDetailService(fil string) ([]Dto, error) {
	dtos := make([]Dto, 0)
	userdetails, err := NewRepoImpl(database.DB).GetListUser(fil)
	if err != nil {
		return dtos, err
	}
	for _, detail := range userdetails {
		dto := detail.ConvertToDto()
		dtos = append(dtos, dto)
	}
	return dtos, nil
}
