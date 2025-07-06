package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/auth/payload"
	"abarobotics-test/src/domain/auth/service"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"
)

// @Summary      Login
// @Description  Login to system and get access token & refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body payload.LoginRequest true "Login credentials"
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Router       /auth/login [post]
func loginApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.LoginRequest
		if err = c.BodyParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation login request")
			return kernel.ResponseErrorValidate(c, err)
		}

		var (
			userAgent = string(c.Request().Header.UserAgent())
			iPAddress = c.IP()
		)

		data, user, err := svc.LoginService(c.UserContext(), request, userAgent, iPAddress)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedLogin)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(data, user),
			Message: msgSuccessLogin,
		})
	}
}
