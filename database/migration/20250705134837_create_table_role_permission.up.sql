CREATE TABLE IF NOT EXISTS "role_permissions" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "role_guid" varchar NOT NULL REFERENCES roles(guid) ON DELETE CASCADE,
    "permission_guid" varchar NOT NULL REFERENCES permissions(guid) ON DELETE CASCADE
);

INSERT INTO 
    role_permissions(role_guid, permission_guid)
VALUES
    (
        (SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'Device Management')
    ),
    (
        (SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'User Management')
    ),
    (
        (SELECT guid FROM roles WHERE name = 'Technician'), (SELECT guid FROM permissions WHERE name = 'Device Management')
    ),
    (
        (SELECT guid FROM roles WHERE name = 'Viewer'), (SELECT guid FROM permissions WHERE name = 'Device Management')
    )
ON CONFLICT DO NOTHING;