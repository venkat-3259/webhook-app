package routes

import (
	"webhook/app/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(a *fiber.App, h *controller.Handler) {

	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/uplink", h.UplinkHandler) // receive data from user

}
