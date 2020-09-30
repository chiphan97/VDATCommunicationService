package userdetail

import "time"

type UserDetail struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	FullName  string     `json:"fullName"`
	UserName  string     `json:"userName"`
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
		FullName: u.FullName,
		Username: u.UserName,
		First:    u.First,
		Last:     u.Last,
		Role:     u.Role,
	}
	return dto
}
