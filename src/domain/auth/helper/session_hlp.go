package helper

import (
	"context"

	"abarobotics-test/src/domain/auth/payload"
	"abarobotics-test/src/handler/jwt"
	"abarobotics-test/src/model"
	"abarobotics-test/toolkit/logger"
)

func GenerateSessionModel(
	ctx context.Context,
	request payload.SessionPayload,
) (data model.Session, err error) {
	accessToken, err := jwt.GenerateAccessToken(request.ToAccessTokenRequest())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error generate access token")
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(request.ToRefreshTokenRequest())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error generate refresh token")
		return
	}

	data = model.Session{
		GUID:                  request.SessionGUID,
		UserGUID:              request.UserGUID,
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.ExpiresAt,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.ExpiresAt,
		IPAddress:             request.IPAddress,
		UserAgent:             request.UserAgent,
	}

	return
}
