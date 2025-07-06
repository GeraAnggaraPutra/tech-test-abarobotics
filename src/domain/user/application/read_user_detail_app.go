package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/user/payload"
	"abarobotics-test/src/domain/user/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"
)

func readUserDetailApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.GUIDRequest
		if err = c.ParamsParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation read user request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, err := svc.ReadUserDetailService(c.UserContext(), request.GUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetUserDetail)
		}

		response, err := payload.ToReadDetailUserResponse(data)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error parse response")
			return
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    response,
			Message: msgSuccessGetUserDetail,
		})
	}
}
