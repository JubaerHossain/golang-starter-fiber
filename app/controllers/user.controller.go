// controllers/user_controller.go
package controllers

import (
	"attendance/app/models"
	"attendance/app/services"
	"attendance/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
	Validate    *validator.Validate
	RedisClient *redis.Client
}

func NewUserController(userService *services.UserService, validate *validator.Validate, redisClient *redis.Client) *UserController {
	return &UserController{
		UserService: userService,
		Validate:    validate,
		RedisClient: redisClient,
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
	// Get pagination parameters from the query parameters
	page := c.Query("page", "1")          // Default to page 1 if not provided
	pageSize := c.Query("pageSize", "10") // Default to 10 items per page if not provided

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Invalid page parameter"}})
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Invalid pageSize parameter"}})
	}

	// Attempt to retrieve data from cache
	cacheKey := "users:" + page + ":" + pageSize
	cachedData, err := ctrl.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, return it
		var users []models.User
		if err := json.Unmarshal([]byte(cachedData), &users); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		return c.Status(http.StatusOK).JSON(utils.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
	}

	// Data not found in cache, fetch from the service
	users, err := ctrl.UserService.GetAllUsers(ctx, pageInt, pageSizeInt)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Store data in cache
	usersJSON, _ := json.Marshal(users)
	ctrl.RedisClient.Set(ctx, cacheKey, string(usersJSON), time.Minute)

	return c.Status(http.StatusOK).JSON(utils.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
}
