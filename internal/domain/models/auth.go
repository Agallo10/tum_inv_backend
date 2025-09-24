package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Usuario representa un usuario del sistema con capacidad de autenticación
type Usuario struct {
	gorm.Model
	Nombre      string `gorm:"not null"`
	Apellido    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Username    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Rol         string `gorm:"check:rol IN ('admin', 'usuario', 'tecnico');default:'usuario'"`
	Activo      bool   `gorm:"default:true"`
	UltimoLogin *time.Time
}

// HashPassword encripta la contraseña del usuario
func (u *Usuario) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifica si la contraseña proporcionada coincide con la almacenada
func (u *Usuario) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// LoginRequest representa los datos necesarios para iniciar sesión
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest representa los datos necesarios para registrar un nuevo usuario
type RegisterRequest struct {
	Nombre   string `json:"nombre" validate:"required"`
	Apellido string `json:"apellido" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Rol      string `json:"rol" validate:"omitempty,oneof=admin usuario tecnico"`
}

// TokenResponse representa la respuesta con el token JWT
type TokenResponse struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	Usuario      Usuario   `json:"usuario"`
}
