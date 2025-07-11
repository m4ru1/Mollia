package service

import (
	"context"
	"errors"
	"mollia/internal/model"
	"mollia/internal/repository"
)

// ArticleService 定义了文章相关的业务逻辑接口
type ArticleService interface {
	GetArticle(ctx context.Context, id uint) (*model.ArticleResponse, error)
	ListArticles(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error)
	CreateArticle(ctx context.Context, req *model.CreateArticleRequest) (int64, error)
	UpdateArticle(ctx context.Context, id uint, req *model.UpdateArticleRequest) error
	DeleteArticle(ctx context.Context, id uint) error
	ListAdminArticles(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error)
	GetAdminArticleByID(ctx context.Context, id uint) (*model.ArticleResponse, error)
}

type articleService struct {
	articleRepo repository.ArticleRepository
	userRepo    repository.UserRepository
}

// NewArticleService 创建一个新的 ArticleService
func NewArticleService(articleRepo repository.ArticleRepository, userRepo repository.UserRepository) ArticleService {
	return &articleService{articleRepo: articleRepo, userRepo: userRepo}
}

func (s *articleService) GetArticle(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	return s.articleRepo.GetByID(ctx, id)
}

func (s *articleService) ListArticles(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error) {
	return s.articleRepo.List(ctx, limit, offset)
}

func (s *articleService) CreateArticle(ctx context.Context, req *model.CreateArticleRequest) (int64, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return 0, errors.New("无效的用户信息")
	}
	authorID, err := s.userRepo.GetIDByUsername(ctx, username)
	if err != nil {
		return 0, err
	}
	return s.articleRepo.Create(ctx, req, authorID)
}

func (s *articleService) UpdateArticle(ctx context.Context, id uint, req *model.UpdateArticleRequest) error {
	return s.articleRepo.Update(ctx, id, req)
}

func (s *articleService) DeleteArticle(ctx context.Context, id uint) error {
	rowsAffected, err := s.articleRepo.Delete(ctx, id)
	if err == nil && rowsAffected == 0 {
		return errors.New("文章未找到或已被删除")
	}
	return err
}

func (s *articleService) ListAdminArticles(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error) {
	return s.articleRepo.ListAdmin(ctx, limit, offset)
}

func (s *articleService) GetAdminArticleByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	return s.articleRepo.GetAdminByID(ctx, id)
}
