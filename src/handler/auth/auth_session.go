package auth

import (
	"context"
	"time"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (a *Auth) ValidateSession(ctx context.Context) (err error) {
	q := query.NewQuery(a.db)

	session, err := q.ReadDetailSessionQuery(ctx, a.claims.GUID)
	if err != nil {
		logger.PrintError(err, "error read session by guid : "+a.claims.GUID)
		return
	}

	if time.Now().After(session.AccessTokenExpiredAt) {
		err = constant.ErrTokenExpired
		return
	}

	return
}
