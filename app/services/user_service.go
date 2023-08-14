// services/user_service.go
package services

import (
	"attendance/app/models"
	"attendance/app/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (service *UserService) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	return service.UserRepo.InsertUser(ctx, user)
}

func (service *UserService) GetUserByID(ctx context.Context, userId string) (*models.User, error) {
	return service.UserRepo.FindUserByID(ctx, userId)
}

func (service *UserService) GetAllUsers(ctx context.Context, page, pageSize int) ([]models.User, error) {
	return service.UserRepo.FindUsers(ctx, page, pageSize)
}
