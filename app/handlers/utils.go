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
			return nil, errors.New("Material already in stock")
		}
	}
	db_materials = append(db_materials, new_materials...)
	return db_materials, nil
}
