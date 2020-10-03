package userdetail

type Repo interface {
	GetListUser() ([]UserDetail, error)
	AddUserDetail(detail UserDetail) error
	UpdateUserDetail(etail UserDetail) error
	GetUserDetailById(id string) (UserDetail, error)
}
