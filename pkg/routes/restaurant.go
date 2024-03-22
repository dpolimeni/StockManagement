package routes

import (
	"dpolimeni/stockmanagement/app/handlers"
	"dpolimeni/stockmanagement/pkg/middleware"
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

// RestaurantRoutes sets up the routes for the restaurant endpoint
func RestaurantRoutes(app *fiber.App, DB database.Storage) {
	route := app.Group("api/v1/restaurant")

	// Create Mongo DB connection
	restaurantHandler := handlers.RestaurantHandler{
		DB: DB,
	}

	route.Use(middleware.JWTMiddleware())
	route.Post("/", restaurantHandler.AddRestaurant)
	route.Delete("/", restaurantHandler.DeleteRestaurant)
}
