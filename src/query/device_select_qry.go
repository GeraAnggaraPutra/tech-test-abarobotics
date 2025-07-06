package query

import (
	"context"

	"abarobotics-test/src/model"
)

type ReadListDeviceParams struct {
	SetSearch bool
	Search    string
	SetStatus bool
	Status    string
	Order     string
	Limit     int
	Offset    int
}

func (q *Query) ReadListDeviceQuery(
	ctx context.Context,
	args ReadListDeviceParams,
) (data []model.Device, err error) {
	const stmt = `
		SELECT
			guid, name, location, status,
			created_at, created_by,
			updated_at, updated_by
		FROM
			devices
		WHERE
			(CASE WHEN $1::bool THEN(
				name ILIKE $2
				OR location ILIKE $2
				OR status ILIKE $2
			) ELSE TRUE END)
			AND (CASE WHEN $3::bool THEN status = $4 ELSE TRUE END)
		ORDER BY
			(CASE WHEN $5 = 'name ASC' THEN name END) ASC,
			(CASE WHEN $5 = 'name DESC' THEN name END) DESC,
			(CASE WHEN $5 = 'location ASC' THEN location END) ASC,
			(CASE WHEN $5 = 'location DESC' THEN location END) DESC,
			(CASE WHEN $5 = 'status ASC' THEN status END) ASC,
			(CASE WHEN $5 = 'status DESC' THEN status END) DESC,
			(CASE WHEN $5 = 'created_at ASC' THEN created_at END) ASC,
			(CASE WHEN $5 = 'created_at DESC' THEN created_at END) DESC,
			(CASE WHEN $5 = 'updated_at ASC' THEN updated_at END) ASC,
			(CASE WHEN $5 = 'updated_at DESC' THEN updated_at END) DESC
		LIMIT $6
		OFFSET $7
	`

	err = q.db.SelectContext(ctx, &data, stmt,
		args.SetSearch,
		args.Search,
		args.SetStatus,
		args.Status,
		args.Order,
		args.Limit,
		args.Offset,
	)

	rows, err := q.db.QueryxContext(ctx, stmt,
		args.SetSearch,
		args.Search,
		args.SetStatus,
		args.Status,
		args.Order,
		args.Limit,
		args.Offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d model.Device

		if err = rows.Scan(
			&d.GUID,
			&d.Name,
			&d.Location,
			&d.Status,
			&d.CreatedAt,
			&d.CreatedBy,
			&d.UpdatedAt,
			&d.UpdatedBy,
		); err != nil {
			return
		}

		data = append(data, d)
	}

	return
}

func (q *Query) GetCountDeviceQuery(
	ctx context.Context,
	args ReadListDeviceParams,
) (count int64, err error) {
	const stmt = `
		SELECT
			COUNT(*)
		FROM
			devices
		WHERE
			(CASE WHEN $1::bool THEN(
				name ILIKE $2
				OR location ILIKE $2
				OR status ILIKE $2
			) ELSE TRUE END)
			AND (CASE WHEN $3::bool THEN status = $4 ELSE TRUE END)
	`

	err = q.db.GetContext(ctx, &count, stmt,
		args.SetSearch,
		args.Search,
		args.SetStatus,
		args.Status,
	)

	return
}

func (q *Query) ReadDetailDeviceQuery(
	ctx context.Context,
	guid string,
) (data model.Device, err error) {
	const statement = `
		SELECT
			guid, name, location, status,
			created_at, created_by,
			updated_at, updated_by
		FROM
			devices
		WHERE
			guid = $1
	`

	err = q.db.QueryRowxContext(ctx, statement, guid).Scan(
		&data.GUID,
		&data.Name,
		&data.Location,
		&data.Status,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)

	return
}
