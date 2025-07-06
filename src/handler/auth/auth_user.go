package auth

import (
	"context"

	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (a *Auth) XX(ctx context.Context) (data model.User, err error) {
	q := query.NewQuery(a.db)

	data, err = q.ReadDetailUserQuery(ctx, a.claims.UserGUID)
	if err != nil {
		logger.PrintError(err, "error find user", "id", a.claims.UserGUID)
		return
	}

	return
}
