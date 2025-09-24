package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// EquipoController maneja las solicitudes HTTP relacionadas con equipos
type EquipoController struct {
	equipoService services.EquipoService
}

// NewEquipoController crea una nueva instancia de EquipoController
func NewEquipoController(equipoService services.EquipoService) *EquipoController {
	return &EquipoController{
		equipoService: equipoService,
	}
}

// CreateEquipo maneja la creación de un nuevo equipo
func (c *EquipoController) CreateEquipo(ctx echo.Context) error {
	equipo := new(models.Equipo)
	if err := ctx.Bind(equipo); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.equipoService.CreateEquipo(equipo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, equipo)
}

// GetEquipo obtiene un equipo por su ID
func (c *EquipoController) GetEquipo(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	equipo, err := c.equipoService.GetEquipoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Equipo no encontrado"})
	}

	return ctx.JSON(http.StatusOK, equipo)
}

// UpdateEquipo actualiza un equipo existente
func (c *EquipoController) UpdateEquipo(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	equipo := new(models.Equipo)
	if err := ctx.Bind(equipo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	equipo.ID = uint(id)
	if err := c.equipoService.UpdateEquipo(equipo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, equipo)
}

// DeleteEquipo elimina un equipo por su ID
func (c *EquipoController) DeleteEquipo(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.equipoService.DeleteEquipo(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Equipo eliminado correctamente"})
}

// GetAllEquipos obtiene todos los equipos
func (c *EquipoController) GetAllEquipos(ctx echo.Context) error {
	equipos, err := c.equipoService.GetAllEquipos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, equipos)
}

// GetEquiposByDependencia obtiene todos los equipos de una dependencia
func (c *EquipoController) GetEquiposByDependencia(ctx echo.Context) error {
	dependenciaID, err := strconv.ParseUint(ctx.Param("dependenciaId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de dependencia inválido"})
	}

	equipos, err := c.equipoService.GetEquiposByDependenciaID(uint(dependenciaID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, equipos)
}

// GetEquiposByDependencia obtiene todos los equipos de una dependencia
func (c *EquipoController) GetEquipoUsuDepByID(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	equipo, err := c.equipoService.GetEquipoUsuDepByID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, equipo)
}
