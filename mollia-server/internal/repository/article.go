package repository

import (
	"context"
	"mollia/internal/model"
)

// ArticleRepository 定义了文章数据操作的接口
type ArticleRepository interface {
	GetByID(ctx context.Context, id uint) (*model.ArticleResponse, error)
	List(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error)
	Create(ctx context.Context, req *model.CreateArticleRequest, authorID uint) (int64, error)
	Update(ctx context.Context, id uint, req *model.UpdateArticleRequest) error
	Delete(ctx context.Context, id uint) (int64, error)
	ListAdmin(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error)
	GetAdminByID(ctx context.Context, id uint) (*model.ArticleResponse, error)
}
