package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// HardwareInternoRepository define las operaciones del repositorio para HardwareInterno
type HardwareInternoRepository interface {
	Create(hardware *models.HardwareInterno) error
	FindByID(id uint) (*models.HardwareInterno, error)
	Update(hardware *models.HardwareInterno) error
	Delete(id uint) error
	FindAll() ([]models.HardwareInterno, error)
	FindByEquipoID(equipoID uint) ([]models.HardwareInterno, error)
}

// hardwareInternoRepository implementa HardwareInternoRepository
type hardwareInternoRepository struct {
	db *gorm.DB
}

// NewHardwareInternoRepository crea una nueva instancia de HardwareInternoRepository
func NewHardwareInternoRepository(db *gorm.DB) HardwareInternoRepository {
	return &hardwareInternoRepository{db: db}
}

// Create crea un nuevo hardware interno en la base de datos
func (r *hardwareInternoRepository) Create(hardware *models.HardwareInterno) error {
	return r.db.Create(hardware).Error
}

// FindByID busca un hardware interno por su ID
func (r *hardwareInternoRepository) FindByID(id uint) (*models.HardwareInterno, error) {
	var hardware models.HardwareInterno
	err := r.db.First(&hardware, id).Error
	if err != nil {
		return nil, err
	}
	return &hardware, nil
}

// Update actualiza un hardware interno existente
func (r *hardwareInternoRepository) Update(hardware *models.HardwareInterno) error {
	return r.db.Save(hardware).Error
}

// Delete elimina un hardware interno por su ID
func (r *hardwareInternoRepository) Delete(id uint) error {
	return r.db.Delete(&models.HardwareInterno{}, id).Error
}

// FindAll retorna todos los componentes de hardware interno
func (r *hardwareInternoRepository) FindAll() ([]models.HardwareInterno, error) {
	var hardwareInternos []models.HardwareInterno
	err := r.db.Find(&hardwareInternos).Error
	return hardwareInternos, err
}

// FindByEquipoID retorna todos los componentes de hardware interno asociados a un equipo
func (r *hardwareInternoRepository) FindByEquipoID(equipoID uint) ([]models.HardwareInterno, error) {
	var hardwareInternos []models.HardwareInterno
	err := r.db.Where("equipo_id = ?", equipoID).Find(&hardwareInternos).Error
	return hardwareInternos, err
}
