package handlers

import (
	"gestor_api/src/internal/models"
	"gestor_api/src/internal/repositories/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupervisorHandler struct {
	db *database.Database
}

func NewSupervisorHandler(db *database.Database) *SupervisorHandler {
	return &SupervisorHandler{db: db}
}

func (h *SupervisorHandler) GetSupervisores(c *gin.Context) {
	supervisores, err := h.db.GetSupervisores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supervisores)
}

func (h *SupervisorHandler) CreateSupervisor(c *gin.Context) {
	var supervisor models.Supervisor
	if err := c.ShouldBindJSON(&supervisor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	supervisor, err := h.db.CreateSupervisor(supervisor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, supervisor)
}

func (h *SupervisorHandler) UpdateSupervisor(c *gin.Context) {
	var supervisor models.Supervisor
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supervisor ID"})
		return
	}
	if err := c.ShouldBindJSON(&supervisor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	supervisor.ID = uint(idUint)
	supervisor, err = h.db.UpdateSupervisor(supervisor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supervisor)
}

func (h *SupervisorHandler) DeleteSupervisor(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supervisor ID"})
		return
	}
	if err := h.db.DeleteSupervisor(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Supervisor deleted successfully"})
}

func (h *SupervisorHandler) GetSupervisorById(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supervisor ID"})
		return
	}
	supervisor, err := h.db.GetSupervisorById(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supervisor)
}
