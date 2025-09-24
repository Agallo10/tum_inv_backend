package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// ReporteServicioService define las operaciones del servicio para ReporteServicio
type ReporteServicioService interface {
	CreateReporteServicio(reporte *models.ReporteServicio) error
	GetReporteServicioByID(id uint) (*models.ReporteServicio, error)
	UpdateReporteServicio(reporte *models.ReporteServicio) error
	DeleteReporteServicio(id uint) error
	GetAllReportesServicio() ([]models.ReporteServicio, error)
	GetReportesServicioByEquipoID(equipoID uint) ([]models.ReporteServicio, error)
}

// reporteServicioService implementa ReporteServicioService
type reporteServicioService struct {
	reporteRepo repositories.ReporteServicioRepository
}

// NewReporteServicioService crea una nueva instancia de ReporteServicioService
func NewReporteServicioService(reporteRepo repositories.ReporteServicioRepository) ReporteServicioService {
	return &reporteServicioService{reporteRepo: reporteRepo}
}

// CreateReporteServicio crea un nuevo reporte de servicio
func (s *reporteServicioService) CreateReporteServicio(reporte *models.ReporteServicio) error {
	if reporte.Dependencia == "" {
		return errors.New("la dependencia es obligatoria")
	}
	if reporte.Ubicacion == "" {
		return errors.New("la ubicación es obligatoria")
	}
	if reporte.ActividadRealizada == "" {
		return errors.New("la actividad realizada es obligatoria")
	}

	return s.reporteRepo.Create(reporte)
}

// GetReporteServicioByID obtiene un reporte de servicio por su ID
func (s *reporteServicioService) GetReporteServicioByID(id uint) (*models.ReporteServicio, error) {
	return s.reporteRepo.FindByID(id)
}

// UpdateReporteServicio actualiza un reporte de servicio existente
func (s *reporteServicioService) UpdateReporteServicio(reporte *models.ReporteServicio) error {
	if reporte.ID == 0 {
		return errors.New("ID de reporte no válido")
	}
	if reporte.Dependencia == "" {
		return errors.New("la dependencia es obligatoria")
	}
	if reporte.Ubicacion == "" {
		return errors.New("la ubicación es obligatoria")
	}
	if reporte.ActividadRealizada == "" {
		return errors.New("la actividad realizada es obligatoria")
	}

	// Verificar si existe el reporte
	existente, err := s.reporteRepo.FindByID(reporte.ID)
	if err != nil && existente != nil {
		return errors.New("reporte no encontrado")
	}

	return s.reporteRepo.Update(reporte)
}

// DeleteReporteServicio elimina un reporte de servicio por su ID
func (s *reporteServicioService) DeleteReporteServicio(id uint) error {
	if id == 0 {
		return errors.New("ID de reporte no válido")
	}
	return s.reporteRepo.Delete(id)
}

// GetAllReportesServicio obtiene todos los reportes de servicio
func (s *reporteServicioService) GetAllReportesServicio() ([]models.ReporteServicio, error) {
	return s.reporteRepo.FindAll()
}

// GetReportesServicioByEquipoID obtiene todos los reportes de servicio asociados a un equipo
func (s *reporteServicioService) GetReportesServicioByEquipoID(equipoID uint) ([]models.ReporteServicio, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.reporteRepo.FindByEquipoID(equipoID)
}
