package requests

import (
	"strings"

	"rea/porticos/internal/modules/accounts/domain/entities"
	domainErrors "rea/porticos/pkg/errors"
)

type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r *CreateAccountRequest) Validate() error {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	r.Role = strings.TrimSpace(strings.ToLower(r.Role))

	if r.Email == "" {
		return domainErrors.NewValidationError("ACCOUNT_EMAIL_REQUIRED", "email es obligatorio")
	}
	if len(r.Password) < 8 {
		return domainErrors.NewValidationError("ACCOUNT_PASSWORD_INVALID", "password debe tener al menos 8 caracteres")
	}
	if !entities.IsValidRole(r.Role) {
		return domainErrors.NewValidationError("ACCOUNT_ROLE_INVALID", "role debe ser admin, reader o partner")
	}

	return nil
}
