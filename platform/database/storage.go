package database

import (
	"dpolimeni/stockmanagement/app/schemas"
)

type Storage interface {
	GetClient() error
	AuthorizeUser(username, password string) (bool, error)
	AddRestaurant(restaurant schemas.Restaurant) error
}
