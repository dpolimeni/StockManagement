package handlers

import (
	"dpolimeni/stockmanagement/platform/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler is the handler for the auth endpoint
type AuthHandler struct {
	Storage database.Storage
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	userLogin := new(UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
}
