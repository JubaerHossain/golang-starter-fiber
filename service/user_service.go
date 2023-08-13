package service

import (
	"attendance/models"
	"attendance/repository"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(user *models.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
