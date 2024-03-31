package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
	"dpolimeni/stockmanagement/platform/database"
	"fmt"

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

// Update the stock with product sells
// This should be called when the restaurant sell some products
// @Summary Sell products
// @Description Sell products from the stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Param Products body []schemas.Product true "Products to sell"
// @Router /api/v1/stock/sell [post]
func (handler StockHandler) SellProducts(c *fiber.Ctx) error {
	// Get the list of products
	var sold_products []schemas.Product
	if err := c.BodyParser(&sold_products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the request",
		})
	}
	// Get the restaurant ID
	restaurantId := c.Query("restaurant")
	// Get the restaurant from the database
	restaurant, err := handler.DB.GetRestaurant(restaurantId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get the restaurant",
		})
	}
	// Get the stock from the restaurant
	stock := restaurant.Stock
	fmt.Println(stock)

	// Iterate over the products and update the stock

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
