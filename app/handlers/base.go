package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/masiucd/go-jwt/app/config"
)

// Home route
func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("Home")
}

// ProtectedRoute func
func ProtectedRoute(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello user with id = " + id,
	})
}

func login(ctx *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "cannot parse json",
			},
		)
	}

	if body.Email != config.UserEmail || body.Password != config.Pass {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "email or password does not match",
		})
	}

	// HS256
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1" // id
	claims["expire"] = time.Now().Add(time.Hour * 24).Unix()

	s, err := token.SignedString([]byte(config.JwtSecret))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"token":   s,
		"user": struct {
			ID       int    `json:"id"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{ID: 1, Email: config.UserEmail, Password: config.Pass},
	})

}
