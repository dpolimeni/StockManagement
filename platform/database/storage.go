package database

type Storage interface {
	GetClient() error
}
