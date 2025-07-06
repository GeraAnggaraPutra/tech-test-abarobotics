package auth

import (
	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
	"context"
)

func (a *Auth) User(ctx context.Context) (data model.User, err error) {
	q := query.NewQuery(a.db)

	data, err = q.ReadDetailUserQuery(ctx, a.claims.UserGUID)
	if err != nil {
		logger.PrintError(err, "error find user", "id", a.claims.UserGUID)
		return
	}

	return
}
