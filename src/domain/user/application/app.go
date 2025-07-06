package application

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/user/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"
)

func AddRoutes(app *fiber.App, k *kernel.Kernel, validate *validator.Validator) {
	svc := service.NewService(k.GetDBX())
	mdw := middleware.NewEnsureToken(k.GetDBX())

	routes := app.Group("/users", mdw.ValidateToken())

	routes.Get("", readUserListApp(svc, validate))
	routes.Get("/me", readProfileApp(svc))
	routes.Get("/:guid", readUserDetailApp(svc, validate))
	routes.Post("", createUserApp(svc, validate))
	routes.Put("/:guid", updateUserApp(svc, validate))
	routes.Delete("/:guid", deleteUserApp(svc, validate))
}
