package utils

import (
	"github.com/JubaerHossain/golang-starter-fiber/config"
	"github.com/JubaerHossain/golang-starter-fiber/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(collectionName string) *mongo.Collection {
	client := database.DB
	collection := client.Database(config.Env("DB_NAME")).Collection(collectionName)
	return collection
}
