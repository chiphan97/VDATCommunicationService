package userdetail

import "database/sql"

type RepoImpl struct {
	Db *sql.DB
}

func NewRepoImpl(db *sql.DB) Repo {
	return &RepoImpl{db}
}
func (u *RepoImpl) GetListUser(filter string) ([]UserDetail, error) {
	details := make([]UserDetail, 0)
	statement := `SELECT * FROM userdetail `
	if len(filter) > 0 {
		statement = statement + `WHERE username LIKE '` + filter + `%'`
	}
	rows, err := u.Db.Query(statement)
	println(err)
	if err != nil {
		return details, err
	}
	for rows.Next() {
		var user UserDetail
		err = rows.Scan(&user.ID,
			&user.Username,
			&user.First,
			&user.Last,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)
		if err != nil {
			return details, err
		}
		details = append(details, user)
	}
	return details, nil
}
func (u *RepoImpl) AddUserDetail(detail UserDetail) error {
	statement := `insert into userdetail(user_id,username,first,last,role) values($1,$2,$3,$4,$5)`
	_, err := u.Db.Exec(statement, detail.ID,
		detail.Username,
		detail.First,
		detail.Last,
		detail.Role)
	if err != nil {
		return err
	}
	return nil
}
func (u *RepoImpl) GetUserDetailById(id string) (UserDetail, error) {
	var detail UserDetail
	statement := `select * from  userdetail where user_id = $1`
	rows, err := u.Db.Query(statement, id)
	if err != nil {
		return detail, err
	}
	if rows.Next() {
		err := rows.Scan(&detail.ID,
			&detail.Username,
			&detail.First,
			&detail.Last,
			&detail.Role,
			&detail.CreatedAt,
			&detail.UpdatedAt,
			&detail.DeletedAt)
		if err != nil {
			return detail, err
		}
	}
	return detail, nil
}
