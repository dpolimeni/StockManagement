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

func InitApp(app *fiber.App) *database.Mongo {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	mongo := &database.Mongo{
		Username: username,
		Password: password,
		Host:     host,
	}
	mongo.GetClient()
	routes.AuthRoute(app, mongo)
	return mongo
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
	app := fiber.New()
	mongo := InitApp(app)
	routes.RestaurantRoutes(app, mongo)
	token := GetToken(app)
	// If token is empty then we have an error
	if token == "" {
		t.Errorf("Token is empty")
	}

	// Create a Restaurant object with name and address
	payload := strings.NewReader(`{
		"address": "string",
		"id": "string",
		"name": "string",
		"products": [
		  {
			"description": "string",
			"id": "string",
			"name": "string",
			"price": 0,
			"rawMaterials": [
			  {
				"description": "string",
				"id": "string",
				"name": "string",
				"quantity": 0,
				"unit": "string"
			  }
			]
		  }
		],
		"stock": {
		  "rawMaterials": [
			{
			  "description": "string",
			  "id": "string",
			  "name": "string",
			  "quantity": 0,
			  "unit": "string"
			}
		  ]
		}
	  }
	`)

	req := httptest.NewRequest("POST", "/api/v1/restaurant", payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := app.Test(req, -1)
	if err != nil {
		fmt.Println(err)
	}
	// If response status code is not 200 then we have an error
	if resp.StatusCode != 200 {
		t.Errorf("Status code is not 200")
	}
}
