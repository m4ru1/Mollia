package model

import "time"

// Category 对应于数据库中的 'categories' 表
type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// CategoryRequest 是用于创建/更新分类的请求体结构
type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
