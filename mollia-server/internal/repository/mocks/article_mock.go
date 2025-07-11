package mocks

import (
	"context"
	"mollia/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockArticleRepository 是 ArticleRepository 的一个模拟实现
type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) GetByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ArticleResponse), args.Error(1)
}

func (m *MockArticleRepository) List(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]model.ArticleInList), int64(args.Int(1)), args.Error(2)
}

func (m *MockArticleRepository) Create(ctx context.Context, req *model.CreateArticleRequest, authorID uint) (int64, error) {
	args := m.Called(ctx, req, authorID)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockArticleRepository) Update(ctx context.Context, id uint, req *model.UpdateArticleRequest) error {
	args := m.Called(ctx, id, req)
	return args.Error(0)
}

func (m *MockArticleRepository) Delete(ctx context.Context, id uint) (int64, error) {
	args := m.Called(ctx, id)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockArticleRepository) ListAdmin(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]model.AdminArticleInList), int64(args.Int(1)), args.Error(2)
}

func (m *MockArticleRepository) GetAdminByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ArticleResponse), args.Error(1)
}
