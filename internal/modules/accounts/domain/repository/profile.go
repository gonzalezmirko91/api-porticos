package repository

import (
	"context"

	"rea/porticos/internal/modules/accounts/domain/entities"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *entities.UserProfile) (*entities.UserProfile, error)
	HasActiveAdmin(ctx context.Context) (bool, error)
}
