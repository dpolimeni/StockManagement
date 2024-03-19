package handlers

import "github.com/gofiber/fiber/v2"

// StockHandler is the handler for the stock endpoint
// @Summary Get stock
// @Description Get stock level from the database
// @Tags stock
// @Accept */*
// @Produce json
// @Success 200 {object} string
func StockHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
