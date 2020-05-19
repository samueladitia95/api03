-- Verify api03:location on pg

BEGIN;

-- XXX Add verifications here.
SELECT location_id, location_code, provider_name, provider_code, created_at
FROM api03.location
WHERE FALSE;

ROLLBACK;
