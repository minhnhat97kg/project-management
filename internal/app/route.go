package app

import (
	"project-management/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoute(app *fiber.App, handlers *handler.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	common := v1.Group("/")

	common.Post("/register", handlers.User.Register)
	common.Post("/login", handlers.User.Login)

}
