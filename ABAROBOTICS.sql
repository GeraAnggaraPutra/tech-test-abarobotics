CREATE TABLE "roles" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

CREATE TABLE "permissions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

CREATE TABLE "actions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

CREATE TABLE "role_permissions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "role_guid" varchar NOT NULL,
  "permission_guid" varchar NOT NULL
);

CREATE TABLE "role_permission_actions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "role_permission_guid" varchar NOT NULL,
  "action_guid" varchar NOT NULL,
  "is_checked" bool NOT NULL DEFAULT false
);

CREATE TABLE "users" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role_guid" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

CREATE TABLE "sessions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "user_guid" varchar NOT NULL,
  "access_token" text NOT NULL,
  "access_token_expired_at" timestamp NOT NULL,
  "refresh_token" text NOT NULL,
  "refresh_token_expired_at" timestamp NOT NULL,
  "ip_address" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "devices" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "location" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_guid") REFERENCES "roles" ("guid");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_guid") REFERENCES "permissions" ("guid");

ALTER TABLE "role_permission_actions" ADD FOREIGN KEY ("role_permission_guid") REFERENCES "role_permissions" ("guid");

ALTER TABLE "role_permission_actions" ADD FOREIGN KEY ("action_guid") REFERENCES "actions" ("guid");

ALTER TABLE "users" ADD FOREIGN KEY ("role_guid") REFERENCES "roles" ("guid");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_guid") REFERENCES "users" ("guid");
