BEGIN;

DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
		CREATE TYPE user_role AS ENUM ('admin', 'reader', 'partner');
	END IF;
END $$;

CREATE TABLE IF NOT EXISTS user_profiles (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	supabase_user_id UUID NOT NULL UNIQUE,
	email VARCHAR(254) NOT NULL UNIQUE,
	role user_role NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_user_profiles_single_active_admin
	ON user_profiles ((role))
	WHERE role = 'admin' AND status = 'active';

CREATE INDEX IF NOT EXISTS idx_user_profiles_role ON user_profiles(role);

COMMIT;
