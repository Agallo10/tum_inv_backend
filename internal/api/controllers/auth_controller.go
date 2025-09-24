package controllers

import (
	"net/http"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// AuthController maneja las solicitudes HTTP relacionadas con la autenticación
type AuthController struct {
	authService services.AuthService
}

// NewAuthController crea una nueva instancia de AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register maneja el registro de un nuevo usuario
func (c *AuthController) Register(ctx echo.Context) error {
	req := new(models.RegisterRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Validar datos
	if req.Nombre == "" || req.Apellido == "" || req.Email == "" || req.Username == "" || req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Todos los campos son obligatorios"})
	}

	// Registrar usuario
	usuario, err := c.authService.Register(*req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Ocultar contraseña en la respuesta
	usuario.Password = ""

	return ctx.JSON(http.StatusCreated, usuario)
}

// Login maneja la autenticación de un usuario
func (c *AuthController) Login(ctx echo.Context) error {
	req := new(models.LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Validar datos
	if req.Username == "" || req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Nombre de usuario y contraseña son obligatorios"})
	}

	// Autenticar usuario
	response, err := c.authService.Login(*req)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	// Ocultar contraseña en la respuesta
	response.Usuario.Password = ""

	return ctx.JSON(http.StatusOK, response)
}

// RefreshToken maneja la renovación de un token JWT
func (c *AuthController) RefreshToken(ctx echo.Context) error {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	req := new(RefreshRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Validar datos
	if req.RefreshToken == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Token de actualización es obligatorio"})
	}

	// Renovar token
	response, err := c.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	// Ocultar contraseña en la respuesta
	response.Usuario.Password = ""

	return ctx.JSON(http.StatusOK, response)
}

// GetProfile obtiene el perfil del usuario autenticado
func (c *AuthController) GetProfile(ctx echo.Context) error {
	// Obtener ID de usuario desde el contexto (establecido por el middleware de autenticación)
	userID, ok := ctx.Get("user_id").(uint)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo obtener el ID de usuario"})
	}

	// Obtener usuario
	usuario, err := c.authService.GetUserByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	// Ocultar contraseña en la respuesta
	usuario.Password = ""

	return ctx.JSON(http.StatusOK, usuario)
}
