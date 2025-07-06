CREATE TABLE IF NOT EXISTS "role_permission_actions" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "role_permission_guid" varchar NOT NULL REFERENCES "role_permissions"("guid") ON DELETE CASCADE,
    "action_guid" varchar NOT NULL REFERENCES "actions"("guid") ON DELETE CASCADE,
    "is_checked" bool NOT NULL DEFAULT false
);

INSERT INTO 
    role_permission_actions(role_permission_guid, action_guid, is_checked)
VALUES
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'create'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'view'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'edit'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'delete'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), (SELECT guid FROM actions WHERE name = 'create'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), (SELECT guid FROM actions WHERE name = 'view'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), (SELECT guid FROM actions WHERE name = 'edit'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), (SELECT guid FROM actions WHERE name = 'delete'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Technician') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'create'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Technician') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'view'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Technician') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'edit'), true
    ),
    (
        (SELECT guid FROM role_permissions WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Viewer') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Device Management')), (SELECT guid FROM actions WHERE name = 'view'), true
    )
ON CONFLICT DO NOTHING;