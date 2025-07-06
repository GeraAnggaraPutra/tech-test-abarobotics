CREATE TABLE IF NOT EXISTS "devices" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "name" varchar NOT NULL,
    "location" varchar NOT NULL,
    "status" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar
);