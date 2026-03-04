package middlewares

import (
	"strings"

	domainErrors "rea/porticos/pkg/errors"
	httpMapper "rea/porticos/pkg/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedRoles))
	for _, r := range allowedRoles {
		normalized := strings.ToLower(strings.TrimSpace(r))
		if normalized != "" {
			allowed[normalized] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		rawRole, ok := c.Get(ContextUserRoleKey)
		if !ok {
			respondForbidden(c, domainErrors.NewForbiddenError("ROLE_REQUIRED", "rol no disponible en el token"))
			return
		}

		role, ok := rawRole.(string)
		if !ok || strings.TrimSpace(role) == "" {
			respondForbidden(c, domainErrors.NewForbiddenError("ROLE_INVALID", "rol inválido"))
			return
		}

		role = strings.ToLower(strings.TrimSpace(role))
		if _, exists := allowed[role]; !exists {
			respondForbidden(c, domainErrors.NewForbiddenError("ROLE_FORBIDDEN", "no tienes permisos para esta operación"))
			return
		}

		c.Next()
	}
}

func respondForbidden(c *gin.Context, err error) {
	status, payload := httpMapper.MapErrorToHttp(err)
	c.AbortWithStatusJSON(status, payload)
}
