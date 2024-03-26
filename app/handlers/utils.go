package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
)

func addRawMaterials(stock_products *[]schemas.Product, products []schemas.Product) {
	// Iterate over the stock products and update the raw materials
	for i, stock_product := range *stock_products {
		// Find the product in the list of products
		for _, product := range products {
			// Check if the product is the same
			if stock_product.Name == product.Name {
				// Update the raw materials
				(*stock_products)[i].RawMaterials = product.RawMaterials
			}
		}
	}
}
