-- Revert api03:appschema from pg

BEGIN;

-- XXX Add DDLs here.
DROP SCHEMA api03;

COMMIT;
