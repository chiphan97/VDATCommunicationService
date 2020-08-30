package model

type Groups struct {
	AbstractModel
	UserCreate string   `json:"user_create""`
	NameGroup  string   `json:"name_group"`
	TypeGroup  string   `json:"type_group"`
	Private    bool     `json:"private"`
	ListUser   []string `json:"list_user"`
}

const (
	ONE  = "ONE_ONE"
	MANY = "MANY_MANY"
)
