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

func createUserApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.CreateUserRequest
		if err = c.BodyParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation create user request")
			return kernel.ResponseErrorValidate(c, err)
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error get auth handler")
			return
		}

		ok, err := svc.IsEmailExistsService(c.UserContext(), request.Email)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedCheckEmail)
		}

		if ok {
			return kernel.ResponseError(c, constant.ErrEmailAlreadyExists, msgEmailAlreadyExists)
		}

		err = svc.CreateUserService(c.UserContext(), request, ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedCreateUser)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusCreated,
			Data:    nil,
			Message: msgSuccessCreateUser,
		})
	}
}
