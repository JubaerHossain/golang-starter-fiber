package api

import (
	"attendance/models"
	"attendance/repository"
	"attendance/service"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App, repo repository.Repository) {
	userService := service.NewUserService(repo)

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")
		user, err := userService.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.JSON(user)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var newUser models.User
		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		err := userService.CreateUser(&newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
		}
		return c.Status(fiber.StatusCreated).JSON(newUser)
	})
}
