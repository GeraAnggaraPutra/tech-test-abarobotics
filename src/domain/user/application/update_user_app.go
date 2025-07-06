package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/user/payload"
	"abarobotics-test/src/domain/user/service"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"

)

func updateUserApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.UpdateUserRequest

		request.GUID = c.Params("guid")
		if request.GUID == "" {
			return kernel.ResponseError(c, constant.ErrGUID, msgFailedGUID)
		}

		if err = c.BodyParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}
		

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation update user request")
			return kernel.ResponseErrorValidate(c, err)
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error get auth handler")
			return
		}

		ok, err := svc.IsEmailExistsExcludeUserService(c.UserContext(), request.Email, request.GUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedCheckEmail)
		}

		if ok {
			return kernel.ResponseError(c, constant.ErrEmailAlreadyExists, msgEmailAlreadyExists)
		}

		err = svc.UpdateUserService(c.UserContext(), request, ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedUpdateUser)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateUser,
		})
	}
}
