package application

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/user/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"
)

func AddRoutes(app *fiber.App, k *kernel.Kernel, validate *validator.Validator) {
	svc := service.NewService(k.GetDBX())
	mdw := middleware.NewEnsureToken(k.GetDBX())

	routes := app.Group("/users", mdw.ValidateToken())

	routes.Get("", mdw.ValidatePermissionAction(constant.Permission[0], constant.Action[0]), readUserListApp(svc, validate))
	routes.Get("/me", readProfileApp(svc))
	routes.Get("/:guid", mdw.ValidatePermissionAction(constant.Permission[0], constant.Action[0]), readUserDetailApp(svc, validate))
	routes.Post("", mdw.ValidatePermissionAction(constant.Permission[0], constant.Action[1]), createUserApp(svc, validate))
	routes.Put("/:guid", mdw.ValidatePermissionAction(constant.Permission[0], constant.Action[2]), updateUserApp(svc, validate))
	routes.Delete("/:guid", mdw.ValidatePermissionAction(constant.Permission[0], constant.Action[3]), deleteUserApp(svc, validate))
}
