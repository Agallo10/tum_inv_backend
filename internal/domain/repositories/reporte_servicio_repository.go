package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// ReporteServicioRepository define las operaciones del repositorio para ReporteServicio
type ReporteServicioRepository interface {
	Create(reporte *models.ReporteServicio) error
	FindByID(id uint) (*models.ReporteServicio, error)
	Update(reporte *models.ReporteServicio) error
	Delete(id uint) error
	FindAll() ([]models.ReporteServicio, error)
	FindByEquipoID(equipoID uint) ([]models.ReporteServicio, error)
}

// reporteServicioRepository implementa ReporteServicioRepository
type reporteServicioRepository struct {
	db *gorm.DB
}

// NewReporteServicioRepository crea una nueva instancia de ReporteServicioRepository
func NewReporteServicioRepository(db *gorm.DB) ReporteServicioRepository {
	return &reporteServicioRepository{db: db}
}

// Create crea un nuevo reporte de servicio en la base de datos
func (r *reporteServicioRepository) Create(reporte *models.ReporteServicio) error {
	return r.db.Create(reporte).Error
}

// FindByID busca un reporte de servicio por su ID
func (r *reporteServicioRepository) FindByID(id uint) (*models.ReporteServicio, error) {
	var reporte models.ReporteServicio
	err := r.db.Preload("TiposMantenimiento").Preload("Repuestos").Preload("Funcionarios").First(&reporte, id).Error
	if err != nil {
		return nil, err
	}
	return &reporte, nil
}

// Update actualiza un reporte de servicio existente
func (r *reporteServicioRepository) Update(reporte *models.ReporteServicio) error {
	return r.db.Save(reporte).Error
}

// Delete elimina un reporte de servicio por su ID
func (r *reporteServicioRepository) Delete(id uint) error {
	return r.db.Delete(&models.ReporteServicio{}, id).Error
}

// FindAll retorna todos los reportes de servicio
func (r *reporteServicioRepository) FindAll() ([]models.ReporteServicio, error) {
	var reportes []models.ReporteServicio
	err := r.db.Preload("TiposMantenimiento").Preload("Repuestos").Preload("Funcionarios").Find(&reportes).Error
	return reportes, err
}

// FindByEquipoID retorna todos los reportes de servicio asociados a un equipo
func (r *reporteServicioRepository) FindByEquipoID(equipoID uint) ([]models.ReporteServicio, error) {
	var reportes []models.ReporteServicio
	err := r.db.Preload("TiposMantenimiento").Preload("Repuestos").Preload("Funcionarios").Where("equipo_id = ?", equipoID).Find(&reportes).Error
	return reportes, err
}