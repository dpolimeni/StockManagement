package routes

import (
	"dpolimeni/stockmanagement/app/handlers"

	"github.com/gofiber/fiber/v2"
)

// StockRoutes sets up the routes for the stock endpoint
func StockRoutes(app *fiber.App) {
	route := app.Group("api/v1/stock")

	route.Get("/", handlers.StockHandler)
}
