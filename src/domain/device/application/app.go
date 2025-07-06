package application

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/device/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"
)

func AddRoutes(app *fiber.App, k *kernel.Kernel, validate *validator.Validator) {
	svc := service.NewService(k.GetDBX())
	mdw := middleware.NewEnsureToken(k.GetDBX())

	routes := app.Group("/devices", mdw.ValidateToken())

	routes.Get("", mdw.ValidatePermissionAction(constant.Permission[1], constant.Action[0]), readDeviceListApp(svc, validate))
	routes.Get("/:guid", mdw.ValidatePermissionAction(constant.Permission[1], constant.Action[0]), readDeviceDetailApp(svc, validate))
	routes.Post("", mdw.ValidatePermissionAction(constant.Permission[1], constant.Action[1]), createDeviceApp(svc, validate))
	routes.Put("/:guid", mdw.ValidatePermissionAction(constant.Permission[1], constant.Action[2]), updateDeviceApp(svc, validate))
	routes.Delete("/:guid", mdw.ValidatePermissionAction(constant.Permission[1], constant.Action[3]), deleteDeviceApp(svc, validate))
}
