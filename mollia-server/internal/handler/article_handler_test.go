package handler_test

import (
	"encoding/json"
	"errors"
	"mollia/internal/handler"
	"mollia/internal/model"
	"mollia/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestArticleHandler_GetArticleByID(t *testing.T) {
	// 设置 Gin 为测试模式
	gin.SetMode(gin.TestMode)

	t.Run("should return article successfully", func(t *testing.T) {
		// 准备
		mockService := new(mocks.MockArticleService)
		expectedArticle := &model.ArticleResponse{ID: 1, Title: "Test Title"}
		// 模拟 Service 层将成功返回文章
		mockService.On("GetArticle", mock.Anything, uint(1)).Return(expectedArticle, nil)

		articleHandler := handler.NewArticleHandler(mockService)

		// 设置路由
		router := gin.Default()
		router.GET("/articles/:id", articleHandler.GetArticleByID)

		// 创建一个 HTTP 请求和响应记录器
		req, _ := http.NewRequest(http.MethodGet, "/articles/1", nil)
		w := httptest.NewRecorder()

		// 执行请求
		router.ServeHTTP(w, req)

		// 断言
		assert.Equal(t, http.StatusOK, w.Code)

		// 验证响应体
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, float64(200), responseBody["code"])

		data := responseBody["data"].(map[string]interface{})
		assert.Equal(t, float64(expectedArticle.ID), data["id"])
		assert.Equal(t, expectedArticle.Title, data["title"])

		mockService.AssertExpectations(t)
	})

	t.Run("should return 404 when article not found", func(t *testing.T) {
		mockService := new(mocks.MockArticleService)
		// 模拟 Service 层返回错误
		mockService.On("GetArticle", mock.Anything, uint(2)).Return(nil, errors.New("文章未找到"))

		articleHandler := handler.NewArticleHandler(mockService)
		router := gin.Default()
		router.GET("/articles/:id", articleHandler.GetArticleByID)

		req, _ := http.NewRequest(http.MethodGet, "/articles/2", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockService.AssertExpectations(t)
	})
}
