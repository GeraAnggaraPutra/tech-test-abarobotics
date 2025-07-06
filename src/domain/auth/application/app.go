package application

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/auth/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"
)

func AddRoutes(app *fiber.App, k *kernel.Kernel, validate *validator.Validator) {
	svc := service.NewService(k.GetDBX())
	mdw := middleware.NewEnsureToken(k.GetDBX())

	routes := app.Group("/auth")

	routes.Post("/login", loginApp(svc, validate))
	routes.Post("/refresh-token", refreshTokenApp(svc, validate))
	routes.Post("/logout", mdw.ValidateToken(), logoutApp(svc))
}
