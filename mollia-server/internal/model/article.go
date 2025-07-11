package model

import "time"

// ArticleResponse 是用于 API 响应的文章结构
type ArticleResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ArticleInList 是用于文章列表 API 响应的简化结构
type ArticleInList struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateArticleRequest 是用于创建文章的请求体结构
type CreateArticleRequest struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	Tags       []string `json:"tags"`
	CategoryID uint     `json:"categoryId" binding:"required"`
	Status     string   `json:"status" binding:"required,oneof=draft published"`
}

// AdminArticleInList 是用于后台文章列表 API 响应的结构
type AdminArticleInList struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UpdateArticleRequest 是用于更新文章的请求体结构
type UpdateArticleRequest struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	Tags       []string `json:"tags"`
	CategoryID uint     `json:"categoryId" binding:"required"`
	Status     string   `json:"status" binding:"required,oneof=draft published"`
}
