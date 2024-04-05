package schemas

type RawMaterial struct {
	Id          string  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty" bson:"description"`
	Quantity    float64 `json:"quantity" binding:"required"`
	Unit        string  `json:"unit,omitempty" bson:"unit"`
}

type Product struct {
	Id           string        `json:"id" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	Description  string        `json:"description,omitempty" bson:"description"`
	Price        float64       `json:"price" binding:"required"`
	RawMaterials []RawMaterial `json:"rawmaterials" binding:"required"`
}

type Stock struct {
	RawMaterials []RawMaterial `json:"rawmaterials" binding:"required"`
}

type Restaurant struct {
	Id       string    `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Address  string    `json:"address" binding:"required"`
	Products []Product `json:"products" binding:"required"`
	Stock    Stock     `json:"stock" binding:"required"`
}
