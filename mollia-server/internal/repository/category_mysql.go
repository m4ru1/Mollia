package repository

import (
	"context"
	"database/sql"
	"mollia/internal/model"
)

type CategoryRepository interface {
	Create(ctx context.Context, name string) (int64, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) (int64, error)
	GetArticleCount(ctx context.Context, id uint) (int, error)
}

type mysqlCategoryRepository struct {
	db *sql.DB
}

func NewMySQLCategoryRepository(db *sql.DB) CategoryRepository {
	return &mysqlCategoryRepository{db: db}
}

func (r *mysqlCategoryRepository) Create(ctx context.Context, name string) (int64, error) {
	query := "INSERT INTO categories (name) VALUES (?)"
	res, err := r.db.ExecContext(ctx, query, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *mysqlCategoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	query := "SELECT id, name, created_at FROM categories ORDER BY id DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, rows.Err()
}

func (r *mysqlCategoryRepository) Update(ctx context.Context, id uint, name string) error {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, name, id)
	return err
}

func (r *mysqlCategoryRepository) Delete(ctx context.Context, id uint) (int64, error) {
	query := "DELETE FROM categories WHERE id = ?"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *mysqlCategoryRepository) GetArticleCount(ctx context.Context, id uint) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM articles WHERE category_id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&count)
	return count, err
}
