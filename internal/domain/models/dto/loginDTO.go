package dto

// LoginDTO contiene las credenciales de inicio de sesión
type LoginDTO struct {
	Email    string `json:"email" validate:"required,customEmail,max=50"`
	Password string `json:"password" validate:"required,min=5,max=12"`
}
