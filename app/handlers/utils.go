package handlers

import (
	"dpolimeni/stockmanagement/app/schemas"
	"errors"
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

func sellProducts(db_materials []schemas.RawMaterial, sold_products []schemas.Product) ([]schemas.RawMaterial, error) {
	// Iterate over the sold products and remove the raw materials from the stock
	for _, sold_product := range sold_products {
		for _, raw_material := range sold_product.RawMaterials {
			found := false
			for i, db_material := range db_materials {
				if db_material.Id == raw_material.Id {
					db_materials[i].Quantity -= raw_material.Quantity
					found = true
				}
			}
			if !found {
				error_string := "Raw material not found: " + raw_material.Name
				return nil, errors.New(error_string)
			}
		}
	}
	return db_materials, nil
}
