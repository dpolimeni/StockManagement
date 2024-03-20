package handlers

import (
	"dpolimeni/stockmanagement/pkg/middleware"
	"dpolimeni/stockmanagement/platform/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	DB database.Storage
}

// Login godoc
// @Summary Login to the application
// @Description Login to the application
// @Tags auth
// @Accept json
// @Produce json
// @Param userLogin body UserLogin true "UserLogin"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/auth/login [post]
func (handler AuthHandler) LoginHandler(c *fiber.Ctx) error {

	userLogin := new(UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Check if the user exists
	authorized, err := handler.DB.AuthorizeUser(userLogin.Username, userLogin.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if authorized {
		fmt.Println("Authorized")
		token, err := middleware.GenerateToken(userLogin.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return c.Status(fiber.StatusOK).JSON(token)
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
}
