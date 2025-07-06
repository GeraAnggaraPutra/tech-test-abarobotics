package service

import (
	"context"

	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) DeleteUserService(
	ctx context.Context,
	guid string,
) (err error) {
	q := query.NewQuery(s.db)

	err = q.DeleteUserQuery(ctx, guid)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error delete user query", "guid", guid)
		return
	}

	return
}
