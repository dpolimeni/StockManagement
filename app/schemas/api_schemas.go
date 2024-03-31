package schemas

type Material struct {
	Id       string  `json:"id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
}

type Purchase struct {
	Materials []Material `json:"materials" binding:"required"`
}
