package web

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected(jwtKey string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(jwtKey)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(ErrorResponse("Missing or malformed JWT"))
	}
	//log.Println(err.Error())
	return c.Status(fiber.StatusUnauthorized).
		JSON(ErrorResponse("Invalid or expired JWT"))
}
