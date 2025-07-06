package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/device/payload"
	"abarobotics-test/src/domain/device/service"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"
)

// @Summary      Update device
// @Description  Update existing device by GUID
// @Tags         Device
// @Accept       json
// @Produce      json
// @Param        guid     path      string  true  "Device GUID"
// @Param        request  body      payload.UpdateDeviceRequest  true  "Device update payload"
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Router       /devices/{guid} [put]
func updateDeviceApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.UpdateDeviceRequest
		request.GUID = c.Params("guid")
		if request.GUID == "" {
			return kernel.ResponseError(c, constant.ErrGUID, msgFailedGUID)
		}

		if err = c.BodyParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation update device request")
			return kernel.ResponseErrorValidate(c, err)
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error get auth handler")
			return
		}

		err = svc.UpdateDeviceService(c.UserContext(), request, ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedUpdateDevice)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateDevice,
		})
	}
}
