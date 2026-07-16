package service

import (
	"context"
	"errors"

	"blipw/internal/models"
	"blipw/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, username string, password string) (models.User, error) {
	if len(username) == 0 || len(password) == 0 {
		return models.User{}, errors.New("username or password is empty")
	}

	_, err := s.userRepo.GetByUsername(ctx, username)
	if err == nil {
		return models.User{}, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.userRepo.Create(ctx, username, string(hashedPassword))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
