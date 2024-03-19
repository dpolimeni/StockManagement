package routes

import (
	"dpolimeni/stockmanagement/app/handlers"
	"dpolimeni/stockmanagement/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// StockRoutes sets up the routes for the stock endpoint
func StockRoutes(app *fiber.App) {
	route := app.Group("api/v1/stock")
	route.Use(middleware.JWTMiddleware())

	route.Get("/", handlers.StockHandler)
}
