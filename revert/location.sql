-- Revert api03:location from pg

BEGIN;

-- XXX Add DDLs here.
DROP TABLE api03.location;

COMMIT;
