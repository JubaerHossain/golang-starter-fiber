package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BuildPaginatedResponse(c *fiber.Ctx, data interface{}, page, pageSize int, hasNext bool) error {
	paginationInfo := fiber.Map{
		"page":     page,
		"pageSize": pageSize,
		"hasNext":  hasNext,
	}
	return c.Status(http.StatusOK).JSON(Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{
		"data":       data,
		"pagination": paginationInfo,
	}})

}
