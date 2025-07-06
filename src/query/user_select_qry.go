package query

import (
	"context"

	"abarobotics-test/src/model"
)

type ReadListUserParams struct {
	SetSearch bool
	Search    string
	Order     string
	Limit     int
	Offset    int
}

func (q *Query) ReadListUserQuery(
	ctx context.Context,
	args ReadListUserParams,
) (data []model.User, err error) {
	const stmt = `
		SELECT
			u.guid, u.email,
			u.role_guid, r.name AS role_name,
			u.created_at, u.created_by,
			u.updated_at, u.updated_by
		FROM
			users u
		LEFT JOIN
			roles r
		ON u.role_guid = r.guid
		WHERE
			(CASE WHEN $1::bool THEN(
				u.email ILIKE $2
				OR r.name ILIKE $2
			) ELSE TRUE END)
		ORDER BY
			(CASE WHEN $3 = 'email ASC' THEN u.email END) ASC,
			(CASE WHEN $3 = 'email DESC' THEN u.email END) DESC,
			(CASE WHEN $3 = 'role_name ASC' THEN r.name END) ASC,
			(CASE WHEN $3 = 'role_name DESC' THEN r.name END) DESC,
			(CASE WHEN $3 = 'created_at ASC' THEN u.created_at END) ASC,
			(CASE WHEN $3 = 'created_at DESC' THEN u.created_at END) DESC,
			(CASE WHEN $3 = 'updated_at ASC' THEN u.updated_at END) ASC,
			(CASE WHEN $3 = 'updated_at DESC' THEN u.updated_at END) DESC
		LIMIT $4
		OFFSET $5
	`

	rows, err := q.db.QueryxContext(ctx, stmt,
		args.SetSearch,
		args.Search,
		args.Order,
		args.Limit,
		args.Offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User

		if err = rows.Scan(
			&u.GUID,
			&u.Email,
			&u.RoleGUID,
			&u.RoleName,
			&u.CreatedAt,
			&u.CreatedBy,
			&u.UpdatedAt,
			&u.UpdatedBy,
		); err != nil {
			return
		}

		data = append(data, u)
	}

	return
}

func (q *Query) GetCountUserQuery(
	ctx context.Context,
	args ReadListUserParams,
) (count int64, err error) {
	const stmt = `
		SELECT
			COUNT(u.guid)
		FROM
			users u
		LEFT JOIN
			roles r
		ON u.role_guid = r.guid
		WHERE
			(CASE WHEN $1::bool THEN(
				u.email ILIKE $2
				OR r.name ILIKE $2
			) ELSE TRUE END)
	`

	err = q.db.GetContext(ctx, &count, stmt,
		args.SetSearch,
		args.Search,
	)

	return
}

func (q *Query) ReadDetailUserQuery(
	ctx context.Context,
	guid string,
) (data model.User, err error) {
	const statement = `
		SELECT
			u.guid, u.email,
			u.role_guid, r.name AS role_name,
			u.created_at, u.created_by,
			u.updated_at, u.updated_by,
			COALESCE(
				json_agg(
					json_build_object(
						'guid', p.guid, 
						'name', p.name, 
						'actions', actions
					) ORDER BY p.name ASC
				), 
			'[]') AS permissions
		FROM
			users u
		LEFT JOIN
			roles r
		ON u.role_guid = r.guid
		LEFT JOIN
			role_permissions rp 
		ON r.guid = rp.role_guid
		LEFT JOIN
			permissions p 
		ON rp.permission_guid = p.guid
		LEFT JOIN (
			SELECT
				rpa.role_permission_guid,
				json_agg(
					json_build_object(
						'guid', a.guid, 
						'name', a.name, 
						'is_checked', COALESCE(rpa.is_checked, false),
						'role_permission_guid', rpa.role_permission_guid
					) ORDER BY a.name
				) AS actions
			FROM
				role_permission_actions rpa
			LEFT JOIN
				actions a ON rpa.action_guid = a.guid
			GROUP BY
				rpa.role_permission_guid
			) AS role_permission_actions 
		ON rp.guid = role_permission_actions.role_permission_guid
		WHERE
			u.guid = $1
		GROUP BY
			u.guid, u.email,
			u.role_guid, r.name,
			u.created_at, u.created_by,
			u.updated_at, u.updated_by
	`

	err = q.db.QueryRowxContext(ctx, statement, guid).Scan(
		&data.GUID,
		&data.Email,
		&data.RoleGUID,
		&data.RoleName,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
		&data.Permissions,
	)

	return
}

func (q *Query) ReadDetailUserByEmailQuery(
	ctx context.Context,
	email string,
) (data model.User, err error) {
	const statement = `
		SELECT
			u.guid, u.email, u.password, u.role_guid, r.name AS role_name, 
			u.created_at, u.created_by, u.updated_at, u.updated_by
		FROM
			users u
		LEFT JOIN
			roles r
		ON u.role_guid = r.guid
		WHERE
			u.email = $1
	`

	err = q.db.QueryRowxContext(ctx, statement, email).Scan(
		&data.GUID,
		&data.Email,
		&data.Password,
		&data.RoleGUID,
		&data.RoleName,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)

	return
}

func (q *Query) IsUserEmailExistsQuery(
	ctx context.Context,
	email string,
) (exists bool, err error) {
	const statement = `
		SELECT EXISTS (
			SELECT
				1
			FROM
				users
			WHERE
				email = $1
		)
	`

	err = q.db.GetContext(ctx, &exists, statement, email)

	return
}

func (q *Query) IsUpdateUserEmailExistsQuery(
	ctx context.Context,
	email, guid string,
) (exists bool, err error) {
	statement := `
		SELECT EXISTS (
			SELECT
				1
			FROM
				users
			WHERE
				email = $1
				AND guid != $2
		)
	`

	err = q.db.GetContext(ctx, &exists, statement, email, guid)

	return
}
