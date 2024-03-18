package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is a struct that holds the connection information for a MongoDB database
type Mongo struct {
	Username string
	Password string
	Host     string
	Client   *mongo.Client
}

// DataBase returns a connection to a MongoDB database
func (m *Mongo) GetClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	uri := "mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=ClusterTestGo"
	uri = fmt.Sprintf(uri, m.Username, m.Password, m.Host)
	uri = url.QueryEscape(uri)
	fmt.Println(uri)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	m.Client = client

	return err
}
