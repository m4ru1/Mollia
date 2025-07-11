package service

import (
	"context"
	"errors"
	"mollia/internal/model"
	"mollia/internal/repository"
)

type TagService interface {
	Create(ctx context.Context, name string) (int64, error)
	GetAll(ctx context.Context) ([]model.Tag, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) Create(ctx context.Context, name string) (int64, error) {
	return s.repo.Create(ctx, name)
}

func (s *tagService) GetAll(ctx context.Context) ([]model.Tag, error) {
	return s.repo.GetAll(ctx)
}

func (s *tagService) Update(ctx context.Context, id uint, name string) error {
	return s.repo.Update(ctx, id, name)
}

func (s *tagService) Delete(ctx context.Context, id uint) error {
	count, err := s.repo.GetArticleCount(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("无法删除：仍有文章使用该标签")
	}
	rowsAffected, err := s.repo.Delete(ctx, id)
	if err == nil && rowsAffected == 0 {
		return errors.New("标签未找到或已被删除")
	}
	return err
}
