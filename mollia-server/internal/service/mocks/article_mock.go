package mocks

import (
	"context"
	"mollia/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockArticleService 是 ArticleService 的一个模拟实现
type MockArticleService struct {
	mock.Mock
}

func (m *MockArticleService) GetArticle(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ArticleResponse), args.Error(1)
}

// NOTE: 为了演示，其他方法将被省略。在完整的项目中，所有接口方法都应被模拟。
func (m *MockArticleService) ListArticles(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error) {
	panic("implement me")
}
func (m *MockArticleService) CreateArticle(ctx context.Context, req *model.CreateArticleRequest) (int64, error) {
	panic("implement me")
}
func (m *MockArticleService) UpdateArticle(ctx context.Context, id uint, req *model.UpdateArticleRequest) error {
	panic("implement me")
}
func (m *MockArticleService) DeleteArticle(ctx context.Context, id uint) error {
	panic("implement me")
}
func (m *MockArticleService) ListAdminArticles(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error) {
	panic("implement me")
}
func (m *MockArticleService) GetAdminArticleByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	panic("implement me")
}
