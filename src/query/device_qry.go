package query

import (
	"context"
)

type CreateDeviceParams struct {
	GUID      string
	Name      string
	Location  string
	Status    string
	CreatedBy string
}

func (q *Query) CreateDeviceQuery(
	ctx context.Context,
	args CreateDeviceParams,
) (err error) {
	const statement = `
		INSERT INTO devices (
			guid, name,
			location, status,
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
		args.Name,
		args.Location,
		args.Status,
		args.CreatedBy,
	)

	return
}

type UpdateDeviceParams struct {
	GUID      string
	Name      string
	Location  string
	Status    string
	UpdatedBy string
}

func (q *Query) UpdateDeviceQuery(
	ctx context.Context,
	args UpdateDeviceParams,
) (err error) {
	const statement = `
		UPDATE
			devices
		SET
			name = CASE WHEN $2 <> '' THEN $2 ELSE name END,
			location = CASE WHEN $3 <> '' THEN $3 ELSE location END,
			status = CASE WHEN $4 <> '' THEN $4 ELSE status END,
			updated_by = $5,
			updated_at = (now() at time zone 'UTC')::TIMESTAMP
		WHERE
			guid = $1
	`

	_, err = q.db.ExecContext(ctx, statement,
		args.GUID,
		args.Name,
		args.Location,
		args.Status,
		args.UpdatedBy,
	)

	return
}

func (q *Query) DeleteDeviceQuery(
	ctx context.Context,
	guid string,
) (err error) {
	const statement = `
		DELETE FROM devices
		WHERE
			guid = $1
	`

	_, err = q.db.ExecContext(ctx, statement, guid)

	return
}
