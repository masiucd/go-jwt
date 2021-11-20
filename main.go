package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello")
	})
	app.Post("/login", login)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
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

	if body.Email != "masiu@ex.com" || body.Password != "123456" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "email or password does not match",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1" // id
	claims["expire"] = time.Now().Add(time.Hour * 24).Unix()

	s, err := token.SignedString([]byte("secret"))

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
		}{ID: 1, Email: "masiu@ex.com", Password: "123456"},
	})

}
