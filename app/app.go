package app

import (
	"attendance/app/controllers"
	"attendance/app/repository"
	"attendance/app/routes"
	"attendance/app/services"
	"attendance/config"
	"attendance/database"
	"attendance/utils"

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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Concurrency: 1000,
		Prefork:    true,
		ServerHeader: "Fiber",
		AppName: "Attendance",

	})

	app.Use(cors.New())
	app.Use(logger.New())
	

	// Set up routes
	routes.UserRoute(app, userController)

	// Serve static files
	app.Static("/", "./views") // serve static files

	// Start the server
	print("Server is running on port: " + config.Env("APP_URL") + ":" + config.Env("PORT"))
	app.Listen(":" + config.Env("PORT"))
}
