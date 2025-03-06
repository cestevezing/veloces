package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Error(c *fiber.Ctx, status int, message string) error {
	response := &Response{
		Status:  status,
		Message: message,
	}
	return c.Status(status).JSON(response)
}

func Success(c *fiber.Ctx, message string, data any) error {
	response := &Response{
		Status:  fiber.StatusOK,
		Message: message,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
