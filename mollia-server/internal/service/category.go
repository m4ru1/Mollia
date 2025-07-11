package service

import (
	"context"
	"errors"
	"mollia/internal/model"
	"mollia/internal/repository"
)

type CategoryService interface {
	Create(ctx context.Context, name string) (int64, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(ctx context.Context, name string) (int64, error) {
	return s.repo.Create(ctx, name)
}

func (s *categoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *categoryService) Update(ctx context.Context, id uint, name string) error {
	return s.repo.Update(ctx, id, name)
}

func (s *categoryService) Delete(ctx context.Context, id uint) error {
	count, err := s.repo.GetArticleCount(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("无法删除：仍有文章属于该分类")
	}
	rowsAffected, err := s.repo.Delete(ctx, id)
	if err == nil && rowsAffected == 0 {
		return errors.New("分类未找到或已被删除")
	}
	return err
}
