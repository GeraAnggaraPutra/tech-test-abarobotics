package auth

import "context"

func (a *Auth) CheckPermissionAction(
	ctx context.Context,
	roleGUID, permissionName, actionName string,
) (exists bool, err error) {
	const statement = `
		SELECT EXISTS (
			SELECT 
				1
			FROM 
				roles r
			LEFT JOIN 
				role_permissions rp
			ON r.guid = rp.role_guid
			LEFT JOIN 
				permissions p
			ON rp.permission_guid = p.guid
			LEFT JOIN 
				role_permission_actions rpa
			ON rp.guid = rpa.role_permission_guid
			LEFT JOIN 
				actions a
			ON rpa.action_guid = a.guid
			WHERE 
				r.guid = $1
				AND p.name = $2
				AND a.name = $3
				AND COALESCE(rpa.is_checked, false) = true
		)
	`

	err = a.db.GetContext(ctx, &exists, statement, roleGUID, permissionName, actionName)

	return
}
