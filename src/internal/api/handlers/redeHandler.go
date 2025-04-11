package handlers

import (
	"gestor_api/src/internal/models"
	"gestor_api/src/internal/repositories/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RedeHandler struct {
	db *database.Database
}

func NewRedeHandler(db *database.Database) *RedeHandler {
	return &RedeHandler{db: db}
}

func (h *RedeHandler) GetRedes(c *gin.Context) {
	redes, err := h.db.GetRedes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, redes)
}

func (h *RedeHandler) CreateRede(c *gin.Context) {
	var rede models.Rede
	if err := c.ShouldBindJSON(&rede); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rede, err := h.db.CreateRede(rede)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rede)
}

func (h *RedeHandler) UpdateRede(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var rede models.Rede
	if err := c.ShouldBindJSON(&rede); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rede.ID = uint(idUint)
	rede, err = h.db.UpdateRede(rede)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rede)
}

func (h *RedeHandler) DeleteRede(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.db.DeleteRede(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rede deleted successfully"})
}
