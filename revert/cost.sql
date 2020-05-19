-- Revert api03:cost from pg

BEGIN;

-- XXX Add DDLs here.
DROP TABLE api03.cost;

COMMIT;
