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
	// Create Mongo DB connection
	stockhandler := handlers.StockHandler{
		DB: DB,
	}

	route.Use(middleware.JWTMiddleware())
	route.Post("/sell", stockhandler.SellProducts)
	route.Get("/", stockhandler.GetStock)
	route.Post("/update", stockhandler.UpdateMaterials)
	route.Post("/rectify", stockhandler.RectifyStock)
}
