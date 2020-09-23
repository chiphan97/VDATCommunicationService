package useronline

import (
	"database/sql"
)

type RepoImpl struct {
	Db *sql.DB
}

func NewRepoImpl(db *sql.DB) Repo {
	return &RepoImpl{Db: db}
}

func (u *RepoImpl) AddUserOnline(online UserOnline) error {
	statement := `INSERT INTO ONLINE (hostname,socket_id,user_id) VALUES ($1,$2,$3)`
	_, err := u.Db.Exec(statement,
		online.HostName,
		online.SocketID,
		online.UserID)
	if err != nil {
		return err
	}
	return nil
}
func (u *RepoImpl) DeleteUserOnline(socketid string) error {
	statement := `DELETE FROM ONLINE WHERE socket_id=$1`
	_, err := u.Db.Exec(statement, socketid)
	if err != nil {
		return err
	}
	return nil
}
func (u *RepoImpl) GetUserOnlineBySocketIdAndHostId(socketID string, hostname string) (UserOnline, error) {
	var user UserOnline
	statement := `SELECT * FROM ONLINE WHERE hostname=$1 AND socket_id=$2`
	rows, err := u.Db.Query(statement, hostname, socketID)
	//println(err)
	if err != nil {
		return user, err
	}
	if rows.Next() {
		err = rows.Scan(&user.HostName,
			&user.SocketID,
			&user.UserID,
			&user.LogAt)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
