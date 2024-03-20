package database

type Storage interface {
	GetClient() error
	AuthorizeUser(username, password string) (bool, error)
}
