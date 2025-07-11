package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mollia/api"
	"mollia/internal/handler"
	"mollia/internal/repository"
	"mollia/internal/service"
	"mollia/pkg/config"
	"mollia/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg, _ := config.LoadConfig()
	db := database.InitDB(cfg)

	// Repositories
	userRepo := repository.NewMySQLUserRepository(db)
	articleRepo := repository.NewMySQLArticleRepository(db)
	categoryRepo := repository.NewMySQLCategoryRepository(db)
	tagRepo := repository.NewMySQLTagRepository(db)

	// Services
	userSvc := service.NewUserService(userRepo, cfg.JWT_SECRET)
	articleSvc := service.NewArticleService(articleRepo, userRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	tagSvc := service.NewTagService(tagRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userSvc)
	articleHandler := handler.NewArticleHandler(articleSvc)
	categoryHandler := handler.NewCategoryHandler(categorySvc)
	tagHandler := handler.NewTagHandler(tagSvc)
	storageHandler := handler.NewStorageHandler()

	router := gin.Default()

	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"} // 允许的前端地址
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// 配置静态文件服务
	router.Static("/static", "./uploads")

	// 注册路由
	api.RegisterRoutes(router, userHandler, articleHandler, categoryHandler, tagHandler, storageHandler, cfg.JWT_SECRET)

	// 4. 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":" + cfg.PORT,
		Handler: router,
	}

	// 5. 在 goroutine 中启动服务器
	go func() {
		log.Printf("服务器正在监听端口: %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 6. 监听中断信号，实现优雅停机
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器 ...")

	// 创建一个 5 秒的超时 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭:", err)
	}

	log.Println("服务器已退出")
}
