-- Verify api03:cost on pg

BEGIN;

-- XXX Add verifications here.
SELECT cost_id, from_code, to_code, provider, service, service_description, etd, cost, updated_at
FROM api03.location
WHERE FALSE;

ROLLBACK;
