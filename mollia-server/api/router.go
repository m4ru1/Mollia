package api

import (
	"mollia/internal/handler"
	"mollia/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	userHandler *handler.UserHandler,
	articleHandler *handler.ArticleHandler,
	categoryHandler *handler.CategoryHandler,
	tagHandler *handler.TagHandler,
	storageHandler *handler.StorageHandler,
	jwtSecret string,
) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
		}

		publicArticles := v1.Group("/articles")
		{
			publicArticles.GET("", articleHandler.GetArticles)
			publicArticles.GET("/:id", articleHandler.GetArticleByID)
		}

		admin := v1.Group("/admin")
		admin.Use(middleware.AuthMiddleware(jwtSecret))
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
				adminCategories.POST("", categoryHandler.CreateCategory)
				adminCategories.GET("", categoryHandler.GetCategories)
				adminCategories.PUT("/:id", categoryHandler.UpdateCategory)
				adminCategories.DELETE("/:id", categoryHandler.DeleteCategory)
			}
			adminTags := admin.Group("/tags")
			{
				adminTags.POST("", tagHandler.CreateTag)
				adminTags.GET("", tagHandler.GetTags)
				adminTags.PUT("/:id", tagHandler.UpdateTag)
				adminTags.DELETE("/:id", tagHandler.DeleteTag)
			}
			adminStorage := admin.Group("/storage")
			{
				adminStorage.POST("/upload", storageHandler.UploadFile)
			}
		}
	}
}
