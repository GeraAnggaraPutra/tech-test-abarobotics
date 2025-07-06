package service

import (
	"context"

	"abarobotics-test/src/handler/jwt"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) LogoutService(
	ctx context.Context,
	claims *jwt.AccessTokenPayload,
) (err error) {
	q := query.NewQuery(s.db)

	err = q.DeleteSessionQuery(ctx, claims.GUID)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error delete session", "session_guid", claims.GUID)
		return
	}

	return
}
