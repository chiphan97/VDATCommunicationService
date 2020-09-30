package groups

import "gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"

type MessageEvent struct {
	Type      string           `json:"type"`
	IdGroup   uint             `json:"groupId"`
	Data      Data             `json:"data"`
	ListGroup []Dto            `json:"listGroup"`
	ListUser  []userdetail.Dto `json:"listUser"`
	UserId    string
	Status    string
}
type Data struct {
	Name        string   `json:"nameGroup"`
	Owner       string   `json:"owner"`
	Type        string   `json:"type"`
	Private     bool     `json:"private"`
	Description string   `json:"description"`
	UserCurrent string   `json:"userCurrent"`
	Users       []string `json:"users"`
}

func (d Data) convertToModel() Groups {
	g := Groups{

		UserCreate:  d.Owner,
		Name:        d.Name,
		Type:        d.Type,
		Private:     d.Private,
		Description: d.Description,
		Users:       d.Users,
	}
	return g
}
func (d Data) ConvertToPayLoad() PayLoad {
	g := PayLoad{
		Name:        d.Name,
		Type:        d.Type,
		Private:     d.Private,
		Description: d.Description,
		Users:       d.Users,
	}
	return g
}
func (d Data) ConvertToData(in Data) Data {
	d.Name = in.Name
	d.Users = in.Users
	d.Description = in.Description
	d.Private = in.Private
	d.Owner = in.Owner

	return d
}
