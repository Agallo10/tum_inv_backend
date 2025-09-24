package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// RepuestoController maneja las solicitudes HTTP relacionadas con repuestos
type RepuestoController struct {
	repuestoService services.RepuestoService
}

// NewRepuestoController crea una nueva instancia de RepuestoController
func NewRepuestoController(repuestoService services.RepuestoService) *RepuestoController {
	return &RepuestoController{
		repuestoService: repuestoService,
	}
}

// CreateRepuesto maneja la creación de un nuevo repuesto
func (c *RepuestoController) CreateRepuesto(ctx echo.Context) error {
	repuesto := new(models.Repuesto)
	if err := ctx.Bind(repuesto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.repuestoService.CreateRepuesto(repuesto); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, repuesto)
}

// GetRepuesto obtiene un repuesto por su ID
func (c *RepuestoController) GetRepuesto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	repuesto, err := c.repuestoService.GetRepuestoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Repuesto no encontrado"})
	}

	return ctx.JSON(http.StatusOK, repuesto)
}

// UpdateRepuesto actualiza un repuesto existente
func (c *RepuestoController) UpdateRepuesto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	repuesto := new(models.Repuesto)
	if err := ctx.Bind(repuesto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	repuesto.ID = uint(id)
	if err := c.repuestoService.UpdateRepuesto(repuesto); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, repuesto)
}

// DeleteRepuesto elimina un repuesto por su ID
func (c *RepuestoController) DeleteRepuesto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.repuestoService.DeleteRepuesto(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Repuesto eliminado correctamente"})
}

// GetAllRepuestos obtiene todos los repuestos
func (c *RepuestoController) GetAllRepuestos(ctx echo.Context) error {
	repuestos, err := c.repuestoService.GetAllRepuestos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, repuestos)
}

// GetRepuestosByReporte obtiene todos los repuestos asociados a un reporte
func (c *RepuestoController) GetRepuestosByReporte(ctx echo.Context) error {
	reporteID, err := strconv.ParseUint(ctx.Param("reporteId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de reporte inválido"})
	}

	repuestos, err := c.repuestoService.GetRepuestosByReporteID(uint(reporteID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, repuestos)
}