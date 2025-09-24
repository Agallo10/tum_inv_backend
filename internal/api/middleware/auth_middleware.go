package middleware

import (
	"net/http"
	"strings"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// JWTMiddleware es un middleware para validar tokens JWT
type JWTMiddleware struct {
	authService services.AuthService
}

// NewJWTMiddleware crea una nueva instancia de JWTMiddleware
func NewJWTMiddleware(authService services.AuthService) *JWTMiddleware {
	return &JWTMiddleware{
		authService: authService,
	}
}

// Authenticate valida el token JWT y establece el ID de usuario en el contexto
func (m *JWTMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtener token del encabezado Authorization
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
		}

		// Verificar formato del token (Bearer <token>)
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Formato de token inválido"})
		}

		token := parts[1]

		// Validar token
		userID, err := m.authService.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		// Establecer ID de usuario en el contexto
		c.Set("user_id", userID)

		return next(c)
	}
}

// RequireRole verifica si el usuario tiene el rol requerido
func (m *JWTMiddleware) RequireRole(role string, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Primero autenticar
		err := m.Authenticate(func(c echo.Context) error { return nil })(c)
		if err != nil {
			return err
		}

		// Obtener ID de usuario del contexto
		userID, ok := c.Get("user_id").(uint)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se pudo obtener el ID de usuario"})
		}

		// Obtener usuario
		usuario, err := m.authService.GetUserByID(userID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
		}

		// Verificar rol
		if usuario.Rol != role {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "No tiene permisos para acceder a este recurso"})
		}

		return next(c)
	}
}