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

func readDeviceListApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.ReadDeviceListRequest
		if err = c.QueryParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation read devices request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, totalData, err := svc.ReadDeviceListService(c.UserContext(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetDeviceList)
		}

		return kernel.ResponsePaginate(c, request.PaginationPayload, totalData, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToReadDeviceResponses(data),
			Message: msgSuccessGetDeviceList,
		})
	}
}
