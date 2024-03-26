package schemas

type RawMaterial struct {
	Name        string
	Description string
	Quantity    float64
}

type Product struct {
	Name         string        `json:"name" binding:"required"`
	Description  string        `json:"description,omitempty" bson:"omitempty"`
	Price        float64       `json:"price" binding:"required"`
	RawMaterials []RawMaterial `json:"raw_materials" binding:"required"`
}

type Stock struct {
	RawMaterials []RawMaterial
	Products     []Product
}

type Restaurant struct {
	Id      string
	Name    string
	Address string
	Stock   Stock
}
