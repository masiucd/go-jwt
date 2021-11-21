package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/masiucd/go-jwt/app/handlers"
	"github.com/masiucd/go-jwt/app/middleware"
)

// Routes init  routes
func Routes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", middleware.Test(), handlers.Home)
	api.Get("/me", middleware.Protected(), handlers.GetMe)
	api.Post("/login", handlers.Login)
}
