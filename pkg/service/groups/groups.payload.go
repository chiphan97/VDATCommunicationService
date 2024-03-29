package groups

type PayLoad struct {
	Name        string   `json:"nameGroup"`
	Type        string   `json:"type"`
	Private     bool     `json:"private"`
	Users       []string `json:"users"`
	Description string   `json:"description"`
}

func (g *PayLoad) ConvertToModel() Groups {
	model := Groups{
		Name:        g.Name,
		Type:        g.Type,
		Private:     g.Private,
		Users:       g.Users,
		Description: g.Description,
	}
	return model
}
