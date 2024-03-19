package main

import (
	"dpolimeni/stockmanagement/pkg/routes"
	"dpolimeni/stockmanagement/platform/database"
	"fmt"

	"os"

	_ "dpolimeni/stockmanagement/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/joho/godotenv"
)

// @title Swagger Of my apis
// @version 1.0
// @description This is a sample server stock management server
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	fmt.Println(username, password, host)

	mongo := database.Mongo{
		Username: username,
		Password: password,
		Host:     host,
	}
	mongo.GetClient()

	app.Use(logger.New())
	// Let the app use the swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", HealthCheck)
	routes.StockRoutes(app)

	app.Listen(":8080")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description Get test on base path.
// @Tags Root Base
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
