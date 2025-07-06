package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/auth/helper"
	"abarobotics-test/src/domain/auth/payload"
	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) LoginService(
	ctx context.Context,
	request payload.LoginRequest,
) (data model.Session, user model.User, err error) {
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

	user, err = q.ReadDetailUserByEmailQuery(ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = logger.PrintNewError(err, constant.ErrAccountNotFound)
			return
		}

		logger.WithContext(ctx).Error(err, "error read user by email : "+request.Email)
		return
	}

	if err = util.CompareHashPassword(request.Password, user.Password); err != nil {
		err = logger.PrintNewError(err, constant.ErrPasswordIncorrect)
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(user.GUID))
	if err != nil {
		logger.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(user.GUID))
		return
	}

	err = q.CreateSessionQuery(ctx, query.SessionParams{
		GUID:                  data.GUID,
		UserGUID:              data.UserGUID,
		AccessToken:           data.AccessToken,
		AccessTokenExpiredAt:  data.AccessTokenExpiredAt,
		RefreshToken:          data.RefreshToken,
		RefreshTokenExpiredAt: data.RefreshTokenExpiredAt,
		IPAddress:             data.IPAddress,
		UserAgent:             data.UserAgent,
	})
	if err != nil {
		logger.WithContext(ctx).Error(err, "error create session", "session model", data)
		return
	}

	if err = tx.Commit(); err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
