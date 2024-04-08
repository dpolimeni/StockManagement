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
		message := fmt.Sprintf("restaurant %s not found", restaurantId)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": message,
		})
	}

	stockLevel := db_restaurant.Stock
	return c.Status(fiber.StatusOK).JSON(stockLevel)
}

// Update the stock with product purchases/wastes
// This should be called when raw materials are purchased/wasted
// @Summary Purchase/Throw away raw materials
// @Description Purchase raw materials for the stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Param StockChange body schemas.StockChange true "Materials to purchase or waste"
// @Param Authorization header string true "Authorization" Default(Bearer )
// @Router /api/v1/stock/update [post]
func (handler StockHandler) UpdateMaterials(c *fiber.Ctx) error {
	restaurant := c.Query("restaurant")
	var stock_change schemas.StockChange
	if err := c.BodyParser(&stock_change); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the request",
		})
	}

	// Check if the type is purchase or waste
	change_type := stock_change.Type
	if change_type != "purchase" && change_type != "waste" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid type",
		})
	}

	// Get the restaurant from the database
	db_restaurant, err := handler.DB.GetRestaurant(restaurant)
	if err != nil {
		message := fmt.Sprintf("restaurant %s not found", restaurant)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": message,
		})
	}
	// Get the stock from the restaurant
	stock := db_restaurant.Stock
	// Iterate over the stock and update the quantities
	for _, material := range stock_change.Materials {
		for i, db_material := range stock.RawMaterials {
			if db_material.Id == material.Id {
				stock.RawMaterials[i].Quantity += material.Quantity * schemas.TypeMap[change_type]
			}
		}
	}
	// Update the stock in the database
	db_restaurant.Stock = stock
	handler.DB.ReplaceRestaurant(db_restaurant)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Stock updated",
		"stock":   stock,
	})
}

// Update the stock with product sells
// This should be called when the restaurant sell some products
// @Summary Sell products
// @Description Sell products from the stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Param Products body []schemas.SoldProducts true "Products to sell"
// @Router /api/v1/stock/sell [post]
func (handler StockHandler) SellProducts(c *fiber.Ctx) error {
	// Get the list of products
	var sold_products []schemas.SoldProducts
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
	alertMaterials, err := sellProduct(&restaurant, sold_products)

	if err != nil {
		err_string := err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err_string,
		})
	}
	// Update the restaurant in the database
	handler.DB.ReplaceRestaurant(restaurant)

	message := "Products sold and stock updated"

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
		"alert":   alertMaterials,
	})
}

// Update the stock with real levels
// This should be called when the stock is updated after a physical count
// @Summary Update stock levels
// @Description Update the stock levels of a restaurant
// @Tags Stock
// @Accept json
// @Produce json
// @Param restaurant query string true "Restaurant ID"
// @Param RectifyStock body schemas.RectifyStock true "Updated stock levels"
// @Param Authorization header string true "Authorization" Default(Bearer )
// @Router /api/v1/stock/rectify [post]
func (handler StockHandler) RectifyStock(c *fiber.Ctx) error {
	restaurantId := c.Query("restaurant")

	restaurant, err := handler.DB.GetRestaurant(restaurantId)
	if err != nil {
		error_string := fmt.Sprintf("Restaurant %s not found", restaurantId)
		return c.Status(400).JSON(fiber.Map{"message": error_string})
	}

	var stock schemas.RectifyStock
	if err := c.BodyParser(&stock); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err = rectifyStock(&restaurant, stock)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	handler.DB.ReplaceRestaurant(restaurant)
	return c.Status(200).JSON(fiber.Map{"message": "Stock updated", "stock": restaurant.Stock})

}
