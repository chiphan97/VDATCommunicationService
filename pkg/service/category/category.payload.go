package category

type PayLoad struct {
	Name     string `json:"name"`
	UserId   string `json:"userId"`
	ParentID int64  `json:"parentId"`
}

func (p PayLoad) convertToModel() Category {
	category := Category{
		Name:      p.Name,
		ParentID:  p.ParentID,
		CreatedBy: p.UserId,
		UpdateBy:  p.UserId,
	}
	return category
}
