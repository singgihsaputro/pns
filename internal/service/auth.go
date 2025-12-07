package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/example/pns/internal/model"
	"github.com/example/pns/internal/repository"
)

type AuthService struct {
	users *repository.UserRepository
}

func NewAuthService(users *repository.UserRepository) *AuthService {
	return &AuthService{users: users}
}

func (s *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *AuthService) Authenticate(ctx context.Context, email, password string) (model.User, error) {
	u, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return model.User{}, errors.New("invalid credentials")
	}
	return u, nil
}

