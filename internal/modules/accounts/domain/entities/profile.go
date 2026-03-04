package entities

import (
	"strings"

	domainErrors "rea/porticos/pkg/errors"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleReader  Role = "reader"
	RolePartner Role = "partner"
)

type UserProfile struct {
	ID             string `json:"id"`
	SupabaseUserID string `json:"supabaseUserId"`
	Email          string `json:"email"`
	Role           Role   `json:"role"`
	Status         string `json:"status"`
}

func (p *UserProfile) Validate() error {
	if strings.TrimSpace(p.SupabaseUserID) == "" {
		return domainErrors.NewValidationError("PROFILE_USER_ID_REQUIRED", "supabaseUserId es obligatorio")
	}
	if strings.TrimSpace(p.Email) == "" {
		return domainErrors.NewValidationError("PROFILE_EMAIL_REQUIRED", "email es obligatorio")
	}
	if !IsValidRole(string(p.Role)) {
		return domainErrors.NewValidationError("PROFILE_ROLE_INVALID", "rol inválido")
	}
	if strings.TrimSpace(p.Status) == "" {
		p.Status = "active"
	}
	return nil
}

func IsValidRole(role string) bool {
	switch strings.ToLower(strings.TrimSpace(role)) {
	case string(RoleAdmin), string(RoleReader), string(RolePartner):
		return true
	default:
		return false
	}
}
