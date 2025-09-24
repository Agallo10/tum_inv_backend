package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// ConfiguracionRedRepository define las operaciones del repositorio para ConfiguracionRed
type ConfiguracionRedRepository interface {
	Create(configuracion *models.ConfiguracionRed) error
	FindByID(id uint) (*models.ConfiguracionRed, error)
	Update(configuracion *models.ConfiguracionRed) error
	Delete(id uint) error
	FindAll() ([]models.ConfiguracionRed, error)
	FindByEquipoID(equipoID uint) (*models.ConfiguracionRed, error)
}

// configuracionRedRepository implementa ConfiguracionRedRepository
type configuracionRedRepository struct {
	db *gorm.DB
}

// NewConfiguracionRedRepository crea una nueva instancia de ConfiguracionRedRepository
func NewConfiguracionRedRepository(db *gorm.DB) ConfiguracionRedRepository {
	return &configuracionRedRepository{db: db}
}

// Create crea una nueva configuración de red en la base de datos
func (r *configuracionRedRepository) Create(configuracion *models.ConfiguracionRed) error {
	return r.db.Create(configuracion).Error
}

// FindByID busca una configuración de red por su ID
func (r *configuracionRedRepository) FindByID(id uint) (*models.ConfiguracionRed, error) {
	var configuracion models.ConfiguracionRed
	err := r.db.First(&configuracion, id).Error
	if err != nil {
		return nil, err
	}
	return &configuracion, nil
}

// Update actualiza una configuración de red existente
func (r *configuracionRedRepository) Update(configuracion *models.ConfiguracionRed) error {
	return r.db.Save(configuracion).Error
}

// Delete elimina una configuración de red por su ID
func (r *configuracionRedRepository) Delete(id uint) error {
	return r.db.Delete(&models.ConfiguracionRed{}, id).Error
}

// FindAll retorna todas las configuraciones de red
func (r *configuracionRedRepository) FindAll() ([]models.ConfiguracionRed, error) {
	var configuraciones []models.ConfiguracionRed
	err := r.db.Find(&configuraciones).Error
	return configuraciones, err
}

// FindByEquipoID retorna la configuración de red asociada a un equipo
func (r *configuracionRedRepository) FindByEquipoID(equipoID uint) (*models.ConfiguracionRed, error) {
	var configuracion models.ConfiguracionRed
	err := r.db.Where("equipo_id = ?", equipoID).First(&configuracion).Error
	if err != nil {
		return nil, err
	}
	return &configuracion, nil
}
