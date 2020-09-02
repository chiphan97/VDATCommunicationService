package model

type Groups struct {
	AbstractModel
	UserCreate string   `json:"user_id""`
	Name       string   `json:"nameGroup"`
	Type       string   `json:"type"`
	Private    bool     `json:"private"`
	ListUser   []string `json:"users"`
}

const (
	ONE  = "one-to-one"
	MANY = "many-to-many"
)
