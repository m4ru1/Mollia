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
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Login(t *testing.T) {
	// 准备：为密码 "password" 生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	mockUser := &model.User{
		ID:           1,
		Username:     "testuser",
		PasswordHash: string(hashedPassword),
	}

	jwtSecret := "test_secret"

	// === 测试用例 1: 登录成功 ===
	t.Run("should login successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		// 当 GetByUsername 以 "testuser" 调用时，返回模拟用户
		mockRepo.On("GetByUsername", mock.Anything, "testuser").Return(mockUser, nil)

		userSvc := service.NewUserService(mockRepo, jwtSecret)
		token, err := userSvc.Login(context.Background(), "testuser", "password")

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)
	})

	// === 测试用例 2: 用户名错误 ===
	t.Run("should fail with wrong username", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		// 当 GetByUsername 以 "wronguser" 调用时，返回未找到错误
		mockRepo.On("GetByUsername", mock.Anything, "wronguser").Return(nil, sql.ErrNoRows)

		userSvc := service.NewUserService(mockRepo, jwtSecret)
		token, err := userSvc.Login(context.Background(), "wronguser", "password")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, "用户名或密码错误", err.Error())
		mockRepo.AssertExpectations(t)
	})

	// === 测试用例 3: 密码错误 ===
	t.Run("should fail with wrong password", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		// 仍然返回正确的用户，但 service 层应能处理密码比较失败的情况
		mockRepo.On("GetByUsername", mock.Anything, "testuser").Return(mockUser, nil)

		userSvc := service.NewUserService(mockRepo, jwtSecret)
		token, err := userSvc.Login(context.Background(), "testuser", "wrongpassword")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, "用户名或密码错误", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
