package application

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/domain/auth/service"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/kernel"
)

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
