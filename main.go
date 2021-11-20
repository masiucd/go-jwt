package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masiucd/go-jwt/app/routes"
)

func main() {
	app := fiber.New()
	routes.Routes(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
