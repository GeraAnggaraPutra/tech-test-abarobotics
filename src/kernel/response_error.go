package kernel

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/handler/validator"
	"abarobotics-test/src/util"
)

type responseErrorPayload struct {
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

func ResponseErrorValidate(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusBadRequest).JSON(responseErrorPayload{
		Error:   validator.ValidationErrors(err),
		Message: constant.ErrMsgValidate,
	})
}

func ResponseError(c *fiber.Ctx, err error, msg string) error {
	e := formatError(err)
	if e != nil {
		return c.Status(http.StatusBadRequest).JSON(responseErrorPayload{
			Error:   e,
			Message: msg,
		})
	}

	// Kalau tidak termasuk error custom, kirim error default
	return c.Status(http.StatusInternalServerError).JSON(responseErrorPayload{
		Message: msg,
	})
}

func formatError(err error) (e map[string]string) {
	switch {
	case errors.Is(err, constant.ErrGUID):
		e = map[string]string{"guid": util.CapitalFirstLetter(err.Error())}
	case errors.Is(err, constant.ErrPasswordIncorrect):
		e = map[string]string{"password": util.CapitalFirstLetter(err.Error())}
	case errors.Is(err, constant.ErrAccountNotFound) || errors.Is(err, constant.ErrEmailAlreadyExists):
		e = map[string]string{"email": util.CapitalFirstLetter(err.Error())}
	}

	return
}
