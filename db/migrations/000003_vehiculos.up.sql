BEGIN;

CREATE TABLE IF NOT EXISTS vehiculos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_supabase_user_id UUID NOT NULL,
    patente VARCHAR(16) NOT NULL,
    tipo_vehiculo VARCHAR(30) NOT NULL,
    alias VARCHAR(80) NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_vehiculo_owner_patente UNIQUE (owner_supabase_user_id, patente)
);

CREATE INDEX IF NOT EXISTS idx_vehiculos_owner
    ON vehiculos(owner_supabase_user_id);

CREATE INDEX IF NOT EXISTS idx_vehiculos_owner_activo
    ON vehiculos(owner_supabase_user_id, activo);

COMMIT;
