package impl

import (
	"database/sql"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository"
)

type GroupUserRepoImpl struct {
	Db *sql.DB
}

func NewGroupUserRepoImpl(db *sql.DB) repository.GroupUserRepo {
	return &GroupUserRepoImpl{Db: db}
}
func (groupuser *GroupUserRepoImpl) GetListSubUserByGroup(idGourp int) ([]string, error) {
	var subUsers []string
	statement := `SELECT sub_user_join FROM Groups_Users 
					WHERE id_group =$1`
	rows, err := groupuser.Db.Query(statement, idGourp)
	if err != nil {
		return subUsers, err
	}
	for rows.Next() {
		var subject string
		err = rows.Scan(&subject)
		if err != nil {
			return subUsers, err
		}
		subUsers = append(subUsers, subject)
	}
	return subUsers, nil
}
