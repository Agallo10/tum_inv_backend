package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// DependenciaController maneja las solicitudes HTTP relacionadas con Dependencia
type DependenciaController struct {
	service services.DependenciaService
}

// NewDependenciaController crea una nueva instancia de DependenciaController
func NewDependenciaController(service services.DependenciaService) *DependenciaController {
	return &DependenciaController{service: service}
}

// CreateDependencia maneja la creación de una nueva dependencia
func (c *DependenciaController) CreateDependencia(ctx echo.Context) error {
	dependencia := new(models.Dependencia)
	if err := ctx.Bind(dependencia); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.service.CreateDependencia(dependencia); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, dependencia)
}

// GetDependencia maneja la obtención de una dependencia por su ID
func (c *DependenciaController) GetDependencia(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	dependencia, err := c.service.GetDependenciaByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Dependencia no encontrada"})
	}

	return ctx.JSON(http.StatusOK, dependencia)
}

// GetAllDependencias maneja la obtención de todas las dependencias
func (c *DependenciaController) GetAllDependencias(ctx echo.Context) error {
	dependencias, err := c.service.GetAllDependencias()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener dependencias"})
	}

	return ctx.JSON(http.StatusOK, dependencias)
}

// UpdateDependencia maneja la actualización de una dependencia existente
func (c *DependenciaController) UpdateDependencia(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	dependencia := new(models.Dependencia)
	if err := ctx.Bind(dependencia); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	dependencia.ID = uint(id)
	if err := c.service.UpdateDependencia(dependencia); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dependencia)
}

// DeleteDependencia maneja la eliminación de una dependencia
func (c *DependenciaController) DeleteDependencia(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.service.DeleteDependencia(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Dependencia eliminada correctamente"})
}

// GetDependenciasBySecretaria maneja la obtención de dependencias por secretaría
func (c *DependenciaController) GetDependenciasBySecretaria(ctx echo.Context) error {
	secretariaID, err := strconv.ParseUint(ctx.Param("secretariaId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de secretaría inválido"})
	}

	dependencias, err := c.service.GetDependenciasBySecretariaID(uint(secretariaID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dependencias)
}

// GetUsuariosByDependencia maneja la obtención de usuarios responsables por dependencia
func (c *DependenciaController) GetUsuariosByDependencia(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	dependencia, err := c.service.GetDependenciaByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Dependencia no encontrada"})
	}

	return ctx.JSON(http.StatusOK, dependencia.UsuarioResponsables)
}

// // GetEquiposByDependencia obtiene todos los equipos de una dependencia
// func (c *DependenciaController) GetEquiposByDependencia(ctx echo.Context) error {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
// 	}

// 	dependencia, err := c.service.GetDependenciaByID(uint(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Dependencia no encontrada"})
// 	}

// 	return ctx.JSON(http.StatusOK, dependencia.Equipos)
// }
