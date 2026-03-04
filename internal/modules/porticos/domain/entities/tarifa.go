package entities

import (
	domainErrors "rea/porticos/pkg/errors"
	"strings"
	"time"
)

type Tarifa struct {
	ID           string          `json:"id"`
	TipoVehiculo string          `json:"tipoVehiculo"`
	Moneda       string          `json:"moneda"`
	Horarios     []TarifaHorario `json:"horarios"`
}

type TarifaHorario struct {
	ID     string    `json:"id"`
	Inicio time.Time `json:"inicio"`
	Fin    time.Time `json:"fin"`
	Monto  int       `json:"monto"`
}

func (t *Tarifa) Validate() error {
	if strings.TrimSpace(t.TipoVehiculo) == "" {
		return domainErrors.NewValidationError("TARIFA_TIPO_VEHICULO_REQUIRED", "tipoVehiculo es obligatorio")
	}
	if strings.TrimSpace(t.Moneda) == "" {
		return domainErrors.NewValidationError("TARIFA_MONEDA_REQUIRED", "moneda es obligatoria")
	}
	if len(t.Horarios) == 0 {
		return domainErrors.NewValidationError("TARIFA_HORARIOS_REQUIRED", "al menos un horario es obligatorio")
	}

	for i := range t.Horarios {
		if err := t.Horarios[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (h *TarifaHorario) Validate() error {
	if !h.Inicio.Before(h.Fin) {
		return domainErrors.NewValidationError("TARIFA_HORARIO_RANGE_INVALID", "inicio debe ser menor que fin")
	}
	if h.Monto <= 0 {
		return domainErrors.NewValidationError("TARIFA_MONTO_INVALID", "monto debe ser mayor que 0")
	}
	return nil
}
