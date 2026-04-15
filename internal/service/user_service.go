package service

import (
	"Go/my-go-app/internal/models"
	"errors"
)

// Repository interface allows us to swap databases easily (Factor IV)
type UserRepository interface {
	GetByID(id string) (*models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetProfile(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}
	return s.repo.GetByID(id)
}