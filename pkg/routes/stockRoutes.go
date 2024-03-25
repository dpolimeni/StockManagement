package routes

import (
	"dpolimeni/stockmanagement/app/handlers"
	"dpolimeni/stockmanagement/pkg/middleware"
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

// StockRoutes sets up the routes for the stock endpoint
func StockRoutes(app *fiber.App, DB *database.Mongo) {
	route := app.Group("api/v1/stock")
	route.Use(middleware.JWTMiddleware())

	// Create Mongo DB connection
	restaurantHandler := handlers.StockHandler{
		DB: DB,
	}

	route.Get("/", restaurantHandler.GetStock)
}
