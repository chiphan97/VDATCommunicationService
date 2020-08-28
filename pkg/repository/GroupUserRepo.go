package repository

type GroupUserRepo interface {
	GetListUserByGroup(idGourp int) ([]string, error)
	AddGroupUser(users []string, idgroup int) error
}
