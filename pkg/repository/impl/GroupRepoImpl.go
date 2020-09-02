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
//		err = rows.Scan(&group.ID, &group.UserCreate, &group.Name,
//			&group.Type, &group.CreatedAt,
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
	statement := `SELECT g.id_group, owner_id,name, type,private, created_at, updated_at, deleted_at 
					FROM groups AS g
					INNER JOIN groups_users AS g_u
						ON g.id_group = g_u.id_group
						WHERE g.type=$1
						AND g.private=$2
						AND ((owner_id = $3 AND g_u.user_id = $4) 
							OR (owner_id = $4 AND g_u.user_id = $3))`
	rows, err := g.Db.Query(statement, model.ONE, true, owner, user)
	if err != nil {
		return groups, err
	}
	if rows.Next() {
		var group model.Groups
		err = rows.Scan(&group.ID, &group.UserCreate, &group.Name, &group.Type, &group.Private, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (g *GroupRepoImpl) GetGroupByUser(user string) ([]model.Groups, error) {
	groups := make([]model.Groups, 0)
	statement := `SELECT g.id_group, owner_id, name, type,private, created_at, updated_at, deleted_at 
 					FROM groups AS g
					INNER JOIN groups_users AS g_u
					ON g.id_group = g_u.id_group
					WHERE  g_u.user_id = $1
					ORDER BY created_at DESC 
					LIMIT 20`
	rows, err := g.Db.Query(statement, user)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var group model.Groups
		err = rows.Scan(&group.ID, &group.UserCreate, &group.Name, &group.Type, &group.Private, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (g *GroupRepoImpl) GetGroupByPrivateAndUser(private bool, user string) ([]model.Groups, error) {
	groups := make([]model.Groups, 0)
	statement := `SELECT g.id_group, owner_id, name, type,private, created_at, updated_at, deleted_at 
 					FROM groups AS g
 					WHERE g.private = $1
 					AND owner_id !=$2
					ORDER BY created_at DESC 
					LIMIT 20`
	rows, err := g.Db.Query(statement, private, user)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var group model.Groups
		err = rows.Scan(&group.ID, &group.UserCreate, &group.Name, &group.Type, &group.Private, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (g *GroupRepoImpl) AddGroupType(group model.Groups) (model.Groups, error) {

	statement := `INSERT INTO groups (owner_id,name ,type,private) VALUES ($1,$2,$3,$4)`
	_, err := g.Db.Exec(statement, group.UserCreate, group.Name, group.Type, group.Private)
	if err != nil {
		return group, err
	}

	statement = `SELECT g.id_group, owner_id,name, type,private, created_at, updated_at, deleted_at 
 					FROM Groups AS g WHERE owner_id = $1
 					ORDER BY created_at DESC
 					LIMIT 1`
	rows, err := g.Db.Query(statement, group.UserCreate)
	if err != nil {
		return group, err
	}
	if rows.Next() {
		err = rows.Scan(&group.ID, &group.UserCreate, &group.Name, &group.Type, &group.Private, &group.CreatedAt, &group.UpdatedAt, &group.DeletedAt)
		if err != nil {
			return group, err
		}
	}
	return group, nil
}
func (g *GroupRepoImpl) UpdateGroup(group model.Groups) (model.Groups, error) {
	var newgroup model.Groups
	statement := `UPDATE groups SET name=$1 WHERE id_group=$2`
	_, err := g.Db.Exec(statement, group.Name, group.ID)
	if err != nil {
		return newgroup, err
	}
	statement = `SELECT id_group, owner_id, name, type,private, created_at, updated_at, deleted_at 
 					FROM groups
 					WHERE  id_group= $1`
	rows, err := g.Db.Query(statement, group.ID)
	if rows.Next() {
		err = rows.Scan(&newgroup.ID, &newgroup.UserCreate, &newgroup.Name, &newgroup.Type, &newgroup.Private, &newgroup.CreatedAt, &newgroup.UpdatedAt, &newgroup.DeletedAt)
		if err != nil {
			return newgroup, err
		}
	}
	return newgroup, nil
}
func (g *GroupRepoImpl) DeleteGroup(idGourp int) error {
	statement := `DELETE FROM Groups_Users WHERE id_group = $1 `
	_, err := g.Db.Exec(statement, idGourp)
	if err != nil {
		return err
	}
	statement = `DELETE FROM Groups WHERE id_group = $1`
	_, err = g.Db.Exec(statement, idGourp)
	if err != nil {
		return err
	}
	return nil
}
func (g *GroupRepoImpl) GetOwnerByGroupAndOwner(owner string, groupId int) (bool, error) {
	statement := `SELECT owner_id FROM Groups WHERE owner_id=$1 AND id_group=$2`
	rows, err := g.Db.Query(statement, owner, groupId)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
