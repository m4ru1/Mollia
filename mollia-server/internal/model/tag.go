package model

import "time"

// Tag 对应于数据库中的 'tags' 表
type Tag struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// TagRequest 是用于创建/更新标签的请求体结构
type TagRequest struct {
	Name string `json:"name" binding:"required"`
}
