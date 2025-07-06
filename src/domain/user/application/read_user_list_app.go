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

// @Summary Get list of users
// @Description Get list of users with search & pagination
// @Tags User
// @Accept json
// @Produce json
// @Param search query string false "Search keyword"
// @Param sort query string false "Sort field (email, role_name, created_at, updated_at)"
// @Param direction query string false "Sort direction (ASC or DESC)"
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Success 200 {object} kernel.responsePaginatePayload
// @Failure 400 {object} kernel.responseErrorPayload
// @Router /users [get]
func readUserListApp(svc *service.Service, validate *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request payload.ReadUserListRequest
		if err = c.QueryParser(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := validate.Validate(request); err != nil {
			logger.PrintError(err, "error validation read users request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, totalData, err := svc.ReadUserListService(c.UserContext(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetUserList)
		}

		return kernel.ResponsePaginate(c, request.PaginationPayload, totalData, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToReadListUserResponses(data),
			Message: msgSuccessGetUserList,
		})
	}
}
