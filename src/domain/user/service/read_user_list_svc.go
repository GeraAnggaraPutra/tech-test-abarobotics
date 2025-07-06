package service

import (
	"context"

	"abarobotics-test/src/domain/user/payload"
	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) ReadUserListService(
	ctx context.Context,
	request payload.ReadUserListRequest,
) (data []model.User, totalData int64, err error) {
	q := query.NewQuery(s.db)

	data, err = q.ReadListUserQuery(ctx, request.ToParams())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error read user list query", "request", request)
		return
	}

	totalData, err = q.GetCountUserQuery(ctx, request.ToParams())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error get count user query", "request", request)
		return
	}

	return
}
