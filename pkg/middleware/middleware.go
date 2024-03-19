package middleware

// Set up JWT AUTH
import (
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 24).Unix(),
		"iat":      time.Now().Unix(),
	})
	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 24 * 60).Unix(),
		"iat":      time.Now().Unix(),
	})
	secret := os.Getenv("SECRET_KEY")
	access, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refresh, err := refreshtoken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access":  access,
		"refresh": refresh,
	}, nil
}

// Create JWT middleware to attach to routes
func JWTMiddleware() func(*fiber.Ctx) error {
	config := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})
	return config
}

// func CheckSuperuser(DbClient *ent.Client, username string) bool {
// 	user, err := DbClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
// 	if err != nil {
// 		return false
// 	}
// 	if user.IsAdmin {
// 		return true
// 	} else {
// 		return false
// 	}
// }
