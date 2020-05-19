-- Verify api03:appschema on pg

BEGIN;

-- XXX Add verifications here.
SELECT pg_catalog.has_schema_privilege('api03', 'usage');

ROLLBACK;
