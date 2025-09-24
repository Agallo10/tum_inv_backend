package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// SoftwareController manjea las solicitudes HTTP relacionadas con los software
type SoftwareController struct {
	softwareService services.SoftwareService
}

// NewSoftwareController crea una nueva instancia de SoftwareController
func NewSoftwareController(softwareService services.SoftwareService) *SoftwareController {
	return &SoftwareController{softwareService: softwareService}
}

// CreateSoftware maneja la creación de un nuevo software
func (c *SoftwareController) CreateSoftware(ctx echo.Context) error {
	software := new(models.Software)
	if err := ctx.Bind(software); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.softwareService.CreateSoftware((software)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, software)
}

// GetSoftware obtiene un software por id
func (c *SoftwareController) GetSoftware(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	software, err := c.softwareService.GetSoftwareByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Software no encontrado"})
	}
	return ctx.JSON(http.StatusOK, software)
}

// UpdateSoftware actualiza un software existente
func (c *SoftwareController) UpdateSoftware(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	software := new(models.Software)
	if err := ctx.Bind(software); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	software.ID = uint(id)
	if err := c.softwareService.UpdateSoftware(software); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, software)
}

// DeleteSoftware elimina un software por su ID
func (c *SoftwareController) DeleteSoftware(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.softwareService.DeleteSoftware(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Software eliminado correctamente"})
}

// GetAllSoftware obtiene todos los software
func (c *SoftwareController) GetAllSoftware(ctx echo.Context) error {
	AllSoftware, err := c.softwareService.GetAllSoftware()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, AllSoftware)
}

// GetAllSoftwareByEquipo obtiene todos los software asociados a un equipo
func (c *SoftwareController) GetAllSoftwareByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	AllSoftware, err := c.softwareService.GetAllSoftwareByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, AllSoftware)
}
