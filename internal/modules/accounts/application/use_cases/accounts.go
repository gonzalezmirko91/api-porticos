package usecases

import (
	"context"
	"strings"

	"rea/porticos/internal/modules/accounts/domain/dtos/requests"
	"rea/porticos/internal/modules/accounts/domain/entities"
	"rea/porticos/internal/modules/accounts/domain/repository"
	domainErrors "rea/porticos/pkg/errors"
)

type SupabaseAdminGateway interface {
	CreateUser(ctx context.Context, email, password string) (string, error)
	DeleteUser(ctx context.Context, userID string) error
}

type AccountsUseCase struct {
	profiles repository.ProfileRepository
	supabase SupabaseAdminGateway
}

func NewAccountsUseCase(profiles repository.ProfileRepository, supabase SupabaseAdminGateway) *AccountsUseCase {
	return &AccountsUseCase{
		profiles: profiles,
		supabase: supabase,
	}
}

func (uc *AccountsUseCase) CreateAccount(ctx context.Context, req requests.CreateAccountRequest) (*entities.UserProfile, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.Role == string(entities.RoleAdmin) {
		exists, err := uc.profiles.HasActiveAdmin(ctx)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, domainErrors.NewConflictError("ACCOUNT_ADMIN_EXISTS", "ya existe una cuenta admin activa")
		}
	}

	userID, err := uc.supabase.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	profile := &entities.UserProfile{
		SupabaseUserID: strings.TrimSpace(userID),
		Email:          req.Email,
		Role:           entities.Role(req.Role),
		Status:         "active",
	}

	created, err := uc.profiles.Create(ctx, profile)
	if err != nil {
		// Compensación best-effort para no dejar usuario huérfano en Supabase.
		_ = uc.supabase.DeleteUser(ctx, userID)
		return nil, err
	}

	return created, nil
}

func (uc *AccountsUseCase) SignupPublic(ctx context.Context, req requests.CreateAccountRequest) (*entities.UserProfile, error) {
	req.Role = string(entities.RoleReader)
	return uc.CreateAccount(ctx, req)
}

func (uc *AccountsUseCase) BootstrapFirstAdmin(ctx context.Context, req requests.CreateAccountRequest) (*entities.UserProfile, error) {
	req.Role = string(entities.RoleAdmin)
	return uc.CreateAccount(ctx, req)
}
