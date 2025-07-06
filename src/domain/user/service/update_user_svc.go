package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/user/payload"
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) UpdateUserService(
	ctx context.Context,
	request payload.UpdateUserRequest,
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

	var password string

	if request.Password != "" {
		password, err = util.GenerateHashPassword(request.Password)
		if err != nil {
			logger.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
			return
		}
	}

	err = q.UpdateUserQuery(ctx, request.ToParams(userGUID, password))
	if err != nil {
		logger.WithContext(ctx).Error(err, "error update user query", "request", request)
		return
	}

	if err = tx.Commit(); err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}

func (s *Service) IsEmailExistsExcludeUserService(
	ctx context.Context,
	email, userGUID string,
) (exists bool, err error) {
	q := query.NewQuery(s.db)

	exists, err = q.IsUpdateUserEmailExistsQuery(ctx, email, userGUID)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error checking email existence", "email", email, "user guid", userGUID)
		return
	}

	return
}
