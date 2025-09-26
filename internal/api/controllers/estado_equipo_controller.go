package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

type EstadoEquipoController struct {
	service *services.EstadoEquipoService
}

func NewEstadoEquipoController(service *services.EstadoEquipoService) *EstadoEquipoController {
	return &EstadoEquipoController{service: service}
}

// GetAllEstados obtiene todos los estados de equipo
// @Summary Obtener todos los estados de equipo
// @Description Obtiene una lista de todos los estados de equipo
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Success 200 {array} models.EstadoEquipo
// @Failure 500 {object} map[string]string
// @Router /estado-equipos [get]
func (c *EstadoEquipoController) GetAllEstados(ctx echo.Context) error {
	estados, err := c.service.GetAllEstados()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, estados)
}

// GetActiveEstados obtiene todos los estados de equipo activos
// @Summary Obtener estados de equipo activos
// @Description Obtiene una lista de todos los estados de equipo activos
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Success 200 {array} models.EstadoEquipo
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/activos [get]
func (c *EstadoEquipoController) GetActiveEstados(ctx echo.Context) error {
	estados, err := c.service.GetActiveEstados()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, estados)
}

// GetEstadoByID obtiene un estado de equipo por su ID
// @Summary Obtener estado de equipo por ID
// @Description Obtiene un estado de equipo específico por su ID
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param id path int true "ID del estado de equipo"
// @Success 200 {object} models.EstadoEquipo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/{id} [get]
func (c *EstadoEquipoController) GetEstadoByID(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	estado, err := c.service.GetEstadoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, estado)
}

// CreateEstado crea un nuevo estado de equipo
// @Summary Crear nuevo estado de equipo
// @Description Crea un nuevo estado de equipo
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param estado body models.EstadoEquipo true "Datos del estado de equipo"
// @Success 201 {object} models.EstadoEquipo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos [post]
func (c *EstadoEquipoController) CreateEstado(ctx echo.Context) error {
	var estado models.EstadoEquipo
	if err := ctx.Bind(&estado); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	if err := c.service.CreateEstado(&estado); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, estado)
}

// UpdateEstado actualiza un estado de equipo existente
// @Summary Actualizar estado de equipo
// @Description Actualiza los datos de un estado de equipo existente
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param id path int true "ID del estado de equipo"
// @Param estado body models.EstadoEquipo true "Datos actualizados del estado de equipo"
// @Success 200 {object} models.EstadoEquipo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/{id} [put]
func (c *EstadoEquipoController) UpdateEstado(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var estado models.EstadoEquipo
	if err := ctx.Bind(&estado); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	if err := c.service.UpdateEstado(uint(id), &estado); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Obtener el estado actualizado para devolverlo
	estadoActualizado, err := c.service.GetEstadoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener el estado actualizado"})
	}

	return ctx.JSON(http.StatusOK, estadoActualizado)
}

// DeleteEstado elimina un estado de equipo
// @Summary Eliminar estado de equipo
// @Description Elimina un estado de equipo por su ID
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param id path int true "ID del estado de equipo"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/{id} [delete]
func (c *EstadoEquipoController) DeleteEstado(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.service.DeleteEstado(uint(id)); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Estado de equipo eliminado correctamente"})
}

// ToggleActivo cambia el estado activo/inactivo de un estado de equipo
// @Summary Toggle estado activo/inactivo
// @Description Cambia el estado activo/inactivo de un estado de equipo
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param id path int true "ID del estado de equipo"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/{id}/toggle-activo [patch]
func (c *EstadoEquipoController) ToggleActivo(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.service.ToggleActivo(uint(id)); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Estado de activación cambiado correctamente"})
}

// GetEquiposByEstado obtiene todos los equipos que tienen un estado específico
// @Summary Obtener equipos por estado
// @Description Obtiene todos los equipos que tienen un estado específico
// @Tags EstadoEquipo
// @Accept json
// @Produce json
// @Param id path int true "ID del estado de equipo"
// @Success 200 {array} models.Equipo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /estado-equipos/{id}/equipos [get]
func (c *EstadoEquipoController) GetEquiposByEstado(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	equipos, err := c.service.GetEquiposByEstado(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, equipos)
}