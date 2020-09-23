package userdetail

type Repo interface {
	GetListUser(filter string) ([]UserDetail, error)
	AddUserDetail(detail UserDetail) error
	UpdateUserDetail(etail UserDetail) error
	GetUserDetailById(id string) (UserDetail, error)
}
