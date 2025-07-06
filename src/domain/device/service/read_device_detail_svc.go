package service

import (
	"context"

	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) ReadDeviceDetailService(
	ctx context.Context,
	guid string,
) (data model.Device, err error) {
	q := query.NewQuery(s.db)

	data, err = q.ReadDetailDeviceQuery(ctx, guid)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error read device detail query", "guid", guid)
		return
	}

	return
}
