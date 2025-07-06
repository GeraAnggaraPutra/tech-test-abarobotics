package query

import (
	"context"

	"abarobotics-test/src/model"
)

func (q *Query) ReadDetailSessionQuery(
	ctx context.Context,
	guid string,
) (data model.Session, err error) {
	const statement = `
		SELECT
			guid, user_guid,
			access_token, access_token_expired_at,
			refresh_token, refresh_token_expired_at
		FROM
			sessions
		WHERE
			guid = $1
	`

	err = q.db.QueryRowxContext(ctx, statement, guid).Scan(
		&data.GUID,
		&data.UserGUID,
		&data.AccessToken,
		&data.AccessTokenExpiredAt,
		&data.RefreshToken,
		&data.RefreshTokenExpiredAt,
	)

	return
}
