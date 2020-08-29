package model

type Groups struct {
	AbstractModel
	SubUserCreat string `json:"sub_user_create""`
	NameGroup    string `json:"name_group"`
	TypeGroup    string `json:"type_group"`
}

const (
	ONE  = "ONE_ONE"
	MANY = "MANY_MANY"
)
