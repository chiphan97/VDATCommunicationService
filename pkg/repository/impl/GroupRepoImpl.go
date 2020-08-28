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

//func (groupuser *GroupRepoImpl) GetListGroupByUser(subUser string) ([]model.Groups, error) {
//	var groups []model.Groups
//	statement := `SELECT * FROM Groups AS g
//					INNER JOIN Groups_Users AS g_u
//					ON g.id_group = g_u.id_group
//					WHERE g_u.sub_user_join = $1
//					ORDER BY created_at DESC
//					LIMIT 20`
//	rows, err := groupuser.Db.Query(statement, subUser)
//	if err != nil {
//		return groups, err
//	}
//	for rows.Next() {
//		group := model.Groups{}
//		err = rows.Scan(&group.ID, &group.SubUserCreat, &group.NameGroup,
//			&group.TypeGroup, &group.CreatedAt,
//			&group.UpdatedAt, &group.DeletedAt)
//		if err != nil {
//			return groups, err
//		}
//		groups = append(groups, group)
//	}
//	return groups, nil
//}
func (g *GroupRepoImpl) GetGroupByOwnerAndUserAndTypeOne(owner string, user string) ([]model.Groups, error) {
	groups := make([]model.Groups, 0)
	statement := `SELECT g.id_group, owner_id, type, created_at, updated_at, deleted_at 
					FROM groups AS g
					INNER JOIN groups_users AS g_u
						ON g.id_group = g_u.id_group
						WHERE g.type=$1 
						AND ((owner_id = $2 AND g_u.user_id = $3) 
							OR (owner_id = $3 AND g_u.user_id = $2))`
	rows, err := g.Db.Query(statement, model.ONE, owner, user)
	if err != nil {
		return groups, err
	}
	if rows.Next() {
		var group model.Groups
		err = rows.Scan(&group.ID, &group.SubUserCreat, &group.TypeGroup, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (g *GroupRepoImpl) GetGroupByUser(user string) ([]model.Groups, error) {
	groups := make([]model.Groups, 0)
	statement := `SELECT g.id_group, owner_id, name, type, created_at, updated_at, deleted_at 
 					FROM groups AS g
					INNER JOIN groups_users AS g_u
					ON g.id_group = g_u.id_group
					WHERE  g_u.user_id = $1
						AND g.type = $2
					ORDER BY created_at DESC 
					LIMIT 20`
	rows, err := g.Db.Query(statement, user, model.MANY)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var group model.Groups
		err = rows.Scan(&group.ID, &group.SubUserCreat, &group.NameGroup, &group.TypeGroup, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (g *GroupRepoImpl) AddGroupTypeONE(owner string) (model.Groups, error) {
	var group model.Groups
	statement := `INSERT INTO groups (OWNER_ID, TYPE)  VALUES ($1,$2)`
	_, err := g.Db.Exec(statement, owner, model.ONE)
	if err != nil {
		return group, err
	}
	statement = `SELECT g.id_group, owner_id, type, created_at, updated_at, deleted_at 
 					FROM Groups AS g WHERE owner_id = $1
 					ORDER BY created_at DESC
 					LIMIT 1`
	rows, err := g.Db.Query(statement, owner)
	if err != nil {
		return group, err
	}
	if rows.Next() {
		err = rows.Scan(&group.ID, &group.SubUserCreat, &group.TypeGroup, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return group, err
		}
	}
	return group, nil
}
