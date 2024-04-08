package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
	"errors"
	"fmt"
	"slices"
)

func addRawMaterials(db_materials []schemas.RawMaterial, new_materials []schemas.RawMaterial) ([]schemas.RawMaterial, error) {
	// Iterate over the new raw materials add them and return an error if any duplicate is found
	database_ids := []string{}
	for _, db_materia := range db_materials {
		database_ids = append(database_ids, db_materia.Id)
	}

	for _, new_material := range new_materials {
		if slices.Contains(database_ids, new_material.Id) {
			duplicate := new_material.Name
			error_string := "Duplicate raw material: " + duplicate
			return nil, errors.New(error_string)
		}
	}
	db_materials = append(db_materials, new_materials...)
	return db_materials, nil
}

func sellProduct(restaurant *schemas.Restaurant, soldProducts []schemas.SoldProducts) ([]string, error) {
	// For each sold product update the stock
	productsSold := map[string]float64{}
	for _, product := range soldProducts {
		productsSold[product.Id] = product.Quantity
	}
	alertMaterials := []string{}

	// Check if all the products are in the restaurant
	productsCheck := map[string]bool{}
	for _, product := range restaurant.Products {
		switch productsSold[product.Id] {
		case 0:
			continue
		default:
			productsCheck[product.Id] = true
			for _, productMaterial := range product.RawMaterials {
				for i, dbRawMaterial := range restaurant.Stock.RawMaterials {
					if productMaterial.Id == dbRawMaterial.Id {
						restaurant.Stock.RawMaterials[i].Quantity -= productMaterial.Quantity * productsSold[product.Id]
						// If the quantity is less than 0 set to 0 and return an alert
						fmt.Println(restaurant.Stock.RawMaterials[i].Quantity)
						if restaurant.Stock.RawMaterials[i].Quantity < 0 {
							restaurant.Stock.RawMaterials[i].Quantity = 0
							alertMaterials = append(alertMaterials, restaurant.Stock.RawMaterials[i].Name)
						}
					}
				}
			}
		}
	}

	// Check if all the products are in the restaurant
	for _, product := range soldProducts {
		if _, ok := productsCheck[product.Id]; !ok {
			error_string := "Product not found: " + product.Name
			return nil, errors.New(error_string)
		}
	}

	return alertMaterials, nil
}
