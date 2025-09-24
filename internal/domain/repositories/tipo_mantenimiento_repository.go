package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// TipoMantenimientoRepository define las operaciones del repositorio para TipoMantenimiento
type TipoMantenimientoRepository interface {
	Create(tipo *models.TipoMantenimiento) error
	FindByID(id uint) (*models.TipoMantenimiento, error)
	Update(tipo *models.TipoMantenimiento) error
	Delete(id uint) error
	FindAll() ([]models.TipoMantenimiento, error)
	FindByReporteID(reporteID uint) ([]models.TipoMantenimiento, error)
}

// tipoMantenimientoRepository implementa TipoMantenimientoRepository
type tipoMantenimientoRepository struct {
	db *gorm.DB
}

// NewTipoMantenimientoRepository crea una nueva instancia de TipoMantenimientoRepository
func NewTipoMantenimientoRepository(db *gorm.DB) TipoMantenimientoRepository {
	return &tipoMantenimientoRepository{db: db}
}

// Create crea un nuevo tipo de mantenimiento en la base de datos
func (r *tipoMantenimientoRepository) Create(tipo *models.TipoMantenimiento) error {
	return r.db.Create(tipo).Error
}

// FindByID busca un tipo de mantenimiento por su ID
func (r *tipoMantenimientoRepository) FindByID(id uint) (*models.TipoMantenimiento, error) {
	var tipo models.TipoMantenimiento
	err := r.db.First(&tipo, id).Error
	if err != nil {
		return nil, err
	}
	return &tipo, nil
}

// Update actualiza un tipo de mantenimiento existente
func (r *tipoMantenimientoRepository) Update(tipo *models.TipoMantenimiento) error {
	return r.db.Save(tipo).Error
}

// Delete elimina un tipo de mantenimiento por su ID
func (r *tipoMantenimientoRepository) Delete(id uint) error {
	return r.db.Delete(&models.TipoMantenimiento{}, id).Error
}

// FindAll retorna todos los tipos de mantenimiento
func (r *tipoMantenimientoRepository) FindAll() ([]models.TipoMantenimiento, error) {
	var tipos []models.TipoMantenimiento
	err := r.db.Find(&tipos).Error
	return tipos, err
}

// FindByReporteID retorna todos los tipos de mantenimiento asociados a un reporte
func (r *tipoMantenimientoRepository) FindByReporteID(reporteID uint) ([]models.TipoMantenimiento, error) {
	var tipos []models.TipoMantenimiento
	err := r.db.Where("reporte_id = ?", reporteID).Find(&tipos).Error
	return tipos, err
}