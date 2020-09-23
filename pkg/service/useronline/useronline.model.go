package useronline

import (
	"time"
)

type UserOnline struct {
	UserID string     `json:"id"`
	LogAt  *time.Time `json:"log_at"`
}

func (u *UserOnline) ConvertToDto() Dto {
	user := Dto{
		UserID: u.UserID,
		LogAt:  u.LogAt,
	}
	return user
}
