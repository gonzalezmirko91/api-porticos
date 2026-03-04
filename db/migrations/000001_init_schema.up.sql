BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS porticos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    codigo VARCHAR(30) NOT NULL UNIQUE,
    nombre VARCHAR(150) NOT NULL,
    latitude NUMERIC(9,6) NOT NULL CHECK (latitude >= -90 AND latitude <= 90),
    longitude NUMERIC(9,6) NOT NULL CHECK (longitude >= -180 AND longitude <= 180),
    bearing NUMERIC(5,2) NULL CHECK (bearing >= 0 AND bearing <= 360),
    detection_radius_meters NUMERIC(8,2) NULL CHECK (detection_radius_meters > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tarifas_portico (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    portico_id UUID NOT NULL REFERENCES porticos(id) ON DELETE CASCADE,
    tipo_vehiculo VARCHAR(30) NOT NULL,
    moneda CHAR(3) NOT NULL DEFAULT 'CLP',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_tarifa_portico_tipo UNIQUE (portico_id, tipo_vehiculo)
);

CREATE TABLE IF NOT EXISTS tarifa_horarios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tarifa_id UUID NOT NULL REFERENCES tarifas_portico(id) ON DELETE CASCADE,
    inicio TIME NOT NULL,
    fin TIME NOT NULL,
    monto INTEGER NOT NULL CHECK (monto > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_horario_rango CHECK (inicio < fin),
    CONSTRAINT uq_tarifa_horario UNIQUE (tarifa_id, inicio, fin)
);

CREATE INDEX IF NOT EXISTS idx_tarifas_portico_portico_id
    ON tarifas_portico(portico_id);

CREATE INDEX IF NOT EXISTS idx_tarifa_horarios_tarifa_id
    ON tarifa_horarios(tarifa_id);

CREATE INDEX IF NOT EXISTS idx_porticos_codigo
    ON porticos(codigo);

COMMIT;
