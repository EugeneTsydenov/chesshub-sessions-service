CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sessions (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id       int NOT NULL,
    ip_address    varchar(45),
    device_info   text,
    is_active     boolean DEFAULT true NOT NULL,
    expired_at    timestamp NOT NULL DEFAULT (now() + interval '30 days'),
    created_at    timestamp DEFAULT now() NOT NULL,
    updated_at    timestamp DEFAULT now() NOT NULL
);