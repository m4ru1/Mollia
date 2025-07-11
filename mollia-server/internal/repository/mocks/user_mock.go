package mocks

import (
	"context"
	"mollia/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository 是 UserRepository 的一个模拟实现
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetIDByUsername(ctx context.Context, username string) (uint, error) {
	args := m.Called(ctx, username)
	return uint(args.Int(0)), args.Error(1)
}
