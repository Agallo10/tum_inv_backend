package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

type EstadoEquipoRepository struct {
	db *gorm.DB
}

func NewEstadoEquipoRepository(db *gorm.DB) *EstadoEquipoRepository {
	return &EstadoEquipoRepository{db: db}
}

// GetAll obtiene todos los estados de equipo
func (r *EstadoEquipoRepository) GetAll() ([]models.EstadoEquipo, error) {
	var estados []models.EstadoEquipo
	err := r.db.Find(&estados).Error
	return estados, err
}

// GetByID obtiene un estado de equipo por su ID
func (r *EstadoEquipoRepository) GetByID(id uint) (*models.EstadoEquipo, error) {
	var estado models.EstadoEquipo
	err := r.db.First(&estado, id).Error
	if err != nil {
		return nil, err
	}
	return &estado, nil
}

// GetActive obtiene todos los estados de equipo activos
func (r *EstadoEquipoRepository) GetActive() ([]models.EstadoEquipo, error) {
	var estados []models.EstadoEquipo
	err := r.db.Where("activo = ?", true).Find(&estados).Error
	return estados, err
}

// Create crea un nuevo estado de equipo
func (r *EstadoEquipoRepository) Create(estado *models.EstadoEquipo) error {
	return r.db.Create(estado).Error
}

// Update actualiza un estado de equipo existente
func (r *EstadoEquipoRepository) Update(estado *models.EstadoEquipo) error {
	return r.db.Save(estado).Error
}

// Delete elimina (soft delete) un estado de equipo
func (r *EstadoEquipoRepository) Delete(id uint) error {
	return r.db.Delete(&models.EstadoEquipo{}, id).Error
}

// ExistsByName verifica si ya existe un estado con el mismo nombre
func (r *EstadoEquipoRepository) ExistsByName(nombre string) (bool, error) {
	var count int64
	err := r.db.Model(&models.EstadoEquipo{}).Where("nombre = ?", nombre).Count(&count).Error
	return count > 0, err
}

// ExistsByNameExcludingID verifica si ya existe un estado con el mismo nombre excluyendo un ID específico
func (r *EstadoEquipoRepository) ExistsByNameExcludingID(nombre string, id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.EstadoEquipo{}).Where("nombre = ? AND id != ?", nombre, id).Count(&count).Error
	return count > 0, err
}

// GetEquiposByEstado obtiene todos los equipos que tienen un estado específico
func (r *EstadoEquipoRepository) GetEquiposByEstado(estadoID uint) ([]models.Equipo, error) {
	var equipos []models.Equipo
	err := r.db.Where("estado_equipo_id = ?", estadoID).Find(&equipos).Error
	return equipos, err
}