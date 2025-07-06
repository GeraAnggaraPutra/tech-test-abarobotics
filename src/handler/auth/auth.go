package auth

import (
	"github.com/gofiber/fiber/v2"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/handler/database"
	"abarobotics-test/src/handler/jwt"
)

type Auth struct {
	db     database.DB
	claims *jwt.AccessTokenPayload
}

func NewAuth(db database.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func GetAuth(c *fiber.Ctx) (*Auth, error) {
	a, ok := c.Locals("auth").(Auth)
	if !ok {
		return nil, constant.ErrTokenUnauthorized
	}

	return &a, nil
}

func (a *Auth) GetClaims() *jwt.AccessTokenPayload {
	return a.claims
}

func (a *Auth) SetClaims(claims *jwt.AccessTokenPayload) {
	a.claims = claims
}
