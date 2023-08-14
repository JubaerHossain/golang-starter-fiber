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

	// Create channels to receive results
	usersCh := make(chan []models.User)
	cacheErrCh := make(chan error)
	serviceErrCh := make(chan error)

	// Concurrently fetch data from cache
	go func() {
		cacheKey := "users:" + page + ":" + pageSize
		cachedData, err := ctrl.RedisClient.Get(ctx, cacheKey).Result()
		if err != nil {
			cacheErrCh <- err
			return
		}
		var users []models.User
		if err := json.Unmarshal([]byte(cachedData), &users); err != nil {
			cacheErrCh <- err
			return
		}
		usersCh <- users
	}()

	// Concurrently fetch data from the service
	go func() {
		users, err := ctrl.UserService.GetAllUsers(ctx, pageInt, pageSizeInt)
		if err != nil {
			serviceErrCh <- err
			return
		}
		usersCh <- users
	}()

	// Wait for results from both channels
	var users []models.User
	var cacheErr, serviceErr error
	for i := 0; i < 2; i++ {
		select {
		case users = <-usersCh:
		case cacheErr = <-cacheErrCh:
		case serviceErr = <-serviceErrCh:
		}
	}

	if cacheErr != nil && serviceErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": cacheErr.Error() + ", " + serviceErr.Error()}})
	} else if cacheErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": cacheErr.Error()}})
	} else if serviceErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": serviceErr.Error()}})
	}

	// Store data in cache (if fetched from the service)
	if serviceErr == nil {
		usersJSON, _ := json.Marshal(users)
		cacheKey := "users:" + page + ":" + pageSize
		ctrl.RedisClient.Set(ctx, cacheKey, string(usersJSON), time.Minute)
	}

	return c.Status(http.StatusOK).JSON(utils.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
}
