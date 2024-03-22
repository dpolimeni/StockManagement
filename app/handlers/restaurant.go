package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

type RestaurantHandler struct {
	DB database.Storage
}

// RestaurantHandler is the handler for the restaurant endpoint
// @Summary Add a new restaurant
// @Description Add a new restaurant to the database
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param restaurant body schemas.Restaurant true "Restaurant object that needs to be added"
// @Success 200 {object} schemas.Restaurant
// @Router /api/v1/restaurant [post]
// @Param Authorization header string true "Authorization" Default(Bearer )
func (handler RestaurantHandler) AddRestaurant(c *fiber.Ctx) error {
	var restaurant schemas.Restaurant
	if err := c.BodyParser(&restaurant); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}
	// Check if the id (string) is provided
	if restaurant.Id == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err := handler.DB.NewRestaurant(restaurant)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	return c.Status(200).JSON(restaurant)
}
