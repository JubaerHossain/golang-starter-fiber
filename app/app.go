package app

import (
	"github.com/JubaerHossain/golang-starter-fiber/app/controllers"
	"github.com/JubaerHossain/golang-starter-fiber/app/repository"
	"github.com/JubaerHossain/golang-starter-fiber/app/routes"
	"github.com/JubaerHossain/golang-starter-fiber/app/services"
	"github.com/JubaerHossain/golang-starter-fiber/config"
	"github.com/JubaerHossain/golang-starter-fiber/database"
	"github.com/JubaerHossain/golang-starter-fiber/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	// Initialize database connection
	database.Connect()

	// Initialize collections and repositories
	userCollection := utils.Collection("users")
	userRepository := repository.NewUserRepository(userCollection)

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize validator
	validate := validator.New()

	redisClient := database.RedisClient()

	// Initialize controllers
	userController := controllers.NewUserController(userService, validate, redisClient)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		Concurrency:  1000,
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "Attendance",
	})

	app.Use(cors.New())
	app.Use(logger.New())

	// Set up routes
	routes.UserRoute(app, userController)

	// Serve static files
	app.Static("/", "./views") // serve static files

	// Start the server
	fmt.Println("Server is running on port: 🔥🔥" + config.Env("APP_URL") + ":" + config.Env("PORT"))
	app.Listen(":" + config.Env("PORT"))
}
