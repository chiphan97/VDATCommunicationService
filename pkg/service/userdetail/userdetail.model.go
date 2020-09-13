package userdetail

import "time"

type UserDetail struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Username  string     `json:"userName"`
	First     string     `json:"first"`
	Last      string     `json:"last"`
	Role      string     `json:"role"`
}

const (
	ADMIN   = "admin"
	DOCTOR  = "doctor"
	PATIENT = "patient"
)

func (u *UserDetail) ConvertToDto() Dto {
	dto := Dto{
		ID:       u.ID,
		Username: u.Username,
		First:    u.First,
		Last:     u.Last,
		Role:     u.Role,
	}
	return dto
}
