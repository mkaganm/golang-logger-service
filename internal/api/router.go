package api

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/logger")
	//routes.Post("/", _)

}
