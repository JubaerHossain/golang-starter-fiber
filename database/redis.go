package database

import (
	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Replace with your Redis server address
		DB:   0,                // Replace with your desired Redis DB number
	})

	return client
}
