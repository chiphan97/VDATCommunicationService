package model

import (
	"time"
)

type UserOnline struct {
	HostName string     `json:"host_name"`
	SocketID string     `json:"socket_id"`
	UserID   string     `json:"id"`
	Username string     `json:"username"`
	First    string     `json:"first"`
	Last     string     `json:"last"`
	LogAt    *time.Time `json:"log_at"`
}

func (u *UserOnline) ConvertToDto() User {
	user := User{
		UserID:   u.UserID,
		Username: u.Username,
		First:    u.First,
		Last:     u.Last,
	}
	return user
}
