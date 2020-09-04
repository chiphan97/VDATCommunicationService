package groups

import (
	"time"
)

type AbstractModel struct {
	ID        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
type Groups struct {
	AbstractModel
	UserCreate string   `json:"userId""`
	Name       string   `json:"nameGroup"`
	Type       string   `json:"type"`
	Private    bool     `json:"private"`
	Thumbnail  string   `json:"thumbnail"`
	Users      []string `json:"users"`
}

type GroupsUsers struct {
	AbstractModel
	UserIDJoin string `json:"userId"`
}

const (
	ONE  = "one-to-one"
	MANY = "many-to-many"
)

func (g *Groups) ConvertToDTO() GroupsDTO {
	dto := GroupsDTO{
		Id:        g.ID,
		Name:      g.Name,
		Type:      g.Type,
		Private:   g.Private,
		Thumbnail: g.Thumbnail,
	}
	return dto
}
