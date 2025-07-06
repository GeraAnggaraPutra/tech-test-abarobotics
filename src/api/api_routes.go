package api

import (
	"os"

	"github.com/gofiber/fiber/v2"

	authDomain "abarobotics-test/src/domain/auth/application"
	deviceDomain "abarobotics-test/src/domain/device/application"
	userDomain "abarobotics-test/src/domain/user/application"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"
	"abarobotics-test/toolkit/config"
)

func routes(app *fiber.App, k *kernel.Kernel) {
	// Register middlewares
	middleware.RecoverMiddleware(app)
	middleware.RateLimiterMiddleware(app)
	middleware.CorsMiddleware(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": os.Getenv("APP_NAME") + " is Running",
		})
	})

	validate := config.NewValidator()

	// domain routes
	authDomain.AddRoutes(app, k, validate)
	userDomain.AddRoutes(app, k, validate)
	deviceDomain.AddRoutes(app, k, validate)
}
