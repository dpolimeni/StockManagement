package main

import (
	"dpolimeni/stockmanagement/pkg/routes"
	"dpolimeni/stockmanagement/platform/database"
	"fmt"
	"log"

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
		log.Print(err)
		// panic(err)
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
	routes.StockRoutes(app, &mongo)
	routes.AuthRoute(app, &mongo)
	routes.RestaurantRoutes(app, &mongo)
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
	test_map := map[string]float64{
		"status": 200,
		"other":  0,
	}

	switch test_map["status"] {
	case 200:
		fmt.Println("Server is running")
		// Return the status of the server
		return c.Status(200).JSON(test_map)
	default:
		fmt.Println("Server is not running")
		return c.Status(500).JSON(test_map)
	}
}
