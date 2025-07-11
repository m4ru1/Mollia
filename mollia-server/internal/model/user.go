package model

import "time"

// User 对应于数据库中的 'users' 表
type User struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // 密码哈希不应通过JSON返回
	Avatar       string    `json:"avatar"`
	Bio          string    `json:"bio"`
	CreatedAt    time.Time `json:"createdAt"`
}
