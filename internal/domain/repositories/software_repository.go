package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// SoftwareRepository define las operaciones del repositorio para Software
type SoftwareRepository interface {
	Create(software *models.Software) error
	FindByID(id uint) (*models.Software, error)
	Update(software *models.Software) error
	Delete(id uint) error
	FindAll() ([]models.Software, error)
	FindByEquipoID(equipoID uint) ([]models.Software, error)
}

// softwareRepository implementa SoftwareRepository
type softwareRepository struct {
	db *gorm.DB
}

// NewSoftwareRepository crea una nueva instancia de SoftwareRepository
func NewSoftwareRepository(db *gorm.DB) SoftwareRepository {
	return &softwareRepository{db: db}
}

// Create crea un nuevo software en la base de datos
func (r *softwareRepository) Create(software *models.Software) error {
	return r.db.Create(software).Error
}

// FindByID busca un software por su ID
func (r *softwareRepository) FindByID(id uint) (*models.Software, error) {
	var software models.Software
	err := r.db.First(&software, id).Error
	if err != nil {
		return nil, err
	}
	return &software, err
}

// Update actualiza un software existente
func (r *softwareRepository) Update(software *models.Software) error {
	return r.db.Save(&software).Error
}

// Delete elimina un software por su ID
func (r *softwareRepository) Delete(id uint) error {
	return r.db.Delete(&models.Software{}, id).Error
}

// FindAll retorna todos los software
func (r *softwareRepository) FindAll() ([]models.Software, error) {
	var allSoftware []models.Software
	err := r.db.Find(&allSoftware).Error
	return allSoftware, err
}

// FindByEquipoID retorna todos los software asociados a un equipo
func (r *softwareRepository) FindByEquipoID(equipoID uint) ([]models.Software, error) {
	var allSoftware []models.Software
	err := r.db.Where("equipo_id = ?", equipoID).Find(&allSoftware).Error
	return allSoftware, err
}
