package mocks

import (
	"context"
	"mollia/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) Create(ctx context.Context, name string) (int64, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockTagRepository) GetAll(ctx context.Context) ([]model.Tag, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Tag), args.Error(1)
}
func (m *MockTagRepository) Update(ctx context.Context, id uint, name string) error {
	args := m.Called(ctx, id, name)
	return args.Error(0)
}
func (m *MockTagRepository) Delete(ctx context.Context, id uint) (int64, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockTagRepository) GetArticleCount(ctx context.Context, id uint) (int, error) {
	args := m.Called(ctx, id)
	return args.Int(0), args.Error(1)
}
