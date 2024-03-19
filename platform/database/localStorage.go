package database

// Storage is an interface that defines the methods that a database storage should implement
type LocalStorage struct {
	Username string
	Password string
	Host     string
}

// GetClient returns a connection to a local storage
func (l *LocalStorage) GetClient() error {
	return nil
}
