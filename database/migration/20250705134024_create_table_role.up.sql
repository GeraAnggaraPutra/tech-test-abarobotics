CREATE TABLE IF NOT EXISTS "roles" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "name" varchar UNIQUE NOT NULL,
    "description" varchar,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar
);

INSERT INTO 
    roles (name, description)
VALUES 
    (
        'Admin', 'This is Admin Role'
    ), 
    (
        'Technician', 'This is Technician Role'
    ), 
    (
        'Viewer', 'This is Viewer Role'
    ) 
ON CONFLICT DO NOTHING;