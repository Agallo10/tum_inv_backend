package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// UsuarioSistemaController maneja las solicitudes HTTP relacionadas con usuarios del sistema
type UsuarioSistemaController struct {
	usuarioService services.UsuarioSistemaService
}

// NewUsuarioSistemaController crea una nueva instancia de UsuarioSistemaController
func NewUsuarioSistemaController(usuarioService services.UsuarioSistemaService) *UsuarioSistemaController {
	return &UsuarioSistemaController{
		usuarioService: usuarioService,
	}
}

// CreateUsuarioSistema maneja la creación de un nuevo usuario del sistema
func (c *UsuarioSistemaController) CreateUsuarioSistema(ctx echo.Context) error {
	usuario := new(models.UsuarioSistema)
	if err := ctx.Bind(usuario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.usuarioService.CreateUsuarioSistema(usuario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, usuario)
}

// GetUsuarioSistema obtiene un usuario del sistema por su ID
func (c *UsuarioSistemaController) GetUsuarioSistema(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	usuario, err := c.usuarioService.GetUsuarioSistemaByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, usuario)
}

// UpdateUsuarioSistema actualiza un usuario del sistema existente
func (c *UsuarioSistemaController) UpdateUsuarioSistema(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	usuario := new(models.UsuarioSistema)
	if err := ctx.Bind(usuario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	usuario.ID = uint(id)
	if err := c.usuarioService.UpdateUsuarioSistema(usuario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuario)
}

// DeleteUsuarioSistema elimina un usuario del sistema por su ID
func (c *UsuarioSistemaController) DeleteUsuarioSistema(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.usuarioService.DeleteUsuarioSistema(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Usuario eliminado correctamente"})
}

// GetAllUsuariosSistema obtiene todos los usuarios del sistema
func (c *UsuarioSistemaController) GetAllUsuariosSistema(ctx echo.Context) error {
	usuarios, err := c.usuarioService.GetAllUsuariosSistema()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuarios)
}

// GetUsuariosSistemaByEquipo obtiene todos los usuarios del sistema asociados a un equipo
func (c *UsuarioSistemaController) GetUsuariosSistemaByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	usuarios, err := c.usuarioService.GetUsuariosSistemaByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuarios)
}

// GetUsuarioSistemaByNombreUsuario obtiene un usuario del sistema por su nombre de usuario y equipo
func (c *UsuarioSistemaController) GetUsuarioSistemaByNombreUsuario(ctx echo.Context) error {
	nombreUsuario := ctx.QueryParam("nombreUsuario")
	equipoID, err := strconv.ParseUint(ctx.QueryParam("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	usuario, err := c.usuarioService.GetUsuarioSistemaByNombreUsuario(nombreUsuario, uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, usuario)
}