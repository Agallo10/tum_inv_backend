package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// UsuarioResponsableController maneja las solicitudes HTTP relacionadas con usuarios responsables
type UsuarioResponsableController struct {
	usuarioService services.UsuarioResponsableService
}

// NewUsuarioResponsableController crea una nueva instancia de UsuarioResponsableController
func NewUsuarioResponsableController(usuarioService services.UsuarioResponsableService) *UsuarioResponsableController {
	return &UsuarioResponsableController{
		usuarioService: usuarioService,
	}
}

// CreateUsuarioResponsable maneja la creación de un nuevo usuario responsable
func (c *UsuarioResponsableController) CreateUsuarioResponsable(ctx echo.Context) error {
	usuario := new(models.UsuarioResponsable)
	if err := ctx.Bind(usuario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.usuarioService.CreateUsuarioResponsable(usuario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, usuario)
}

// GetUsuarioResponsable obtiene un usuario responsable por su ID
func (c *UsuarioResponsableController) GetUsuarioResponsable(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	usuario, err := c.usuarioService.GetUsuarioResponsableByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, usuario)
}

// UpdateUsuarioResponsable actualiza un usuario responsable existente
func (c *UsuarioResponsableController) UpdateUsuarioResponsable(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	usuario := new(models.UsuarioResponsable)
	if err := ctx.Bind(usuario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	usuario.ID = uint(id)
	if err := c.usuarioService.UpdateUsuarioResponsable(usuario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuario)
}

// DeleteUsuarioResponsable elimina un usuario responsable por su ID
func (c *UsuarioResponsableController) DeleteUsuarioResponsable(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.usuarioService.DeleteUsuarioResponsable(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Usuario eliminado correctamente"})
}

// GetAllUsuariosResponsables obtiene todos los usuarios responsables
func (c *UsuarioResponsableController) GetAllUsuariosResponsables(ctx echo.Context) error {
	usuarios, err := c.usuarioService.GetAllUsuariosResponsables()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuarios)
}

// GetUsuarioResponsableByCedula obtiene un usuario responsable por su cédula
func (c *UsuarioResponsableController) GetUsuarioResponsableByCedula(ctx echo.Context) error {
	cedula := ctx.QueryParam("cedula")
	if cedula == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "La cédula es obligatoria"})
	}

	usuario, err := c.usuarioService.GetUsuarioResponsableByCedula(cedula)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, usuario)
}

// // GetEquiposByUsuarioResponsable obtiene todos los equipos asociados a un usuario responsable
// func (c *UsuarioResponsableController) GetEquiposByUsuarioResponsable(ctx echo.Context) error {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
// 	}

// 	usuario, err := c.usuarioService.GetUsuarioResponsableByID(uint(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
// 	}

// 	return ctx.JSON(http.StatusOK, usuario.Equipo)
// }

// GetUsuariosByDependencia obtiene todos los usuarios responsables de una dependencia
func (c *UsuarioResponsableController) GetUsuariosByDependencia(ctx echo.Context) error {
	dependenciaID, err := strconv.ParseUint(ctx.Param("dependenciaId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de dependencia inválido"})
	}

	usuarios, err := c.usuarioService.GetUsuariosByDependenciaID(uint(dependenciaID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, usuarios)
}
