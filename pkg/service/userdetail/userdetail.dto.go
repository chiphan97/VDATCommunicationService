package userdetail

type Dto struct {
	ID       string `json:"id"`
	Username string `json:"userName"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Role     string `json:"role"`
}
