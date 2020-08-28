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
