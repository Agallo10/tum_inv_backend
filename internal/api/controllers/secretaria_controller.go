package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// SecretariaController maneja las solicitudes HTTP relacionadas con Secretaria
type SecretariaController struct {
	service services.SecretariaService
}

// NewSecretariaController crea un nuevo controlador de Secretaria
func NewSecretariaController(service services.SecretariaService) *SecretariaController {
	return &SecretariaController{service: service}
}

// CreateSecretaria maneja la creación de una nueva Secretaria
func (c *SecretariaController) CreateSecretaria(ctx echo.Context) error {
	secretaria := new(models.Secretaria)
	if err := ctx.Bind(secretaria); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.service.CreateSecretaria(secretaria); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, secretaria)
}

// GetSecretaria obtiene una Secretaria por su ID
func (c *SecretariaController) GetSecretaria(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	secretaria, err := c.service.GetSecretariaByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Secretaría/ no encontrada"})
	}

	return ctx.JSON(http.StatusOK, secretaria)
}

// GetAllSecretarias obtiene todas las Secretarias
func (c *SecretariaController) GetAllSecretarias(ctx echo.Context) error {
	secretarias, err := c.service.GetAllSecretarias()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener secretarías/s"})
	}

	return ctx.JSON(http.StatusOK, secretarias)
}

// UpdateSecretaria actualiza una Secretaria existente
func (c *SecretariaController) UpdateSecretaria(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	secretaria := new(models.Secretaria)
	if err := ctx.Bind(secretaria); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	secretaria.ID = uint(id)
	if err := c.service.UpdateSecretaria(secretaria); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, secretaria)
}

// DeleteSecretaria elimina una Secretaria por su ID
func (c *SecretariaController) DeleteSecretaria(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.service.DeleteSecretaria(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Secretaría/ eliminada correctamente"})
}

// GetDependenciasBySecretaria obtiene todas las dependencias de una secretaría
func (c *SecretariaController) GetDependenciasBySecretaria(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	secretaria, err := c.service.GetSecretariaByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Secretaría no encontrada"})
	}

	return ctx.JSON(http.StatusOK, secretaria.Dependencias)
}
