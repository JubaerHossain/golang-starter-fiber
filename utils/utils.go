package utils

import (
	"attendance/config"
	"attendance/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(collectionName string) *mongo.Collection {
	client := database.DB
	collection := client.Database(config.Env("DB_NAME")).Collection(collectionName)
	return collection
}
