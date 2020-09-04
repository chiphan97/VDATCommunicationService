package groups

type GroupsDTO struct {
	Id        uint   `json:"id"`
	Name      string `json:"nameGroup"`
	Type      string `json:"type"`
	Private   bool   `json:"private"`
	Thumbnail string `json:"thumbnail"`
}
