package query

import (
	"context"

)

type CreateUserParams struct {
	GUID      string
	Email     string
	Password  string
	RoleGUID  string
	CreatedBy string
}

func (q *Query) CreateUserQuery(
	ctx context.Context,
	args CreateUserParams,
) (err error) {
	const statement = `
		INSERT INTO users (
			guid, email,
			password, role_guid,
			created_by, created_at
		)
		VALUES (
			$1, $2,
			$3, $4,
			$5, (now() at time zone 'UTC')::TIMESTAMP
		)
	`

	_, err = q.db.ExecContext(ctx, statement,
		args.GUID,
		args.Email,
		args.Password,
		args.RoleGUID,
		args.CreatedBy,
	)

	return
}

type UpdateUserParams struct {
	GUID      string
	Email     string
	Password  string
	RoleGUID  string
	UpdatedBy string
}

func (q *Query) UpdateUserQuery(
	ctx context.Context,
	args UpdateUserParams,
) (err error) {
	const statement = `
		UPDATE
			users
		SET
			email = CASE WHEN $2 <> '' THEN $2 ELSE email END,
			password = CASE WHEN $3 <> '' THEN $3 ELSE password END,
			role_guid = CASE WHEN $4 <> '' THEN $4 ELSE role_guid END,
			updated_by = $5,
			updated_at = (now() at time zone 'UTC')::TIMESTAMP
		WHERE
			guid = $1
	`

	_, err = q.db.ExecContext(ctx, statement,
		args.GUID,
		args.Email,
		args.Password,
		args.RoleGUID,
		args.UpdatedBy,
	)

	return
}

func (q *Query) DeleteUserQuery(
	ctx context.Context,
	guid string,
) (err error) {
	const statement = `
		DELETE FROM users
		WHERE
			guid = $1
	`

	_, err = q.db.ExecContext(ctx, statement, guid)

	return
}
