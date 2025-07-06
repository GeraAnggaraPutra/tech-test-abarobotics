package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/handler/auth"
	"abarobotics-test/src/handler/database"
	"abarobotics-test/src/handler/jwt"
	"abarobotics-test/toolkit/logger"
)

type EnsureToken struct {
	auth *auth.Auth
}

func NewEnsureToken(db database.DB) *EnsureToken {
	ah := auth.NewAuth(db)
	return &EnsureToken{auth: ah}
}

func (et *EnsureToken) ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.Get(constant.DefaultMdwHeaderToken)
		token, err := parseHeaderToken(tokenHeader)
		if err != nil {
			logger.PrintError(err, "error parse header token")
			return fiber.NewError(http.StatusUnauthorized, err.Error())
		}

		accessTokenClaims, err := jwt.ClaimsAccessToken(token)
		if err != nil {
			logger.PrintError(err, "error claims access token")
			return fiber.NewError(http.StatusUnauthorized, err.Error())
		}

		et.auth.SetClaims(&accessTokenClaims)

		err = et.auth.ValidateSession(c.UserContext())
		if err != nil {
			return fiber.NewError(http.StatusUnauthorized, err.Error())
		}

		c.Locals("auth", *et.auth)
		return c.Next()
	}
}

func parseHeaderToken(headerDataToken string) (string, error) {
	if !strings.Contains(headerDataToken, "Bearer") {
		return "", constant.ErrHeaderTokenNotFound
	}

	splitToken := strings.Split(headerDataToken, fmt.Sprintf("%s ", constant.DefaultMdwHeaderBearer))
	if len(splitToken) <= 1 {
		return "", constant.ErrHeaderTokenInvalid
	}

	return splitToken[1], nil
}

func (et *EnsureToken) ValidatePermissionAction(permissionName, actionName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleGUID := et.auth.GetClaims().RoleGUID

		ok, err := et.auth.CheckPermissionAction(c.UserContext(), roleGUID, permissionName, actionName)
		if err != nil {
			logger.PrintError(err, "error check permission action query")
			return fiber.NewError(http.StatusForbidden, "failed to check permission")
		}

		if !ok {
			return fiber.NewError(http.StatusForbidden, constant.ErrForbiddenPermission.Error())
		}

		return c.Next()
	}
}
