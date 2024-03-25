package handlers

import (
	"dpolimeni/stockmanagement/platform/database"

	"github.com/gofiber/fiber/v2"
)

// RestaurantHandler is the handler for the restaurant endpoint
type StockHandler struct {
	DB *database.Mongo
}

// Get stock levels
// This should be called when the stock levels are needed
// @Summary Get the stock levels
// @Description Get the stock levels of a restaurant
// @Tags Stock
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Param Authorization header string true "Authorization" Default(Bearer )
// @Router /api/v1/stock [get]
func (handler StockHandler) GetStock(c *fiber.Ctx) error {
	// Create an empty restaurant
	restaurantId := c.Query("restaurant")

	// Get the stock levels
	db_restaurant, err := handler.DB.GetRestaurant(restaurantId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get the stock levels",
		})
	}

	stockLevel := db_restaurant.Stock
	return c.Status(fiber.StatusOK).JSON(stockLevel)
}

// Initial handler to populate the stock
func (handler RestaurantHandler) CreateStock(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Update the stock with product sells
// This should be called when products are sell
func (handler RestaurantHandler) SellProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Update the stock with product purchases
// This should be called when raw materials are purchased (at every order)
func (handler RestaurantHandler) PurchaseMaterial(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Update the stock with real levels
// This should be called when the stock is updated after a physical count
func (handler RestaurantHandler) UpdateStock(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
