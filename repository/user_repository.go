package repository

import (
	"attendance/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type repository struct {
	mongoClient *mongo.Client
	ctx         context.Context
}

func NewRepository(mongoClient *mongo.Client, ctx context.Context) Repository {
	return &repository{mongoClient: mongoClient, ctx: ctx}
}

func (r *repository) GetUserByID(id string) (*models.User, error) {
	collection := r.mongoClient.Database("your_db_name").Collection("users")
	var user models.User
	err := collection.FindOne(r.ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateUser(user *models.User) error {
	collection := r.mongoClient.Database("your_db_name").Collection("users")
	_, err := collection.InsertOne(r.ctx, user)
	if err != nil {
		return err
	}
	return nil
}
