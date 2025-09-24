package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// UsuarioSistemaRepository define las operaciones del repositorio para UsuarioSistema
type UsuarioSistemaRepository interface {
	Create(usuario *models.UsuarioSistema) error
	FindByID(id uint) (*models.UsuarioSistema, error)
	Update(usuario *models.UsuarioSistema) error
	Delete(id uint) error
	FindAll() ([]models.UsuarioSistema, error)
	FindByEquipoID(equipoID uint) ([]models.UsuarioSistema, error)
	FindByNombreUsuario(nombreUsuario string, equipoID uint) (*models.UsuarioSistema, error)
}

// usuarioSistemaRepository implementa UsuarioSistemaRepository
type usuarioSistemaRepository struct {
	db *gorm.DB
}

// NewUsuarioSistemaRepository crea una nueva instancia de UsuarioSistemaRepository
func NewUsuarioSistemaRepository(db *gorm.DB) UsuarioSistemaRepository {
	return &usuarioSistemaRepository{db: db}
}

// Create crea un nuevo usuario del sistema en la base de datos
func (r *usuarioSistemaRepository) Create(usuario *models.UsuarioSistema) error {
	return r.db.Create(usuario).Error
}

// FindByID busca un usuario del sistema por su ID
func (r *usuarioSistemaRepository) FindByID(id uint) (*models.UsuarioSistema, error) {
	var usuario models.UsuarioSistema
	err := r.db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// Update actualiza un usuario del sistema existente
func (r *usuarioSistemaRepository) Update(usuario *models.UsuarioSistema) error {
	return r.db.Save(usuario).Error
}

// Delete elimina un usuario del sistema por su ID
func (r *usuarioSistemaRepository) Delete(id uint) error {
	return r.db.Delete(&models.UsuarioSistema{}, id).Error
}

// FindAll retorna todos los usuarios del sistema
func (r *usuarioSistemaRepository) FindAll() ([]models.UsuarioSistema, error) {
	var usuarios []models.UsuarioSistema
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

// FindByEquipoID retorna todos los usuarios del sistema asociados a un equipo
func (r *usuarioSistemaRepository) FindByEquipoID(equipoID uint) ([]models.UsuarioSistema, error) {
	var usuarios []models.UsuarioSistema
	err := r.db.Where("equipo_id = ?", equipoID).Find(&usuarios).Error
	return usuarios, err
}

// FindByNombreUsuario busca un usuario del sistema por su nombre de usuario y equipo
func (r *usuarioSistemaRepository) FindByNombreUsuario(nombreUsuario string, equipoID uint) (*models.UsuarioSistema, error) {
	var usuario models.UsuarioSistema
	err := r.db.Where("nombre_usuario = ? AND equipo_id = ?", nombreUsuario, equipoID).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}