package database

import (
	"github.com/JubaerHossain/golang-starter-fiber/config"

	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Env("REDIS_HOST") + ":" + config.Env("REDIS_PORT"),
		Password: config.Env("REDIS_PASSWORD"), // no password set
		DB:       0,                            // Replace with your desired Redis DB number                        // use default 3
	})

	return client
}
