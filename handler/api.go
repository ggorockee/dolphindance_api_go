package handler

import (
	"go_dev/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	jsonResponse := response.Response{
		Error:   false,
		Message: "Hello! World!",
		Data:    nil,
	}

	return c.Status(http.StatusOK).JSON(jsonResponse)
}
