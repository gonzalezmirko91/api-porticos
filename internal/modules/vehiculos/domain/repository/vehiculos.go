package repository

import (
	"context"

	"rea/porticos/internal/modules/vehiculos/domain/entities"
)

type ListVehiculosFilter struct {
	Limit  int
	Offset int
}

type VehiculoRepository interface {
	Create(ctx context.Context, vehiculo *entities.Vehiculo) (*entities.Vehiculo, error)
	ListByOwner(ctx context.Context, ownerID string, filter ListVehiculosFilter) ([]entities.Vehiculo, error)
	ListAll(ctx context.Context, filter ListVehiculosFilter) ([]entities.Vehiculo, error)
	GetByID(ctx context.Context, ownerID, id string) (*entities.Vehiculo, error)
	GetByIDAny(ctx context.Context, id string) (*entities.Vehiculo, error)
	Update(ctx context.Context, vehiculo *entities.Vehiculo) (*entities.Vehiculo, error)
	Delete(ctx context.Context, ownerID, id string) error
}
