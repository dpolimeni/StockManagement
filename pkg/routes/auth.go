package routes

import (
	"dpolimeni/stockmanagement/app/handlers"
	"dpolimeni/stockmanagement/pkg/middleware"
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, DB database.Storage) {
	route := app.Group("api/v1/auth")

	authHandler := handlers.AuthHandler{
		DB: DB,
	}

	route.Post("/login", authHandler.LoginHandler)
	route.Use(middleware.JWTMiddleware())
	route.Post("/refresh", authHandler.RefreshToken)
}
