package requests

import (
	"strings"

	"rea/porticos/internal/modules/vehiculos/domain/entities"
	domainErrors "rea/porticos/pkg/errors"
)

type CreateVehiculoRequest struct {
	Patente      string `json:"patente"`
	TipoVehiculo string `json:"tipoVehiculo"`
	Alias        string `json:"alias,omitempty"`
	Activo       *bool  `json:"activo,omitempty"`
}

type UpdateVehiculoRequest struct {
	Patente      string `json:"patente"`
	TipoVehiculo string `json:"tipoVehiculo"`
	Alias        string `json:"alias,omitempty"`
	Activo       *bool  `json:"activo,omitempty"`
}

func (r *CreateVehiculoRequest) ToEntity(ownerID string) (*entities.Vehiculo, error) {
	ownerID = strings.TrimSpace(ownerID)
	if ownerID == "" {
		return nil, domainErrors.NewValidationError("VEHICULO_OWNER_REQUIRED", "usuario no autenticado")
	}

	activo := true
	if r.Activo != nil {
		activo = *r.Activo
	}

	out := &entities.Vehiculo{
		OwnerSupabaseUserID: ownerID,
		Patente:             r.Patente,
		TipoVehiculo:        strings.ToLower(strings.TrimSpace(r.TipoVehiculo)),
		Alias:               strings.TrimSpace(r.Alias),
		Activo:              activo,
	}
	if err := out.Validate(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *UpdateVehiculoRequest) ToEntity(ownerID, id string) (*entities.Vehiculo, error) {
	ownerID = strings.TrimSpace(ownerID)
	id = strings.TrimSpace(id)
	if ownerID == "" || id == "" {
		return nil, domainErrors.NewValidationError("VEHICULO_REQUIRED_FIELDS", "usuario e id son obligatorios")
	}

	activo := true
	if r.Activo != nil {
		activo = *r.Activo
	}

	out := &entities.Vehiculo{
		ID:                  id,
		OwnerSupabaseUserID: ownerID,
		Patente:             r.Patente,
		TipoVehiculo:        strings.ToLower(strings.TrimSpace(r.TipoVehiculo)),
		Alias:               strings.TrimSpace(r.Alias),
		Activo:              activo,
	}
	if err := out.Validate(); err != nil {
		return nil, err
	}
	return out, nil
}
