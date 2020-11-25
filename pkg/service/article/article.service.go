package article

import (
	"context"
	"time"
)

type ServiceImpl struct {
	repo           Repo
	contextTimeout time.Duration
}

func NewServiceImpl(r Repo, time time.Duration) Service {
	return &ServiceImpl{
		repo:           r,
		contextTimeout: time,
	}
}
func (s *ServiceImpl) Fetch(ctx context.Context) (results []Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	list, err := s.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	for _, a := range list {
		dto := a.convertToDto()
		results = append(results, dto)
	}
	return
}
func (s *ServiceImpl) GetByID(ctx context.Context, id int64) (dto Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return Dto{}, err
	}
	dto = model.convertToDto()
	return
}
func (s *ServiceImpl) GetByTitle(ctx context.Context, title string) (results []Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	list, err := s.repo.GetByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	for _, a := range list {
		dto := a.convertToDto()
		results = append(results, dto)
	}
	return
}
func (s *ServiceImpl) GetByUserId(ctx context.Context, userid string) (results []Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	list, err := s.repo.GetByUserId(ctx, userid)
	if err != nil {
		return nil, err
	}
	for _, a := range list {
		dto := a.convertToDto()
		results = append(results, dto)
	}
	return
}
func (s *ServiceImpl) Update(ctx context.Context, p *Payload) (dto Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	err = s.repo.Update(ctx, p.convertToModel())
	if err != nil {
		return
	}
	model, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return
	}
	dto = model.convertToDto()
	return
}
func (s *ServiceImpl) Store(ctx context.Context, p *Payload) (dto Dto, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	lastId, err := s.repo.Store(ctx, p.convertToModel())
	if err != nil {
		return
	}
	model, err := s.repo.GetByID(ctx, lastId)
	if err != nil {
		return
	}
	dto = model.convertToDto()
	return
}
func (s *ServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	err = s.repo.Delete(ctx, id)
	return
}
