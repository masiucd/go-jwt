package middleware

import "github.com/gofiber/fiber/v2"

// SimpleMiddleware test middleware
func SimpleMiddleware() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		header := ctx.Request().Header.Peek("Authorization")
		if string(header) != "foobar" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return ctx.Next()
	}

}
