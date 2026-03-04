package data

import (
	"context"
	"errors"

	"rea/porticos/internal/modules/accounts/domain/entities"
	"rea/porticos/internal/modules/accounts/domain/repository"
	domainErrors "rea/porticos/pkg/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfilePostgresRepository struct {
	pool *pgxpool.Pool
}

func NewProfilePostgresRepository(pool *pgxpool.Pool) repository.ProfileRepository {
	return &ProfilePostgresRepository{pool: pool}
}

func (r *ProfilePostgresRepository) Create(ctx context.Context, profile *entities.UserProfile) (*entities.UserProfile, error) {
	if profile == nil {
		return nil, domainErrors.NewValidationError("PROFILE_REQUIRED", "profile es obligatorio")
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}

	err := r.pool.QueryRow(ctx, `
		INSERT INTO user_profiles (supabase_user_id, email, role, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id::text
	`, profile.SupabaseUserID, profile.Email, profile.Role, profile.Status).Scan(&profile.ID)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, domainErrors.NewConflictError("PROFILE_CONFLICT", "email o perfil ya registrado, o ya existe un admin activo")
		}
		return nil, domainErrors.NewInternalError("PROFILE_CREATE_ERROR", "error al crear profile")
	}
	return profile, nil
}

func (r *ProfilePostgresRepository) HasActiveAdmin(ctx context.Context) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM user_profiles
			WHERE role = 'admin' AND status = 'active'
		)
	`).Scan(&exists)
	if err != nil {
		return false, domainErrors.NewInternalError("PROFILE_ADMIN_CHECK_ERROR", "error validando admin activo")
	}
	return exists, nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}
	return pgErr.Code == "23505"
}
