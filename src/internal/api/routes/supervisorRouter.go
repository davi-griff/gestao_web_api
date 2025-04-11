package routes

import (
	"gestor_api/src/internal/api/handlers"
	"gestor_api/src/internal/repositories/database"

	"github.com/gin-gonic/gin"
)

type SupervisorRouter struct {
	db      *database.Database
	handler *handlers.SupervisorHandler
	router  *gin.Engine
}

func NewSupervisorRouter(db *database.Database, handler *handlers.SupervisorHandler, router *gin.Engine) {
	supervisorRouter := &SupervisorRouter{
		db:      db,
		handler: handler,
		router:  router,
	}
	supervisorRouter.SetupRoutes()
}

func (r *SupervisorRouter) SetupRoutes() {
	supervisorRouter := r.router.Group("/supervisores")
	{
		supervisorRouter.GET("", r.handler.GetSupervisores)
		supervisorRouter.POST("", r.handler.CreateSupervisor)
		supervisorRouter.PUT("/:id", r.handler.UpdateSupervisor)
		supervisorRouter.DELETE("/:id", r.handler.DeleteSupervisor)
		supervisorRouter.GET("/:id", r.handler.GetSupervisorById)
	}
}
