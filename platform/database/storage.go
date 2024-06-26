package database

import (
	"dpolimeni/stockmanagement/app/schemas"
)

type Storage interface {
	GetClient() error
	AuthorizeUser(username, password string) (bool, error)
	NewRestaurant(restaurant schemas.RestaurantCreate) error
	DeleteRestaurant(restaurantId string) error
	GetRestaurant(restaurantId string) (schemas.Restaurant, error)
	ReplaceRestaurant(restaurant schemas.Restaurant) error
}
