BEGIN;

DROP INDEX IF EXISTS uq_user_profiles_single_active_admin;
DROP INDEX IF EXISTS idx_user_profiles_role;
DROP TABLE IF EXISTS user_profiles;
DROP TYPE IF EXISTS user_role;

COMMIT;
