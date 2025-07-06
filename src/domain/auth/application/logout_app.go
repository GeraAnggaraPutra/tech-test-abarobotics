package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/auth/service"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/kernel"
)

// @Summary      Logout
// @Description  Logout user and invalidate current session (need valid JWT)
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Failure      401  {object}  kernel.responseErrorPayload
// @Security     BearerAuth
// @Router       /auth/logout [post]
func logoutApp(svc *service.Service) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		ah, err := auth.GetAuth(c)
		if err != nil {
			return
		}

		err = svc.LogoutService(c.UserContext(), ah.GetClaims())
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedLogout)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessLogout,
		})
	}
}
