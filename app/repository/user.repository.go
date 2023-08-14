// repository/user_repository.go
package repository

import (
	"attendance/app/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		Collection: collection,
	}
}

func (repo *UserRepository) InsertUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	return repo.Collection.InsertOne(ctx, user)
}

func (repo *UserRepository) FindUserByID(ctx context.Context, userId string) (*models.User, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	var user models.User
	err := repo.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindUsers(ctx context.Context, page, pageSize int) ([]models.User, error) {
	var users []models.User

	// Calculate the number of documents to skip based on the page and pageSize
	skip := (page - 1) * pageSize

	// Define options for pagination
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))

	cursor, err := repo.Collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
