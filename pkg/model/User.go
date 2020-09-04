package model

type User struct {
	UserID   string `json:"userId"`
	Username string `json:"fullName"`
	First    string `json:"firstName"`
	Last     string `json:"lastName"`
}
