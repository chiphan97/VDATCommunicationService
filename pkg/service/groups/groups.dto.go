package groups

type Dto struct {
	Id          uint   `json:"id"`
	Name        string `json:"nameGroup"`
	Type        string `json:"type"`
	Private     bool   `json:"private"`
	Owner       string `json:"owner"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}
