package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// PerifericoController maneja las solicitudes HTTP relacionadas con periféricos
type PerifericoController struct {
	perifericoService services.PerifericoService
}

// NewPerifericoController crea una nueva instancia de PerifericoController
func NewPerifericoController(perifericoService services.PerifericoService) *PerifericoController {
	return &PerifericoController{
		perifericoService: perifericoService,
	}
}

// CreatePeriferico maneja la creación de un nuevo periférico
func (c *PerifericoController) CreatePeriferico(ctx echo.Context) error {
	periferico := new(models.Periferico)
	if err := ctx.Bind(periferico); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.perifericoService.CreatePeriferico(periferico); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, periferico)
}

// GetPeriferico obtiene un periférico por su ID
func (c *PerifericoController) GetPeriferico(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	periferico, err := c.perifericoService.GetPerifericoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Periférico no encontrado"})
	}

	return ctx.JSON(http.StatusOK, periferico)
}

// UpdatePeriferico actualiza un periférico existente
func (c *PerifericoController) UpdatePeriferico(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	periferico := new(models.Periferico)
	if err := ctx.Bind(periferico); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	periferico.ID = uint(id)
	if err := c.perifericoService.UpdatePeriferico(periferico); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, periferico)
}

// DeletePeriferico elimina un periférico por su ID
func (c *PerifericoController) DeletePeriferico(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.perifericoService.DeletePeriferico(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Periférico eliminado correctamente"})
}

// GetAllPerifericos obtiene todos los periféricos
func (c *PerifericoController) GetAllPerifericos(ctx echo.Context) error {
	perifericos, err := c.perifericoService.GetAllPerifericos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, perifericos)
}

// GetPerifericosByEquipo obtiene todos los periféricos asociados a un equipo
func (c *PerifericoController) GetPerifericosByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)

	fmt.Println(equipoID)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	perifericos, err := c.perifericoService.GetPerifericosByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, perifericos)
}
