package service

import (
	"context"
	"errors"
	"mollia/internal/repository"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type userService struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func NewUserService(repo repository.UserRepository, jwtSecret string) UserService {
	return &userService{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
