package handlers

import "github.com/gofiber/fiber/v2"

// StockHandler is the handler for the stock endpoint
// @Summary Show the status of server.
// @Description Get test on base path.
// @Tags Stock
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/stock [get]
// @Param Authorization header string true "Authorization" Default(Bearer )
func StockHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Get stock levels
// This should be called when the stock levels are needed
func (handler RestaurantHandler) GetStock(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
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
