package schemas

type Material struct {
	Id       string  `json:"id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
}

var TypeMap = map[string]float64{
	"purchase": 1,
	"waste":    -1,
}

type StockChange struct {
	Materials []Material `json:"materials" binding:"required"`
	Type      string     `json:"type" binding:"required"` // Must be purchase or waste
}
