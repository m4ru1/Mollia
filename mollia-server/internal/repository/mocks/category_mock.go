package mocks

import (
	"context"
	"mollia/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository 是 CategoryRepository 的一个模拟实现
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, name string) (int64, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Category), args.Error(1)
}
func (m *MockCategoryRepository) Update(ctx context.Context, id uint, name string) error {
	args := m.Called(ctx, id, name)
	return args.Error(0)
}
func (m *MockCategoryRepository) Delete(ctx context.Context, id uint) (int64, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockCategoryRepository) GetArticleCount(ctx context.Context, id uint) (int, error) {
	args := m.Called(ctx, id)
	return args.Int(0), args.Error(1)
}
