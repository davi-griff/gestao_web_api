package routes

import (
	"gestor_api/src/internal/api/handlers"
	"gestor_api/src/internal/repositories/database"

	"github.com/gin-gonic/gin"
)

type CelulaRouter struct {
	router  *gin.Engine
	handler *handlers.CelulaHandler
	db      *database.Database
}

func NewCelulaRouter(db *database.Database, handler *handlers.CelulaHandler, router *gin.Engine) {
	celulaRouter := &CelulaRouter{
		router:  router,
		handler: handler,
		db:      db,
	}
	celulaRouter.setupRoutes()
}

func (r *CelulaRouter) setupRoutes() {
	celulas := r.router.Group("/celulas")
	{
		celulas.GET("", r.handler.GetCelulas)
		celulas.POST("", r.handler.CreateCelula)
		celulas.PUT(":id", r.handler.UpdateCelula)
		celulas.DELETE(":id", r.handler.DeleteCelula)
		celulas.GET(":id", r.handler.GetCelulaById)
		celulas.GET(":id/membros", r.handler.GetMembrosCelula)
		celulas.POST(":id/membros", r.handler.AdicionarMembroCelula)
		celulas.GET(":id/encontros", r.handler.GetEncontrosCelula)
		celulas.POST(":id/encontros", r.handler.CreateEncontro)
		celulas.PUT(":id/encontros/:encontroId", r.handler.UpdateEncontro)
	}
}
