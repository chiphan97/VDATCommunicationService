package impl

import (
	"database/sql"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository"
)

type GroupRepoImpl struct {
	Db *sql.DB
}

func NewGroupRepoImpl(db *sql.DB) repository.GroupRepo {
	return &GroupRepoImpl{Db: db}
}
func (groupuser *GroupRepoImpl) GetListGroupByUser(subUser string) ([]model.Groups, error) {
	var groups []model.Groups
	statement := `SELECT * FROM Groups AS g 
					INNER JOIN Groups_Users AS g_u 
					ON g.id_group = g_u.id_group
					WHERE g_u.sub_user_join = $1
					ORDER BY created_at DESC 
					LIMIT 20`
	rows, err := groupuser.Db.Query(statement, subUser)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		group := model.Groups{}
		err = rows.Scan(&group.ID, &group.SubUserCreat, &group.NameGroup,
			&group.TypeGroup, &group.CreatedAt,
			&group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
