package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// HardwareInternoController maneja las solicitudes HTTP relacionadas con hardware interno
type HardwareInternoController struct {
	hardwareService services.HardwareInternoService
}

// NewHardwareInternoController crea una nueva instancia de HardwareInternoController
func NewHardwareInternoController(hardwareService services.HardwareInternoService) *HardwareInternoController {
	return &HardwareInternoController{
		hardwareService: hardwareService,
	}
}

// CreateHardwareInterno maneja la creación de un nuevo componente de hardware interno
func (c *HardwareInternoController) CreateHardwareInterno(ctx echo.Context) error {
	hardware := new(models.HardwareInterno)
	if err := ctx.Bind(hardware); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.hardwareService.CreateHardwareInterno(hardware); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, hardware)
}

// GetHardwareInterno obtiene un componente de hardware interno por su ID
func (c *HardwareInternoController) GetHardwareInterno(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	hardware, err := c.hardwareService.GetHardwareInternoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Componente de hardware no encontrado"})
	}

	return ctx.JSON(http.StatusOK, hardware)
}

// UpdateHardwareInterno actualiza un componente de hardware interno existente
func (c *HardwareInternoController) UpdateHardwareInterno(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	hardware := new(models.HardwareInterno)
	if err := ctx.Bind(hardware); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	hardware.ID = uint(id)
	if err := c.hardwareService.UpdateHardwareInterno(hardware); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, hardware)
}

// DeleteHardwareInterno elimina un componente de hardware interno por su ID
func (c *HardwareInternoController) DeleteHardwareInterno(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.hardwareService.DeleteHardwareInterno(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Componente de hardware eliminado correctamente"})
}

// GetAllHardwareInterno obtiene todos los componentes de hardware interno
func (c *HardwareInternoController) GetAllHardwareInterno(ctx echo.Context) error {
	hardwareInternos, err := c.hardwareService.GetAllHardwareInterno()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, hardwareInternos)
}

// GetHardwareInternoByEquipo obtiene todos los componentes de hardware interno asociados a un equipo
func (c *HardwareInternoController) GetHardwareInternoByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	hardwareInternos, err := c.hardwareService.GetHardwareInternoByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, hardwareInternos)
}