package repository

import (
	"context"
	"mollia/internal/model"
)

// UserRepository 定义了用户数据操作的接口
type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetIDByUsername(ctx context.Context, username string) (uint, error)
}
