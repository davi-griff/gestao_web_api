package routes

import (
	"gestor_api/src/internal/api/handlers"
	"gestor_api/src/internal/repositories/database"

	"github.com/gin-gonic/gin"
)

type RedeRouter struct {
	router  *gin.Engine
	handler *handlers.RedeHandler
	db      *database.Database
}

func NewRedeRouter(db *database.Database, handler *handlers.RedeHandler, router *gin.Engine) {
	redeRouter := &RedeRouter{
		router:  router,
		handler: handler,
		db:      db,
	}
	redeRouter.setupRoutes()
}

func (r *RedeRouter) setupRoutes() {
	redes := r.router.Group("/redes")
	{
		redes.GET("", r.handler.GetRedes)
		redes.POST("", r.handler.CreateRede)
		redes.PUT(":id", r.handler.UpdateRede)
		redes.DELETE(":id", r.handler.DeleteRede)
	}
}
