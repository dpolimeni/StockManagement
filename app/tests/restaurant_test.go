package tests

import (
	"dpolimeni/stockmanagement/pkg/routes"
	"dpolimeni/stockmanagement/platform/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func InitApp() (*fiber.App, database.Mongo) {
	app := fiber.New()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	mongo := database.Mongo{
		Username: username,
		Password: password,
		Host:     host,
	}
	routes.AuthRoute(app, &mongo)
	return app, mongo
}

func GetToken(app *fiber.App) string {
	// Login and take the token
	// Create a UserLogin object with username and password
	payload := strings.NewReader(`{
		"username": "stockmanagement",
		"password": "stockmanagement"
	}`)

	req := httptest.NewRequest("POST", "/api/v1/auth/login", payload)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	// Return the response body
	body, err := ioutil.ReadAll(resp.Body)
	// Jsonify the response body
	fmt.Println(string(body))
	jsonBody := make(map[string]interface{})
	json.Unmarshal(body, &jsonBody)
	// Return the token
	if jsonBody["access"] != nil {
		return jsonBody["access"].(string)
	}
	return ""
}

func TestAddRestaurant(t *testing.T) {
	currentDir, _ := os.Getwd()
	projectDir := filepath.Dir(filepath.Dir(currentDir))

	// Check if we have reached the root dir
	fmt.Println(projectDir)
	err := godotenv.Load(filepath.Join(projectDir, ".env"))

	fmt.Println(err)

	app, mongo := InitApp()
	routes.RestaurantRoutes(app, &mongo)
	token := GetToken(app)
	// If token is empty then we have an error
	if token == "" {
		t.Errorf("Token is empty")
	}

	// Create a Restaurant object with name and address
	payload := strings.NewReader(`{
		"name": "Test Restaurant",
		"address": "Test Address"
	}`)
	req := httptest.NewRequest("POST", "/api/v1/restaurant", payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
