package userdetail

type Repo interface {
	GetListUser(filter string) ([]UserDetail, error)
	AddUserDetail(detail UserDetail) error
	GetUserDetailById(id string) (UserDetail, error)
}
