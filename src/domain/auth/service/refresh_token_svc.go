package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"abarobotics-test/src/constant"
	"abarobotics-test/src/domain/auth/helper"
	"abarobotics-test/src/domain/auth/payload"
	"abarobotics-test/src/handler/jwt"
	"abarobotics-test/src/model"
	"abarobotics-test/src/query"
	"abarobotics-test/toolkit/logger"
)

func (s *Service) RefreshTokenService(
	ctx context.Context,
	request payload.RefreshTokenRequest,
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

	refreshTokenClaims, err := jwt.ClaimsRefreshToken(request.RefreshToken)
	if err != nil {
		logger.PrintError(err, "error claims refresh token : "+request.RefreshToken)
		return
	}

	session, err := q.ReadDetailSessionQuery(ctx, refreshTokenClaims.GUID)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error find session by GUID : ", refreshTokenClaims.GUID)
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(session))
	if err != nil {
		logger.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(session))
		return
	}

	user, err = q.ReadDetailUserQuery(ctx, session.UserGUID)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error find user", "guid", session.UserGUID)
		return
	}

	err = q.UpdateSessionQuery(ctx, query.SessionParams{
		GUID:                  data.GUID,
		AccessToken:           data.AccessToken,
		AccessTokenExpiredAt:  data.AccessTokenExpiredAt,
		RefreshToken:          data.RefreshToken,
		RefreshTokenExpiredAt: data.RefreshTokenExpiredAt,
	})
	if err != nil {
		logger.WithContext(ctx).Error(err, "error update session", "session model", data)
		return
	}

	if err = tx.Commit(); err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
