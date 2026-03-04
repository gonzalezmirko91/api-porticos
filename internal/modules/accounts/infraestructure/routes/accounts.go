package routes

import (
	"rea/porticos/internal/modules/accounts/infraestructure/handler"
	"rea/porticos/pkg/middlewares"
	"rea/porticos/pkg/version"

	"github.com/gin-gonic/gin"
)

func ConfigAccountsVersion(accountsHandler *handler.AccountsHandler) {
	wrapperFunc := func(rg *gin.RouterGroup, ctrl any) {
		accountsCtrl := ctrl.(*handler.AccountsHandler)
		RegisterAccountsRoutes(rg, accountsCtrl)
	}

	version.ConfigControllerVersion("accounts", accountsHandler, wrapperFunc)
}

func RegisterAccountsRoutes(rg *gin.RouterGroup, h *handler.AccountsHandler) {
	adminOnly := middlewares.RequireRoles("admin")
	rg.POST("/signup", h.Signup)
	rg.POST("", h.CreateFirstAdmin)
	rg.POST("/managed", adminOnly, h.CreateManaged)
}
