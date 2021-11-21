package handlers

import (
	"fmt"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	config "github.com/masiucd/go-jwt/app/config"
	types "github.com/masiucd/go-jwt/types"
	"github.com/masiucd/go-jwt/util"
)

// Home route
func Home(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello and welcome to Jwt tutorial with Go",
	})
}

// GetMe is the protected route
// We need to be authenticated hen access this route
func GetMe(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := claims["sub"]
	s := fmt.Sprintf("%f", id) // convert float64 to string
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello user with id #= " + strings.ReplaceAll(s, "0", ""),
	})
}

// Login route
// and to receive the Jwt token
func Login(ctx *fiber.Ctx) error {
	var body types.UserRequest
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

	tokenString, err := util.GenerateToken()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "cannot generate token",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"user": types.UserResponse{
			ID:       1,
			Token:    tokenString,
			Email:    config.UserEmail,
			Password: config.Pass,
		},
	})

}
