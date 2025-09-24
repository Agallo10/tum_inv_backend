package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// TipoMantenimientoService define las operaciones del servicio para TipoMantenimiento
type TipoMantenimientoService interface {
	CreateTipoMantenimiento(tipo *models.TipoMantenimiento) error
	GetTipoMantenimientoByID(id uint) (*models.TipoMantenimiento, error)
	UpdateTipoMantenimiento(tipo *models.TipoMantenimiento) error
	DeleteTipoMantenimiento(id uint) error
	GetAllTiposMantenimiento() ([]models.TipoMantenimiento, error)
	GetTiposMantenimientoByReporteID(reporteID uint) ([]models.TipoMantenimiento, error)
}

// tipoMantenimientoService implementa TipoMantenimientoService
type tipoMantenimientoService struct {
	tipoRepo repositories.TipoMantenimientoRepository
}

// NewTipoMantenimientoService crea una nueva instancia de TipoMantenimientoService
func NewTipoMantenimientoService(tipoRepo repositories.TipoMantenimientoRepository) TipoMantenimientoService {
	return &tipoMantenimientoService{tipoRepo: tipoRepo}
}

// CreateTipoMantenimiento crea un nuevo tipo de mantenimiento
func (s *tipoMantenimientoService) CreateTipoMantenimiento(tipo *models.TipoMantenimiento) error {
	if tipo.ReporteID == 0 {
		return errors.New("el ID del reporte es obligatorio")
	}
	if tipo.Tipo == "" {
		return errors.New("el tipo de mantenimiento es obligatorio")
	}

	return s.tipoRepo.Create(tipo)
}

// GetTipoMantenimientoByID obtiene un tipo de mantenimiento por su ID
func (s *tipoMantenimientoService) GetTipoMantenimientoByID(id uint) (*models.TipoMantenimiento, error) {
	return s.tipoRepo.FindByID(id)
}

// UpdateTipoMantenimiento actualiza un tipo de mantenimiento existente
func (s *tipoMantenimientoService) UpdateTipoMantenimiento(tipo *models.TipoMantenimiento) error {
	if tipo.ID == 0 {
		return errors.New("ID de tipo de mantenimiento no válido")
	}
	if tipo.ReporteID == 0 {
		return errors.New("el ID del reporte es obligatorio")
	}
	if tipo.Tipo == "" {
		return errors.New("el tipo de mantenimiento es obligatorio")
	}

	// Verificar si existe el tipo de mantenimiento
	existente, err := s.tipoRepo.FindByID(tipo.ID)
	if err != nil && existente != nil {
		return errors.New("tipo de mantenimiento no encontrado")
	}

	return s.tipoRepo.Update(tipo)
}

// DeleteTipoMantenimiento elimina un tipo de mantenimiento por su ID
func (s *tipoMantenimientoService) DeleteTipoMantenimiento(id uint) error {
	if id == 0 {
		return errors.New("ID de tipo de mantenimiento no válido")
	}
	return s.tipoRepo.Delete(id)
}

// GetAllTiposMantenimiento obtiene todos los tipos de mantenimiento
func (s *tipoMantenimientoService) GetAllTiposMantenimiento() ([]models.TipoMantenimiento, error) {
	return s.tipoRepo.FindAll()
}

// GetTiposMantenimientoByReporteID obtiene todos los tipos de mantenimiento asociados a un reporte
func (s *tipoMantenimientoService) GetTiposMantenimientoByReporteID(reporteID uint) ([]models.TipoMantenimiento, error) {
	if reporteID == 0 {
		return nil, errors.New("ID de reporte no válido")
	}
	return s.tipoRepo.FindByReporteID(reporteID)
}
