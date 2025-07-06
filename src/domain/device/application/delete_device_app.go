package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/device/payload"
	"abarobotics-test/src/domain/device/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"
)

// @Summary      Delete device
// @Description  Delete device by GUID
// @Tags         Device
// @Accept       json
// @Produce      json
// @Param        guid   path      string  true  "Device GUID"
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Router       /devices/{guid} [delete]
func deleteDeviceApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.GUIDRequest
		if err = c.ParamsParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation delete device request")
			return kernel.ResponseErrorValidate(c, err)
		}

		err = svc.DeleteDeviceService(c.UserContext(), request.GUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedDeleteDevice)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteDevice,
		})
	}
}
