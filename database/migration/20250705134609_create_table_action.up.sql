CREATE TABLE IF NOT EXISTS "actions" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar
);

INSERT INTO 
    actions(name)
VALUES
    ('create'),
    ('edit'),
    ('view'),
    ('delete')
ON CONFLICT DO NOTHING;
