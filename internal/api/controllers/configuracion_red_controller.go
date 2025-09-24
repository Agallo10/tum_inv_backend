package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// ConfiguracionRedController maneja las solicitudes HTTP relacionadas con configuraciones de red
type ConfiguracionRedController struct {
	configuracionService services.ConfiguracionRedService
}

// NewConfiguracionRedController crea una nueva instancia de ConfiguracionRedController
func NewConfiguracionRedController(configuracionService services.ConfiguracionRedService) *ConfiguracionRedController {
	return &ConfiguracionRedController{
		configuracionService: configuracionService,
	}
}

// CreateConfiguracionRed maneja la creación de una nueva configuración de red
func (c *ConfiguracionRedController) CreateConfiguracionRed(ctx echo.Context) error {
	configuracion := new(models.ConfiguracionRed)
	if err := ctx.Bind(configuracion); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.configuracionService.CreateConfiguracionRed(configuracion); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, configuracion)
}

// GetConfiguracionRed obtiene una configuración de red por su ID
func (c *ConfiguracionRedController) GetConfiguracionRed(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	configuracion, err := c.configuracionService.GetConfiguracionRedByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Configuración de red no encontrada"})
	}

	return ctx.JSON(http.StatusOK, configuracion)
}

// UpdateConfiguracionRed actualiza una configuración de red existente
func (c *ConfiguracionRedController) UpdateConfiguracionRed(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	configuracion := new(models.ConfiguracionRed)
	if err := ctx.Bind(configuracion); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	configuracion.ID = uint(id)
	if err := c.configuracionService.UpdateConfiguracionRed(configuracion); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, configuracion)
}

// DeleteConfiguracionRed elimina una configuración de red por su ID
func (c *ConfiguracionRedController) DeleteConfiguracionRed(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.configuracionService.DeleteConfiguracionRed(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Configuración de red eliminada correctamente"})
}

// GetAllConfiguracionesRed obtiene todas las configuraciones de red
func (c *ConfiguracionRedController) GetAllConfiguracionesRed(ctx echo.Context) error {
	configuraciones, err := c.configuracionService.GetAllConfiguracionesRed()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, configuraciones)
}

// GetConfiguracionRedByEquipo obtiene la configuración de red asociada a un equipo
func (c *ConfiguracionRedController) GetConfiguracionRedByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	configuracion, err := c.configuracionService.GetConfiguracionRedByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, configuracion)
}
