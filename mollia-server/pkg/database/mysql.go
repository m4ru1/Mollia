package database

import (
	"database/sql"
	"fmt"
	"log"
	"mollia/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB(cfg config.Config) *sql.DB {
	dsn := cfg.DB_DSN
	if dsn == "" {
		log.Fatal("数据库连接字符串 (DB_DSN) 未在配置中找到")
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	err = DB.Ping()
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	fmt.Println("数据库连接成功！")
	return DB
}
