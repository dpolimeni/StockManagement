package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
	"dpolimeni/stockmanagement/platform/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// RestaurantHandler is the handler for the restaurant endpoint
type RestaurantHandler struct {
	DB database.Storage
}

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

// DeleteRestaurant godoc
// @Summary Delete a restaurant
// @Description Delete a restaurant from the database
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/restaurant [delete]
// @Param Authorization header string true "Authorization" Default(Bearer )
func (handler RestaurantHandler) DeleteRestaurant(c *fiber.Ctx) error {
	restaurantId := c.Query("restaurant")
	if restaurantId == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err := handler.DB.DeleteRestaurant(restaurantId)
	if err != nil {
		error_string := fmt.Sprintf("Internal Server Error %s", err.Error())
		return c.Status(500).JSON(fiber.Map{"message": error_string})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Deleted"})
}
