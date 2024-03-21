package handler

import (
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

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
func RestaurantHandler(c *fiber.Ctx, DB database.Storage) error {

}
