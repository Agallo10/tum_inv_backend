package services

import (
	"errors"
	"time"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
	"tum_inv_backend/internal/infrastructure/config"

	"github.com/golang-jwt/jwt/v5"
)

// Duración de los tokens
const (
	AccessTokenDuration  = 24 * time.Hour
	RefreshTokenDuration = 7 * 24 * time.Hour
)

// JWTClaims representa los claims del token JWT
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Rol      string `json:"rol"`
	jwt.RegisteredClaims
}

// AuthService define las operaciones del servicio de autenticación
type AuthService interface {
	Register(req models.RegisterRequest) (*models.Usuario, error)
	Login(req models.LoginRequest) (*models.TokenResponse, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	RefreshToken(refreshToken string) (*models.TokenResponse, error)
	GetUserByID(id uint) (*models.Usuario, error)
}

// authService implementa AuthService
type authService struct {
	usuarioRepo repositories.UsuarioRepository
	jwtSecret   string
}

// NewAuthService crea una nueva instancia de AuthService
func NewAuthService(usuarioRepo repositories.UsuarioRepository, cfg *config.Config) AuthService {
	return &authService{
		usuarioRepo: usuarioRepo,
		jwtSecret:   cfg.JWTSecret,
	}
}

// Register registra un nuevo usuario
func (s *authService) Register(req models.RegisterRequest) (*models.Usuario, error) {
	// Verificar si el nombre de usuario ya existe
	existingUser, err := s.usuarioRepo.FindByUsername(req.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	// Verificar si el correo electrónico ya existe
	existingEmail, err := s.usuarioRepo.FindByEmail(req.Email)
	if err == nil && existingEmail != nil {
		return nil, errors.New("el correo electrónico ya está en uso")
	}

	// Crear nuevo usuario
	usuario := &models.Usuario{
		Nombre:   req.Nombre,
		Apellido: req.Apellido,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Rol:      req.Rol,
		Activo:   true,
	}

	// Encriptar contraseña
	if err := usuario.HashPassword(); err != nil {
		return nil, err
	}

	// Guardar usuario en la base de datos
	if err := s.usuarioRepo.Create(usuario); err != nil {
		return nil, err
	}

	return usuario, nil
}

// Login autentica a un usuario y genera tokens JWT
func (s *authService) Login(req models.LoginRequest) (*models.TokenResponse, error) {
	// Buscar usuario por nombre de usuario
	usuario, err := s.usuarioRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Verificar si el usuario está activo
	if !usuario.Activo {
		return nil, errors.New("cuenta desactivada")
	}

	// Verificar contraseña
	if !usuario.CheckPassword(req.Password) {
		return nil, errors.New("credenciales inválidas")
	}

	// Actualizar último login
	s.usuarioRepo.UpdateLastLogin(usuario.ID)

	// Generar tokens
	accessToken, expiresAt, err := s.generateAccessToken(usuario)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(usuario)
	if err != nil {
		return nil, err
	}

	return &models.TokenResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		Usuario: models.Usuario{
			Model:    usuario.Model,
			Nombre:   usuario.Nombre,
			Apellido: usuario.Apellido,
			Email:    usuario.Email,
			Username: usuario.Username,
			Rol:      usuario.Rol,
			Activo:   usuario.Activo,
		},
	}, nil
}

// ValidateToken valida un token JWT y devuelve sus claims
func (s *authService) ValidateToken(tokenString string) (*JWTClaims, error) {
	// Parsear token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Verificar si el token es válido
	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	// Obtener claims
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("no se pudieron obtener los claims del token")
	}

	return claims, nil
}

// RefreshToken genera un nuevo token de acceso a partir de un token de actualización
func (s *authService) RefreshToken(refreshToken string) (*models.TokenResponse, error) {
	// Validar refresh token
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Obtener usuario
	usuario, err := s.usuarioRepo.FindByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	// Verificar si el usuario está activo
	if !usuario.Activo {
		return nil, errors.New("cuenta desactivada")
	}

	// Generar nuevo token de acceso
	accessToken, expiresAt, err := s.generateAccessToken(usuario)
	if err != nil {
		return nil, err
	}

	return &models.TokenResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		Usuario: models.Usuario{
			Model:    usuario.Model,
			Nombre:   usuario.Nombre,
			Apellido: usuario.Apellido,
			Email:    usuario.Email,
			Username: usuario.Username,
			Rol:      usuario.Rol,
			Activo:   usuario.Activo,
		},
	}, nil
}

// GetUserByID obtiene un usuario por su ID
func (s *authService) GetUserByID(id uint) (*models.Usuario, error) {
	return s.usuarioRepo.FindByID(id)
}

// generateAccessToken genera un token de acceso JWT
func (s *authService) generateAccessToken(usuario *models.Usuario) (string, time.Time, error) {
	expiresAt := time.Now().Add(AccessTokenDuration)

	claims := JWTClaims{
		UserID:   usuario.ID,
		Username: usuario.Username,
		Email:    usuario.Email,
		Rol:      usuario.Rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   usuario.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

// generateRefreshToken genera un token de actualización JWT
func (s *authService) generateRefreshToken(usuario *models.Usuario) (string, error) {
	expiresAt := time.Now().Add(RefreshTokenDuration)

	claims := JWTClaims{
		UserID:   usuario.ID,
		Username: usuario.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   usuario.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
