package service

import (
	"context"

	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) DeleteDeviceService(
	ctx context.Context,
	guid string,
) (err error) {
	q := query.NewQuery(s.db)

	err = q.DeleteDeviceQuery(ctx, guid)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error delete device query", "guid", guid)
		return
	}

	return
}
