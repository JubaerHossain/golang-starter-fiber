package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type PaginatedResponse struct {
	Ctx        *fiber.Ctx
	Data       interface{}
	Page       int
	PageSize   int
	TotalUsers int
}

func NewErrorResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func NewSuccessResponse(status int, message string, data *fiber.Map) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
