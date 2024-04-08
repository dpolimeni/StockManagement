package database

import (
	"context"
	"dpolimeni/stockmanagement/app/schemas"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(uri)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	m.Client = client

	return err
}

// AuthorizeUser checks if the user is authorized to access the database
func (m *Mongo) AuthorizeUser(username, password string) (bool, error) {
	// Just if want to use mongo
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// collection := m.Client.Database("stockmanagement").Collection("users")
	// err := collection.FindOne(ctx, map[string]string{"username": username, "password": password}).Err()
	// if err != nil {
	// 	return false, err
	// }

	// Get username and password from enviuronment variables
	envusername := os.Getenv("AUTH_USERNAME")
	envpassword := "$2a$10$coMj77X0nLj8vspt942Wfe3vWwriR72ICHTWAgAvFFARW/IOqza8C" // os.Getenv("HASHED_PASSWORD")
	fmt.Println(envusername, envpassword)

	if envusername != username {
		return false, nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(envpassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// AddRestaurant adds a new restaurant to the database
func (m *Mongo) NewRestaurant(restaurant schemas.RestaurantCreate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("stockmanagement").Collection("restaurants")
	_, err := collection.InsertOne(ctx, restaurant)
	if err != nil {
		return err
	}
	return nil
}

// Delete Restaurant deletes a restaurant from the database
func (m *Mongo) DeleteRestaurant(restaurantId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // This set a time limit to execute this function
	defer cancel()
	collection := m.Client.Database("stockmanagement").Collection("restaurants")
	deleted, err := collection.DeleteOne(ctx, bson.M{"id": restaurantId})
	if err != nil {
		return err
	}
	if deleted.DeletedCount == 0 {
		return fmt.Errorf("restaurant not found")
	}
	return nil
}

func (m *Mongo) GetRestaurant(restaurantId string) (schemas.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("stockmanagement").Collection("restaurants")
	var restaurant schemas.Restaurant
	err := collection.FindOne(ctx, bson.M{"id": restaurantId}).Decode(&restaurant)
	if err != nil {
		return restaurant, err
	}
	return restaurant, nil
}

func (m *Mongo) ReplaceRestaurant(restaurant schemas.Restaurant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("stockmanagement").Collection("restaurants")
	_, err := collection.ReplaceOne(ctx, bson.M{"id": restaurant.Id}, restaurant)
	if err != nil {
		return err
	}
	return nil
}
