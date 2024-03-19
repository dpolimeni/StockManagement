package schemas

import (
	"fmt"
)

type Stock struct {
	Name  string
	Price float64
}

func (s *Stock) String() string {
	return fmt.Sprintf("Name: %s, Price: %f", s.Name, s.Price)
}
