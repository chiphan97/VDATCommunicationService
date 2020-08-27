package repository

type GroupUserRepo interface {
	GetListSubUserByGroup(idGourp int) ([]string, error)
}
