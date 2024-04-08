package schemas

type Material struct {
	Id       string  `json:"id" binding:"required"`       // The id of the material
	Quantity float64 `json:"quantity" binding:"required"` // The quantity of the material
}

var TypeMap = map[string]float64{
	"purchase": 1,
	"waste":    -1,
}

type StockChange struct {
	Materials []Material `json:"materials" binding:"required"`
	Type      string     `json:"type" binding:"required"` // Must be purchase or waste
}

type RectifyStock struct {
	Materials []Material `json:"materials" binding:"required"`
}

type SoldProducts struct {
	Id       string  `json:"id" binding:"required"`       // The id of the product
	Name     string  `json:"name" binding:"required"`     // The name of the product
	Quantity float64 `json:"quantity" binding:"required"` // The quantity of the product sold
}

type ProductCreate struct {
	Id   string `json:"id" binding:"required"`   // The id of the product
	Name string `json:"name" binding:"required"` // The name of the product
}

type RestaurantCreate struct {
	Id       string          `json:"id" binding:"required"`   // The id of the restaurant
	Name     string          `json:"name" binding:"required"` // The name of the restaurant
	Products []ProductCreate `json:"products" binding:"required"`
}
