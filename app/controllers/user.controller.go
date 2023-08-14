// controllers/user_controller.go
package controllers

import (
	"attendance/app/models"
	"attendance/app/services"
	"attendance/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
	Validate    *validator.Validate
}

func NewUserController(userService *services.UserService, validate *validator.Validate) *UserController {
	return &UserController{
		UserService: userService,
		Validate:    validate,
	}
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := ctrl.Validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	result, err := ctrl.UserService.CreateUser(ctx, user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(utils.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func (ctrl *UserController) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	users, err := ctrl.UserService.GetAllUsers(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(utils.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
}

// Implement other controller methods using UserService
