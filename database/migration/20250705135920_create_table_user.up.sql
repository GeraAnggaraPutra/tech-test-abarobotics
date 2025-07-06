CREATE TABLE IF NOT EXISTS "users" (
    "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "role_guid" varchar NOT NULL REFERENCES roles(guid) ON DELETE CASCADE,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "created_by" varchar NOT NULL DEFAULT 'system',
    "updated_at" timestamp,
    "updated_by" varchar
);

INSERT INTO 
    users (email, password, role_guid)
VALUES 
    (
        'admin@gmail.com', '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC', (SELECT guid FROM roles WHERE name = 'Admin')
    ), 
    (
        'technician@gmail.com', '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC', (SELECT guid FROM roles WHERE name = 'Technician')
    ), 
    (
        'viewer@gmail.com', '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC', (SELECT guid FROM roles WHERE name = 'Viewer')
    ) 
ON CONFLICT DO NOTHING;