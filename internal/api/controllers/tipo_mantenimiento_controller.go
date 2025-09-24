package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// TipoMantenimientoController maneja las solicitudes HTTP relacionadas con tipos de mantenimiento
type TipoMantenimientoController struct {
	tipoService services.TipoMantenimientoService
}

// NewTipoMantenimientoController crea una nueva instancia de TipoMantenimientoController
func NewTipoMantenimientoController(tipoService services.TipoMantenimientoService) *TipoMantenimientoController {
	return &TipoMantenimientoController{
		tipoService: tipoService,
	}
}

// CreateTipoMantenimiento maneja la creación de un nuevo tipo de mantenimiento
func (c *TipoMantenimientoController) CreateTipoMantenimiento(ctx echo.Context) error {
	tipo := new(models.TipoMantenimiento)
	if err := ctx.Bind(tipo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.tipoService.CreateTipoMantenimiento(tipo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, tipo)
}

// GetTipoMantenimiento obtiene un tipo de mantenimiento por su ID
func (c *TipoMantenimientoController) GetTipoMantenimiento(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	tipo, err := c.tipoService.GetTipoMantenimientoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Tipo de mantenimiento no encontrado"})
	}

	return ctx.JSON(http.StatusOK, tipo)
}

// UpdateTipoMantenimiento actualiza un tipo de mantenimiento existente
func (c *TipoMantenimientoController) UpdateTipoMantenimiento(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	tipo := new(models.TipoMantenimiento)
	if err := ctx.Bind(tipo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	tipo.ID = uint(id)
	if err := c.tipoService.UpdateTipoMantenimiento(tipo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, tipo)
}

// DeleteTipoMantenimiento elimina un tipo de mantenimiento por su ID
func (c *TipoMantenimientoController) DeleteTipoMantenimiento(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.tipoService.DeleteTipoMantenimiento(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Tipo de mantenimiento eliminado correctamente"})
}

// GetAllTiposMantenimiento obtiene todos los tipos de mantenimiento
func (c *TipoMantenimientoController) GetAllTiposMantenimiento(ctx echo.Context) error {
	tipos, err := c.tipoService.GetAllTiposMantenimiento()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, tipos)
}

// GetTiposMantenimientoByReporte obtiene todos los tipos de mantenimiento asociados a un reporte
func (c *TipoMantenimientoController) GetTiposMantenimientoByReporte(ctx echo.Context) error {
	reporteID, err := strconv.ParseUint(ctx.Param("reporteId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de reporte inválido"})
	}

	tipos, err := c.tipoService.GetTiposMantenimientoByReporteID(uint(reporteID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, tipos)
}