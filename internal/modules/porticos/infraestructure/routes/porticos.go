package routes

import (
	"rea/porticos/internal/modules/porticos/infraestructure/handler"
	"rea/porticos/pkg/middlewares"
	"rea/porticos/pkg/version"

	"github.com/gin-gonic/gin"
)

func ConfigPorticosVersion(porticosHandler *handler.PorticosHandler) {
	wrapperFunc := func(rg *gin.RouterGroup, ctrl any) {
		porticosCtrl := ctrl.(*handler.PorticosHandler)
		RegisterPorticosRoutes(rg, porticosCtrl)
	}

	version.ConfigControllerVersion("porticos", porticosHandler, wrapperFunc)
}

func RegisterPorticosRoutes(rg *gin.RouterGroup, h *handler.PorticosHandler) {
	readRoles := middlewares.RequireRoles("reader", "partner", "admin")
	adminOnly := middlewares.RequireRoles("admin")

	rg.GET("", readRoles, h.List)
	rg.GET("/:id", readRoles, h.GetByID)
	rg.POST("", adminOnly, h.Create)
	rg.PUT("/:id", adminOnly, h.Update)
	rg.DELETE("/:id", adminOnly, h.Delete)
}
