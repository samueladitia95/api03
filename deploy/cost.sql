-- Deploy api03:cost to pg

BEGIN;

-- XXX Add DDLs here.
CREATE TABLE api03.cost (
    cost_id SERIAL PRIMARY KEY,
    from_code TEXT,
    to_code TEXT,
    provider TEXT,
    service TEXT,
    service_description TEXT,
    etd TEXT,
    cost INT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
