package service

import (
	"context"

	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) ReadUserDetailService(
	ctx context.Context,
	guid string,
) (data model.User, err error) {
	q := query.NewQuery(s.db)

	data, err = q.ReadDetailUserQuery(ctx, guid)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error read user detail query", "guid", guid)
		return
	}

	return
}
