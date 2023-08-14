package main

import (
	"attendance/config"
	"attendance/database"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	database.ConnectDB()  // connect to database


    app.Static("/", "./views") // serve static files


    // Routes


    print("Server is running on port: " + config.Env("APP_URL") + ":" + config.Env("PORT"))
	app.Listen(":"+config.Env("PORT"))
}
