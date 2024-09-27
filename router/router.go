package router

import (
	"go_dev/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1/healthcheck", logger.New())
	api.Get("/", handler.Hello)
}
