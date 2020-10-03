package groups

import (
	"encoding/json"
	"fmt"
)

type Dto struct {
	Id          uint   `json:"id"`
	Name        string `json:"nameGroup"`
	Type        string `json:"type"`
	Private     bool   `json:"private"`
	Owner       string `json:"owner"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}
type Dtos []Dto

func (d Dto) MarshalToJsonString() string {
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func (d Dto) ConvertToData() Data {
	data := Data{
		Name:        d.Name,
		Owner:       d.Owner,
		Type:        d.Type,
		Private:     d.Private,
		Description: d.Description,
	}
	return data
}
