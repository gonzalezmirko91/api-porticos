package routes

import (
	"rea/porticos/internal/modules/vehiculos/infraestructure/handler"
	"rea/porticos/pkg/middlewares"
	"rea/porticos/pkg/version"

	"github.com/gin-gonic/gin"
)

func ConfigVehiculosVersion(vehiculosHandler *handler.VehiculosHandler) {
	wrapperFunc := func(rg *gin.RouterGroup, ctrl any) {
		vehiculosCtrl := ctrl.(*handler.VehiculosHandler)
		RegisterVehiculosRoutes(rg, vehiculosCtrl)
	}

	version.ConfigControllerVersion("vehiculos", vehiculosHandler, wrapperFunc)
}

func RegisterVehiculosRoutes(rg *gin.RouterGroup, h *handler.VehiculosHandler) {
	allowed := middlewares.RequireRoles("reader", "partner", "admin")

	rg.GET("", allowed, h.List)
	rg.GET("/:id", allowed, h.GetByID)
	rg.POST("", allowed, h.Create)
	rg.PUT("/:id", allowed, h.Update)
	rg.DELETE("/:id", allowed, h.Delete)
}
