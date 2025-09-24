package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// RepuestoService define las operaciones del servicio para Repuesto
type RepuestoService interface {
	CreateRepuesto(repuesto *models.Repuesto) error
	GetRepuestoByID(id uint) (*models.Repuesto, error)
	UpdateRepuesto(repuesto *models.Repuesto) error
	DeleteRepuesto(id uint) error
	GetAllRepuestos() ([]models.Repuesto, error)
	GetRepuestosByReporteID(reporteID uint) ([]models.Repuesto, error)
}

// repuestoService implementa RepuestoService
type repuestoService struct {
	repuestoRepo repositories.RepuestoRepository
}

// NewRepuestoService crea una nueva instancia de RepuestoService
func NewRepuestoService(repuestoRepo repositories.RepuestoRepository) RepuestoService {
	return &repuestoService{repuestoRepo: repuestoRepo}
}

// CreateRepuesto crea un nuevo repuesto
func (s *repuestoService) CreateRepuesto(repuesto *models.Repuesto) error {
	if repuesto.SerialNumeroParte == "" {
		return errors.New("el serial o número de parte es obligatorio")
	}
	if repuesto.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}
	if repuesto.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}

	return s.repuestoRepo.Create(repuesto)
}

// GetRepuestoByID obtiene un repuesto por su ID
func (s *repuestoService) GetRepuestoByID(id uint) (*models.Repuesto, error) {
	return s.repuestoRepo.FindByID(id)
}

// UpdateRepuesto actualiza un repuesto existente
func (s *repuestoService) UpdateRepuesto(repuesto *models.Repuesto) error {
	if repuesto.ID == 0 {
		return errors.New("ID de repuesto no válido")
	}
	if repuesto.SerialNumeroParte == "" {
		return errors.New("el serial o número de parte es obligatorio")
	}
	if repuesto.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}
	if repuesto.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}

	// Verificar si existe el repuesto
	existente, err := s.repuestoRepo.FindByID(repuesto.ID)
	if err != nil && existente != nil {
		return errors.New("repuesto no encontrado")
	}

	return s.repuestoRepo.Update(repuesto)
}

// DeleteRepuesto elimina un repuesto por su ID
func (s *repuestoService) DeleteRepuesto(id uint) error {
	if id == 0 {
		return errors.New("ID de repuesto no válido")
	}
	return s.repuestoRepo.Delete(id)
}

// GetAllRepuestos obtiene todos los repuestos
func (s *repuestoService) GetAllRepuestos() ([]models.Repuesto, error) {
	return s.repuestoRepo.FindAll()
}

// GetRepuestosByReporteID obtiene todos los repuestos asociados a un reporte
func (s *repuestoService) GetRepuestosByReporteID(reporteID uint) ([]models.Repuesto, error) {
	if reporteID == 0 {
		return nil, errors.New("ID de reporte no válido")
	}
	return s.repuestoRepo.FindByReporteID(reporteID)
}
