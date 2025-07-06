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

// @Summary      Refresh Token
// @Description  Get new access token using refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body payload.RefreshTokenRequest true "Refresh token payload"
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Failure      401  {object}  kernel.responseErrorPayload
// @Router       /auth/refresh-token [post]
func refreshTokenApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.RefreshTokenRequest
		if err = c.BodyParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation refresh token request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, user, err := svc.RefreshTokenService(c.UserContext(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedRefreshToken)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(data, user),
			Message: msgSuccessRefreshToken,
		})
	}
}
