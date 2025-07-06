package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/user/payload"
	"abarobotics-test/src/domain/user/service"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/logger"
)

// @Summary      Get user profile
// @Description  Get profile of a user logged in
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Router       /users/me [get]
func readProfileApp(svc *service.Service) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error get auth handler")
			return
		}

		data, err := svc.ReadUserDetailService(c.UserContext(), ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetUserProfile)
		}

		response, err := payload.ToReadDetailUserResponse(data)
		if err != nil {
			logger.WithContext(c.UserContext()).Error(err, "error parse response")
			return
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    response,
			Message: msgSuccessGetUserProfile,
		})
	}
}
