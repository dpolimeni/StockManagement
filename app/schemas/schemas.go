package schemas

type RawMaterial struct {
	Name        string
	Description string
	Quantity    float64
}

type Product struct {
	Name         string
	Description  string
	Price        float64
	RawMaterials []RawMaterial
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
