package service_test

import (
	"context"
	"mollia/internal/repository/mocks"
	"mollia/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCategoryService_Delete(t *testing.T) {
	t.Run("should delete category successfully when no articles are associated", func(t *testing.T) {
		mockRepo := new(mocks.MockCategoryRepository)
		mockRepo.On("GetArticleCount", mock.Anything, uint(1)).Return(0, nil)
		mockRepo.On("Delete", mock.Anything, uint(1)).Return(int64(1), nil)
		categorySvc := service.NewCategoryService(mockRepo)
		err := categorySvc.Delete(context.Background(), 1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not delete category when articles are associated", func(t *testing.T) {
		mockRepo := new(mocks.MockCategoryRepository)
		mockRepo.On("GetArticleCount", mock.Anything, uint(2)).Return(5, nil)
		categorySvc := service.NewCategoryService(mockRepo)
		err := categorySvc.Delete(context.Background(), 2)
		assert.Error(t, err)
		assert.Equal(t, "无法删除：仍有文章属于该分类", err.Error())
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "Delete", mock.Anything, uint(2))
	})
}
