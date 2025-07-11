package api

import (
	"mollia/internal/handler"
	"mollia/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(r *gin.Engine, articleHandler *handler.ArticleHandler) {
	// 创建 v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 认证路由
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handler.Login)
		}

		// 文章路由（公共）
		publicArticles := v1.Group("/articles")
		{
			publicArticles.GET("", articleHandler.GetArticles)
			publicArticles.GET("/:id", articleHandler.GetArticleByID)
		}

		// 管理员路由（需要认证）
		admin := v1.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			adminArticles := admin.Group("/articles")
			{
				adminArticles.POST("", articleHandler.CreateArticle)
				adminArticles.GET("", articleHandler.GetAdminArticles)
				adminArticles.GET("/:id", articleHandler.GetAdminArticleByID)
				adminArticles.PUT("/:id", articleHandler.UpdateArticle)
				adminArticles.DELETE("/:id", articleHandler.DeleteArticle)
			}

			adminCategories := admin.Group("/categories")
			{
				adminCategories.POST("", handler.CreateCategory)
				adminCategories.GET("", handler.GetCategories)
				adminCategories.PUT("/:id", handler.UpdateCategory)
				adminCategories.DELETE("/:id", handler.DeleteCategory)
			}

			adminTags := admin.Group("/tags")
			{
				adminTags.POST("", handler.CreateTag)
				adminTags.GET("", handler.GetTags)
				adminTags.PUT("/:id", handler.UpdateTag)
				adminTags.DELETE("/:id", handler.DeleteTag)
			}

			adminStorage := admin.Group("/storage")
			{
				adminStorage.POST("/upload", handler.UploadFile)
			}
		}
	}
}
