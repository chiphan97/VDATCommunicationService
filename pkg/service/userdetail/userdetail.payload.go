package userdetail

type Payload struct {
	ID       string `json:"id"`
	Username string `json:"userName"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Role     string `json:"role"`
}

func (p *Payload) convertToModel() UserDetail {
	u := UserDetail{
		ID:       p.ID,
		Username: p.Username,
		First:    p.First,
		Last:     p.Last,
		Role:     p.Role,
	}
	return u
}
