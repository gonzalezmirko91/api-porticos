package entities

import (
	"strings"
	"unicode"

	domainErrors "rea/porticos/pkg/errors"
)

type Vehiculo struct {
	ID                  string `json:"id"`
	OwnerSupabaseUserID string `json:"ownerSupabaseUserId"`
	Patente             string `json:"patente"`
	TipoVehiculo        string `json:"tipoVehiculo"`
	Alias               string `json:"alias,omitempty"`
	Activo              bool   `json:"activo"`
}

func (v *Vehiculo) Validate() error {
	if strings.TrimSpace(v.OwnerSupabaseUserID) == "" {
		return domainErrors.NewValidationError("VEHICULO_OWNER_REQUIRED", "ownerSupabaseUserId es obligatorio")
	}
	if strings.TrimSpace(v.Patente) == "" {
		return domainErrors.NewValidationError("VEHICULO_PATENTE_REQUIRED", "patente es obligatoria")
	}
	if strings.TrimSpace(v.TipoVehiculo) == "" {
		return domainErrors.NewValidationError("VEHICULO_TIPO_REQUIRED", "tipoVehiculo es obligatorio")
	}
	v.Patente = NormalizePatente(v.Patente)
	if len(v.Patente) < 4 || len(v.Patente) > 16 {
		return domainErrors.NewValidationError("VEHICULO_PATENTE_INVALID", "patente inválida")
	}
	return nil
}

func NormalizePatente(in string) string {
	in = strings.ToUpper(strings.TrimSpace(in))
	var b strings.Builder
	b.Grow(len(in))
	for _, r := range in {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
