-- Deploy api03:location to pg

BEGIN;

-- XXX Add DDLs here.
CREATE TABLE api03.location (
    location_id SERIAL PRIMARY KEY,
    location_code TEXT,
    provider_name TEXT,
    provider_code TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
