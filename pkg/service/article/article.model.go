package article

import (
	"context"
	"time"
)

type Article struct {
	ID        int64
	Content   string
	Title     string
	Thumbnail string
	Version   int64
	CreatedBy string
	UpdateBy  string
	CreatedAt *time.Time
	UpdateAt  *time.Time
	Slug      string
}

func (a Article) convertToDto() Dto {
	dto := Dto{
		ID:        a.ID,
		Content:   a.Content,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Version:   a.Version,
		CreatedBy: a.CreatedBy,
		UpdateBy:  a.UpdateBy,
		CreatedAt: a.CreatedAt,
		UpdateAt:  a.UpdateAt,
		Slug:      a.Slug,
	}
	return dto
}

type Repo interface {
	Fetch(ctx context.Context) (results []Article, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (results []Article, err error)
	GetByUserId(ctx context.Context, userid string) (results []Article, err error)
	Update(ctx context.Context, a *Article) error
	Store(ctx context.Context, a *Article) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type Service interface {
	Fetch(ctx context.Context) (results []Dto, err error)
	GetByID(ctx context.Context, id int64) (Dto, error)
	GetByTitle(ctx context.Context, title string) (results []Dto, err error)
	GetByUserId(ctx context.Context, userid string) (results []Dto, err error)
	Update(ctx context.Context, p *Payload) (Dto, error)
	Store(ctx context.Context, p *Payload) (Dto, error)
	Delete(ctx context.Context, id int64) error
}
