package services

import (
	"chicup/internal/models"
	"chicup/internal/repository"
)

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verify password (omitted for brevity)
	return user, nil
}
