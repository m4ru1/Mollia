package repository

import (
	"context"
	"database/sql"
	"mollia/internal/model"
)

type TagRepository interface {
	Create(ctx context.Context, name string) (int64, error)
	GetAll(ctx context.Context) ([]model.Tag, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) (int64, error)
	GetArticleCount(ctx context.Context, id uint) (int, error)
}

type mysqlTagRepository struct {
	db *sql.DB
}

func NewMySQLTagRepository(db *sql.DB) TagRepository {
	return &mysqlTagRepository{db: db}
}

func (r *mysqlTagRepository) Create(ctx context.Context, name string) (int64, error) {
	query := "INSERT INTO tags (name) VALUES (?)"
	res, err := r.db.ExecContext(ctx, query, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *mysqlTagRepository) GetAll(ctx context.Context) ([]model.Tag, error) {
	query := "SELECT id, name, created_at FROM tags ORDER BY id DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]model.Tag, 0)
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, rows.Err()
}

func (r *mysqlTagRepository) Update(ctx context.Context, id uint, name string) error {
	query := "UPDATE tags SET name = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, name, id)
	return err
}

func (r *mysqlTagRepository) Delete(ctx context.Context, id uint) (int64, error) {
	query := "DELETE FROM tags WHERE id = ?"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *mysqlTagRepository) GetArticleCount(ctx context.Context, id uint) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM article_tags WHERE tag_id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&count)
	return count, err
}
