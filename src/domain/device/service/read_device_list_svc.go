package service

import (
	"context"

	"abarobotics-test/src/domain/device/payload"
	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) ReadDeviceListService(
	ctx context.Context,
	request payload.ReadDeviceListRequest,
) (data []model.Device, totalData int64, err error) {
	q := query.NewQuery(s.db)

	data, err = q.ReadListDeviceQuery(ctx, request.ToParams())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error read device list query", "request", request)
		return
	}

	totalData, err = q.GetCountDeviceQuery(ctx, request.ToParams())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error get count device query", "request", request)
		return
	}

	return
}
