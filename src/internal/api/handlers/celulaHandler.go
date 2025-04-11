package handlers

import (
	"gestor_api/src/internal/models"
	"gestor_api/src/internal/repositories/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CelulaHandler struct {
	db *database.Database
}

func NewCelulaHandler(db *database.Database) *CelulaHandler {
	return &CelulaHandler{db: db}
}

func (h *CelulaHandler) GetCelulas(c *gin.Context) {
	celulas, err := h.db.GetCelulas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, celulas)
}

func (h *CelulaHandler) CreateCelula(c *gin.Context) {
	var celula models.Celula
	if err := c.ShouldBindJSON(&celula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	celula, err := h.db.CreateCelula(celula)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, celula)
}

func (h *CelulaHandler) UpdateCelula(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var celula models.Celula
	if err := c.ShouldBindJSON(&celula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	celula.ID = uint(idUint)
	celula, err = h.db.UpdateCelula(celula)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, celula)
}

func (h *CelulaHandler) DeleteCelula(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.db.DeleteCelula(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CelulaHandler) GetCelulaById(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	celula, err := h.db.GetCelulaById(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, celula)
}

func (h *CelulaHandler) GetMembrosCelula(c *gin.Context) {
	celulaId := c.Param("id")
	celulaIdUint, err := strconv.ParseUint(celulaId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	membrosCelula, err := h.db.GetMembrosCelula(uint(celulaIdUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, membrosCelula)
}

func (h *CelulaHandler) GetEncontrosCelula(c *gin.Context) {
	celulaId := c.Param("id")
	celulaIdUint, err := strconv.ParseUint(celulaId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	encontrosCelula, err := h.db.GetEncontroByIdCelula(uint(celulaIdUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, encontrosCelula)
}

func (h *CelulaHandler) CreateEncontro(c *gin.Context) {
	var encontro models.Encontro
	if err := c.ShouldBindJSON(&encontro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encontro, err := h.db.CreateEncontro(encontro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, encontro)
}

func (h *CelulaHandler) UpdateEncontro(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var encontro models.Encontro
	if err := c.ShouldBindJSON(&encontro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encontro.ID = uint(idUint)
	encontro, err = h.db.UpdateEncontro(encontro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, encontro)
}
