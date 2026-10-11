package service

import (
	"context"
	"errors"

	"blipw/internal/models"
	"blipw/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserAlreadyExists = errors.New("user already exists")

var ErrInvalidInput = errors.New("invalid input")

var ErrInvalidCredentials = errors.New("invalid username or password")

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

func (s *AuthService) Login(ctx context.Context, username, password string) (models.User, error) {
	if len(username) == 0 || len(password) == 0 {
		return models.User{}, ErrInvalidInput
	}
	if user, err := s.userRepo.GetByUsername(ctx, username); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return models.User{}, ErrInvalidCredentials
		}
		return models.User{}, err
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return models.User{}, ErrInvalidCredentials
		}
		return user, nil
	}
}
