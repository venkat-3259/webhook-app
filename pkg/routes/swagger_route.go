package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// RegisterSwaggerRoute func for describe group of API Docs routes.
func RegisterSwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.New(swagger.Config{
		DocExpansion: "none",
	}))
}
