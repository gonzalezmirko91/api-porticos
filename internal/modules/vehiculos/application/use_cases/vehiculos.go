package usecases

import (
	"context"
	"strings"

	"rea/porticos/internal/modules/vehiculos/domain/entities"
	"rea/porticos/internal/modules/vehiculos/domain/repository"
	domainErrors "rea/porticos/pkg/errors"
)

type VehiculosUseCase struct {
	repo repository.VehiculoRepository
}

func NewVehiculosUseCase(repo repository.VehiculoRepository) *VehiculosUseCase {
	return &VehiculosUseCase{repo: repo}
}

func (uc *VehiculosUseCase) Create(ctx context.Context, vehiculo *entities.Vehiculo) (*entities.Vehiculo, error) {
	if vehiculo == nil {
		return nil, domainErrors.NewValidationError("VEHICULO_REQUIRED", "vehiculo es obligatorio")
	}
	if err := vehiculo.Validate(); err != nil {
		return nil, err
	}
	return uc.repo.Create(ctx, vehiculo)
}

func (uc *VehiculosUseCase) ListByOwner(ctx context.Context, ownerID string, limit, offset int) ([]entities.Vehiculo, error) {
	ownerID = strings.TrimSpace(ownerID)
	if ownerID == "" {
		return nil, domainErrors.NewValidationError("VEHICULO_OWNER_REQUIRED", "usuario no autenticado")
	}
	return uc.repo.ListByOwner(ctx, ownerID, repository.ListVehiculosFilter{Limit: limit, Offset: offset})
}

func (uc *VehiculosUseCase) GetByID(ctx context.Context, ownerID, id string) (*entities.Vehiculo, error) {
	if strings.TrimSpace(ownerID) == "" || strings.TrimSpace(id) == "" {
		return nil, domainErrors.NewValidationError("VEHICULO_REQUIRED_FIELDS", "usuario e id son obligatorios")
	}
	return uc.repo.GetByID(ctx, ownerID, id)
}

func (uc *VehiculosUseCase) Update(ctx context.Context, vehiculo *entities.Vehiculo) (*entities.Vehiculo, error) {
	if vehiculo == nil {
		return nil, domainErrors.NewValidationError("VEHICULO_REQUIRED", "vehiculo es obligatorio")
	}
	if err := vehiculo.Validate(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(vehiculo.ID) == "" {
		return nil, domainErrors.NewValidationError("VEHICULO_ID_REQUIRED", "id es obligatorio")
	}
	return uc.repo.Update(ctx, vehiculo)
}

func (uc *VehiculosUseCase) Delete(ctx context.Context, ownerID, id string) error {
	if strings.TrimSpace(ownerID) == "" || strings.TrimSpace(id) == "" {
		return domainErrors.NewValidationError("VEHICULO_REQUIRED_FIELDS", "usuario e id son obligatorios")
	}
	return uc.repo.Delete(ctx, ownerID, id)
}
