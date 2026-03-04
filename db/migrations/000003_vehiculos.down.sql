BEGIN;

DROP INDEX IF EXISTS idx_vehiculos_owner_activo;
DROP INDEX IF EXISTS idx_vehiculos_owner;
DROP TABLE IF EXISTS vehiculos;

COMMIT;
