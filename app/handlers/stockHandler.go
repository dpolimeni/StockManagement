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
