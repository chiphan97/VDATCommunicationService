package groups

type GroupsPayLoad struct {
	Name    string   `json:"nameGroup"`
	Type    string   `json:"type"`
	Private bool     `json:"private"`
	Users   []string `json:"users"`
}

func (g *GroupsPayLoad) ConvertToModel() Groups {
	model := Groups{
		Name:    g.Name,
		Type:    g.Type,
		Private: g.Private,
		Users:   g.Users,
	}
	return model
}
