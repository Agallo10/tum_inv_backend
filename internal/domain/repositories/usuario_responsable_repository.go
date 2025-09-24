package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// UsuarioResponsableRepository define las operaciones del repositorio para UsuarioResponsable
type UsuarioResponsableRepository interface {
	Create(usuario *models.UsuarioResponsable) error
	FindByID(id uint) (*models.UsuarioResponsable, error)
	Update(usuario *models.UsuarioResponsable) error
	Delete(id uint) error
	FindAll() ([]models.UsuarioResponsable, error)
	FindByCedula(cedula string) (*models.UsuarioResponsable, error)
	FindByDependenciaID(dependenciaID uint) ([]models.UsuarioResponsable, error)
}

// usuarioResponsableRepository implementa UsuarioResponsableRepository
type usuarioResponsableRepository struct {
	db *gorm.DB
}

// NewUsuarioResponsableRepository crea una nueva instancia de UsuarioResponsableRepository
func NewUsuarioResponsableRepository(db *gorm.DB) UsuarioResponsableRepository {
	return &usuarioResponsableRepository{db: db}
}

// Create crea un nuevo usuario responsable en la base de datos
func (r *usuarioResponsableRepository) Create(usuario *models.UsuarioResponsable) error {
	return r.db.Create(usuario).Error
}

// FindByID busca un usuario responsable por su ID
func (r *usuarioResponsableRepository) FindByID(id uint) (*models.UsuarioResponsable, error) {
	var usuario models.UsuarioResponsable
	err := r.db.Preload("Equipos").First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// Update actualiza un usuario responsable existente
func (r *usuarioResponsableRepository) Update(usuario *models.UsuarioResponsable) error {
	return r.db.Save(usuario).Error
}

// Delete elimina un usuario responsable por su ID
func (r *usuarioResponsableRepository) Delete(id uint) error {
	return r.db.Delete(&models.UsuarioResponsable{}, id).Error
}

// FindAll retorna todos los usuarios responsables
func (r *usuarioResponsableRepository) FindAll() ([]models.UsuarioResponsable, error) {
	var usuarios []models.UsuarioResponsable
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

// FindByCedula busca un usuario responsable por su número de cédula
func (r *usuarioResponsableRepository) FindByCedula(cedula string) (*models.UsuarioResponsable, error) {
	var usuario models.UsuarioResponsable
	err := r.db.Where("cedula = ?", cedula).Preload("Equipos").First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// FindByDependenciaID busca usuarios responsables por su dependencia
func (r *usuarioResponsableRepository) FindByDependenciaID(dependenciaID uint) ([]models.UsuarioResponsable, error) {
	var usuarios []models.UsuarioResponsable
	err := r.db.Where("dependencia_id = ?", dependenciaID).Find(&usuarios).Error
	return usuarios, err
}
