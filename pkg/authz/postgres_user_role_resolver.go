package authz

import (
	"context"
	"strings"

	domainErrors "rea/porticos/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRoleResolver struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRoleResolver(pool *pgxpool.Pool) *PostgresUserRoleResolver {
	return &PostgresUserRoleResolver{pool: pool}
}

func (r *PostgresUserRoleResolver) ResolveRole(ctx context.Context, supabaseUserID string) (string, error) {
	supabaseUserID = strings.TrimSpace(supabaseUserID)
	if supabaseUserID == "" {
		return "", domainErrors.NewForbiddenError("ROLE_SUBJECT_REQUIRED", "usuario no autorizado")
	}

	var role string
	err := r.pool.QueryRow(ctx, `
		SELECT role::text
		FROM user_profiles
		WHERE supabase_user_id = $1
		  AND status = 'active'
		LIMIT 1
	`, supabaseUserID).Scan(&role)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", domainErrors.NewForbiddenError("ROLE_PROFILE_NOT_FOUND", "perfil no autorizado")
		}
		return "", domainErrors.NewInternalError("ROLE_RESOLUTION_ERROR", "error resolviendo permisos del usuario")
	}

	role = strings.ToLower(strings.TrimSpace(role))
	if role == "" {
		return "", domainErrors.NewForbiddenError("ROLE_INVALID", "perfil no autorizado")
	}

	return role, nil
}
