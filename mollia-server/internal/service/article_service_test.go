package service_test

import (
	"context"
	"database/sql"
	"mollia/internal/model"
	"mollia/internal/repository/mocks"
	"mollia/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestArticleService_GetArticle(t *testing.T) {
	// === 测试用例 1: 成功获取文章 ===
	t.Run("should get article successfully", func(t *testing.T) {
		// 1. 准备
		mockRepo := new(mocks.MockArticleRepository)

		// 准备期望返回的数据
		expectedArticle := &model.ArticleResponse{
			ID:    1,
			Title: "Test Title",
		}

		// 设置模拟行为：当 GetByID 以任何 context 和 ID 1 被调用时,
		// 返回我们准备好的数据和 nil 错误。
		mockRepo.On("GetByID", mock.Anything, uint(1)).Return(expectedArticle, nil)

		// 创建被测试的服务实例，注入模拟仓库
		articleSvc := service.NewArticleService(mockRepo, nil) // userRepo 在此测试中不使用，传入 nil

		// 2. 执行
		article, err := articleSvc.GetArticle(context.Background(), 1)

		// 3. 断言
		assert.NoError(t, err)                                // 断言没有错误发生
		assert.NotNil(t, article)                             // 断言返回的文章不是 nil
		assert.Equal(t, expectedArticle.Title, article.Title) // 断言标题匹配

		// 验证模拟对象的方法是否被按预期调用
		mockRepo.AssertExpectations(t)
	})

	// === 测试用例 2: 文章未找到 ===
	t.Run("should return error when article not found", func(t *testing.T) {
		// 1. 准备
		mockRepo := new(mocks.MockArticleRepository)

		// 设置模拟行为：当 GetByID 以任何 context 和 ID 2 被调用时,
		// 返回 nil 和 sql.ErrNoRows 错误。
		mockRepo.On("GetByID", mock.Anything, uint(2)).Return(nil, sql.ErrNoRows)

		articleSvc := service.NewArticleService(mockRepo, nil)

		// 2. 执行
		article, err := articleSvc.GetArticle(context.Background(), 2)

		// 3. 断言
		assert.Error(t, err)                // 断言有错误发生
		assert.Nil(t, article)              // 断言返回的文章是 nil
		assert.Equal(t, sql.ErrNoRows, err) // 断言错误类型是 sql.ErrNoRows

		mockRepo.AssertExpectations(t)
	})
}
