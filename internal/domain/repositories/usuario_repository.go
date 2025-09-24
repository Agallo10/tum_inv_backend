package repositories

import (
	"time"
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// UsuarioRepository define las operaciones del repositorio para Usuario
type UsuarioRepository interface {
	Create(usuario *models.Usuario) error
	FindByID(id uint) (*models.Usuario, error)
	FindByUsername(username string) (*models.Usuario, error)
	FindByEmail(email string) (*models.Usuario, error)
	Update(usuario *models.Usuario) error
	Delete(id uint) error
	FindAll() ([]models.Usuario, error)
	UpdateLastLogin(id uint) error
}

// usuarioRepository implementa UsuarioRepository
type usuarioRepository struct {
	db *gorm.DB
}

// NewUsuarioRepository crea una nueva instancia de UsuarioRepository
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{db: db}
}

// Create crea un nuevo usuario en la base de datos
func (r *usuarioRepository) Create(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}

// FindByID busca un usuario por su ID
func (r *usuarioRepository) FindByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// FindByUsername busca un usuario por su nombre de usuario
func (r *usuarioRepository) FindByUsername(username string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("username = ?", username).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// FindByEmail busca un usuario por su correo electrónico
func (r *usuarioRepository) FindByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// Update actualiza un usuario existente
func (r *usuarioRepository) Update(usuario *models.Usuario) error {
	return r.db.Save(usuario).Error
}

// Delete elimina un usuario por su ID
func (r *usuarioRepository) Delete(id uint) error {
	return r.db.Delete(&models.Usuario{}, id).Error
}

// FindAll obtiene todos los usuarios
func (r *usuarioRepository) FindAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

// UpdateLastLogin actualiza la fecha del último inicio de sesión
func (r *usuarioRepository) UpdateLastLogin(id uint) error {
	now := time.Now()
	return r.db.Model(&models.Usuario{}).Where("id = ?", id).Update("ultimo_login", &now).Error
}