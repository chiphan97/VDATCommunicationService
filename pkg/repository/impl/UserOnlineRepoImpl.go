package impl

import (
	"database/sql"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/repository"
)

type UserOnlineRepoImpl struct {
	Db *sql.DB
}

func NewUserOnlineRepoImpl(db *sql.DB) repository.UserOnlineRepo {
	return &UserOnlineRepoImpl{Db: db}
}
func (u *UserOnlineRepoImpl) GetListUSerOnline(filter string) ([]model.UserOnline, error) {
	userOnlines := make([]model.UserOnline, 0)
	statement := `SELECT user_id,username,first,last,log_at FROM ONLINE `
	if len(filter) > 0 {
		statement = statement + `WHERE username LIKE '` + filter + `%'`
	}
	rows, err := u.Db.Query(statement)
	println(err)
	if err != nil {
		return userOnlines, err
	}
	for rows.Next() {
		var user model.UserOnline
		err = rows.Scan(&user.UserID, &user.Username, &user.First, &user.Last, &user.LogAt)
		if err != nil {
			return userOnlines, err
		}
		userOnlines = append(userOnlines, user)
	}
	return userOnlines, nil
}
func (u *UserOnlineRepoImpl) AddUserOnline(online model.UserOnline) error {
	statement := `INSERT INTO ONLINE VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := u.Db.Exec(statement, online.HostName, online.SocketID, online.UserID, online.Username, online.First, online.Last)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserOnlineRepoImpl) DeleteUserOnline(socketid string) error {
	statement := `DELETE FROM ONLINE WHERE socket_id=$1`
	_, err := u.Db.Exec(statement, socketid)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserOnlineRepoImpl) GetUserOnline(userId string) (model.UserOnline, error) {
	var user model.UserOnline
	statement := `SELECT user_id,username,first,last,log_at FROM ONLINE WHERE user_id=$1`
	rows, err := u.Db.Query(statement, userId)
	//println(err)
	if err != nil {
		return user, err
	}
	if rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.First, &user.Last, &user.LogAt)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
