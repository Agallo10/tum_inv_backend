package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// AccesoRemotoController maneja las solicitudes HTTP relacionadas con accesos remotos
type AccesoRemotoController struct {
	accesoService services.AccesoRemotoService
}

// NewAccesoRemotoController crea una nueva instancia de AccesoRemotoController
func NewAccesoRemotoController(accesoService services.AccesoRemotoService) *AccesoRemotoController {
	return &AccesoRemotoController{
		accesoService: accesoService,
	}
}

// CreateAccesoRemoto maneja la creación de un nuevo acceso remoto
func (c *AccesoRemotoController) CreateAccesoRemoto(ctx echo.Context) error {
	acceso := new(models.AccesoRemoto)
	if err := ctx.Bind(acceso); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.accesoService.CreateAccesoRemoto(acceso); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, acceso)
}

// GetAccesoRemoto obtiene un acceso remoto por su ID
func (c *AccesoRemotoController) GetAccesoRemoto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	acceso, err := c.accesoService.GetAccesoRemotoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Acceso remoto no encontrado"})
	}

	return ctx.JSON(http.StatusOK, acceso)
}

// UpdateAccesoRemoto actualiza un acceso remoto existente
func (c *AccesoRemotoController) UpdateAccesoRemoto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	acceso := new(models.AccesoRemoto)
	if err := ctx.Bind(acceso); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	acceso.ID = uint(id)
	if err := c.accesoService.UpdateAccesoRemoto(acceso); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, acceso)
}

// DeleteAccesoRemoto elimina un acceso remoto por su ID
func (c *AccesoRemotoController) DeleteAccesoRemoto(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.accesoService.DeleteAccesoRemoto(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Acceso remoto eliminado correctamente"})
}

// GetAllAccesosRemotos obtiene todos los accesos remotos
func (c *AccesoRemotoController) GetAllAccesosRemotos(ctx echo.Context) error {
	accesos, err := c.accesoService.GetAllAccesosRemotos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, accesos)
}

// GetAccesosRemotosByEquipo obtiene todos los accesos remotos asociados a un equipo
func (c *AccesoRemotoController) GetAccesosRemotosByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	accesos, err := c.accesoService.GetAccesosRemotosByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, accesos)
}