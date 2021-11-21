package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/masiucd/go-jwt/app/config"
)

// Protected protect routes
func Protected() fiber.Handler {

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.JwtSecret),
		ErrorHandler: jwtError,
	})
}

// Test protect routes
func Test() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		x := ctx.Request().Header.Peek("Authorization")
		fmt.Println("Hello", string(x))
		if string(x) != "foobar" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return ctx.Next()
	}

}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
