package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/device/payload"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) UpdateDeviceService(
	ctx context.Context,
	request payload.UpdateDeviceRequest,
	userGUID string,
) (err error) {
	tx, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.WithContext(ctx).Error(err, "failed to begin transaction")
		err = errors.WithStack(constant.ErrUnknownSource)

		return
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				logger.WithContext(ctx).Error(err, "error rollback", errRollback)
				err = errors.WithStack(constant.ErrUnknownSource)
				return
			}
		}
	}()

	q := query.NewQuery(tx)

	err = q.UpdateDeviceQuery(ctx, request.ToParams(userGUID))
	if err != nil {
		logger.WithContext(ctx).Error(err, "error update device query", "request", request)
	}

	if err = tx.Commit(); err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
