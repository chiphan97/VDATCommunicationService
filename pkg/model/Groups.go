package model

type Groups struct {
	AbstractModel
	UserCreate string   `json:"user_id""`
	NameGroup  string   `json:"nameGroup"`
	TypeGroup  string   `json:"type"`
	Private    bool     `json:"private"`
	ListUser   []string `json:"users"`
}

const (
	ONE  = "one-to-one"
	MANY = "many-to-many"
)
