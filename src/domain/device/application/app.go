package application

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/device/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/src/middleware"

)

func AddRoutes(app *fiber.App, k *kernel.Kernel, validate *validator.Validator) {
	svc := service.NewService(k.GetDBX())
	mdw := middleware.NewEnsureToken(k.GetDBX())

	routes := app.Group("/devices", mdw.ValidateToken())

	routes.Get("", readDeviceListApp(svc, validate))
	routes.Get("/:guid", readDeviceDetailApp(svc, validate))
	routes.Post("", createDeviceApp(svc, validate))
	routes.Put("/:guid", updateDeviceApp(svc, validate))
	routes.Delete("/:guid", deleteDeviceApp(svc, validate))
}
