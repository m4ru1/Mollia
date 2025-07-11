package repository

import (
	"context"
	"database/sql"
	"mollia/internal/model"
)

type mysqlUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository 创建一个新的 UserRepository 实现
func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	query := "SELECT id, username, password_hash, avatar, bio, created_at FROM users WHERE username = ?"
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Avatar, &user.Bio, &user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) GetIDByUsername(ctx context.Context, username string) (uint, error) {
	var id uint
	query := "SELECT id FROM users WHERE username = ?"
	err := r.db.QueryRowContext(ctx, query, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *mysqlUserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	// ... 实现 GetByID ...
	return nil, nil
}
