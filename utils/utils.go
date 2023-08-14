package utils

import (
	"attendance/config"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(config.Env("DB_NAME")).Collection(collectionName)
	return collection
}



