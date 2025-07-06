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

// @Summary      Delete user
// @Description  Delete user by GUID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        guid   path      string  true  "User GUID"
// @Success      200  {object}  kernel.responseDataPayload
// @Failure      400  {object}  kernel.responseErrorPayload
// @Router       /users/{guid} [delete]
func deleteUserApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.GUIDRequest
		if err = c.ParamsParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation delete user request")
			return kernel.ResponseErrorValidate(c, err)
		}

		err = svc.DeleteUserService(c.UserContext(), request.GUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedDeleteUser)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteUser,
		})
	}
}
