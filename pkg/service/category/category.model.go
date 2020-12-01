package category

import "time"

type Category struct {
	ID        int64
	Name      string
	ParentID  int64
	Num       int64
	Version   int64
	CreatedBy string
	UpdateBy  string
	CreatedAt *time.Time
	UpdateAt  *time.Time
}

func (category Category) convertToDto() Dto {
	dto := Dto{
		ID:        category.ID,
		Name:      category.Name,
		ParentID:  category.ParentID,
		Num:       category.Num,
		Version:   category.Version,
		CreatedBy: category.CreatedBy,
		UpdateBy:  category.UpdateBy,
		CreatedAt: category.CreatedAt,
		UpdateAt:  category.UpdateAt,
	}
	return dto
}
