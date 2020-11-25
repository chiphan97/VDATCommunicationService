package article

type Payload struct {
	ID        int64
	Content   string `json:"content"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	CreatedBy string
	UpdateBy  string
}

func (p Payload) convertToModel() *Article {
	a := &Article{
		ID:        p.ID,
		Content:   p.Content,
		Title:     p.Title,
		Thumbnail: p.Thumbnail,
		CreatedBy: p.CreatedBy,
		UpdateBy:  p.UpdateBy,
	}
	return a
}
