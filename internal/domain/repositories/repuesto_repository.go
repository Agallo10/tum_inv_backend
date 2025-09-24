package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// RepuestoRepository define las operaciones del repositorio para Repuesto
type RepuestoRepository interface {
	Create(repuesto *models.Repuesto) error
	FindByID(id uint) (*models.Repuesto, error)
	Update(repuesto *models.Repuesto) error
	Delete(id uint) error
	FindAll() ([]models.Repuesto, error)
	FindByReporteID(reporteID uint) ([]models.Repuesto, error)
}

// repuestoRepository implementa RepuestoRepository
type repuestoRepository struct {
	db *gorm.DB
}

// NewRepuestoRepository crea una nueva instancia de RepuestoRepository
func NewRepuestoRepository(db *gorm.DB) RepuestoRepository {
	return &repuestoRepository{db: db}
}

// Create crea un nuevo repuesto en la base de datos
func (r *repuestoRepository) Create(repuesto *models.Repuesto) error {
	return r.db.Create(repuesto).Error
}

// FindByID busca un repuesto por su ID
func (r *repuestoRepository) FindByID(id uint) (*models.Repuesto, error) {
	var repuesto models.Repuesto
	err := r.db.First(&repuesto, id).Error
	if err != nil {
		return nil, err
	}
	return &repuesto, nil
}

// Update actualiza un repuesto existente
func (r *repuestoRepository) Update(repuesto *models.Repuesto) error {
	return r.db.Save(repuesto).Error
}

// Delete elimina un repuesto por su ID
func (r *repuestoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Repuesto{}, id).Error
}

// FindAll retorna todos los repuestos
func (r *repuestoRepository) FindAll() ([]models.Repuesto, error) {
	var repuestos []models.Repuesto
	err := r.db.Find(&repuestos).Error
	return repuestos, err
}

// FindByReporteID retorna todos los repuestos asociados a un reporte
func (r *repuestoRepository) FindByReporteID(reporteID uint) ([]models.Repuesto, error) {
	var repuestos []models.Repuesto
	err := r.db.Where("reporte_id = ?", reporteID).Find(&repuestos).Error
	return repuestos, err
}