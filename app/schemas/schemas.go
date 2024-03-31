package schemas

type RawMaterial struct {
	Id          string
	Name        string
	Description string
	Quantity    float64
	Unit        string
}

type Product struct {
	Id           string        `json:"id" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	Description  string        `json:"description,omitempty" bson:"description"`
	Price        float64       `json:"price" binding:"required"`
	RawMaterials []RawMaterial `json:"raw_materials" binding:"required"`
}

type Stock struct {
	RawMaterials []RawMaterial
}

type Restaurant struct {
	Id       string
	Name     string
	Address  string
	Products []Product
	Stock    Stock
}
