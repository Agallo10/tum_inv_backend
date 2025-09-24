package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// HardwareInternoService define las operaciones del servicio para HardwareInterno
type HardwareInternoService interface {
	CreateHardwareInterno(hardware *models.HardwareInterno) error
	GetHardwareInternoByID(id uint) (*models.HardwareInterno, error)
	UpdateHardwareInterno(hardware *models.HardwareInterno) error
	DeleteHardwareInterno(id uint) error
	GetAllHardwareInterno() ([]models.HardwareInterno, error)
	GetHardwareInternoByEquipoID(equipoID uint) ([]models.HardwareInterno, error)
}

// hardwareInternoService implementa HardwareInternoService
type hardwareInternoService struct {
	hardwareRepo repositories.HardwareInternoRepository
}

// NewHardwareInternoService crea una nueva instancia de HardwareInternoService
func NewHardwareInternoService(hardwareRepo repositories.HardwareInternoRepository) HardwareInternoService {
	return &hardwareInternoService{hardwareRepo: hardwareRepo}
}

// CreateHardwareInterno crea un nuevo componente de hardware interno
func (s *hardwareInternoService) CreateHardwareInterno(hardware *models.HardwareInterno) error {
	if hardware.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if hardware.Componente == "" {
		return errors.New("el tipo de componente es obligatorio")
	}
	if hardware.Tecnologia == "" {
		return errors.New("la tecnología del componente es obligatoria")
	}
	if hardware.Capacidad == "" {
		return errors.New("la capacidad del componente es obligatoria")
	}
	return s.hardwareRepo.Create(hardware)
}

// GetHardwareInternoByID obtiene un componente de hardware interno por su ID
func (s *hardwareInternoService) GetHardwareInternoByID(id uint) (*models.HardwareInterno, error) {
	return s.hardwareRepo.FindByID(id)
}

// UpdateHardwareInterno actualiza un componente de hardware interno existente
func (s *hardwareInternoService) UpdateHardwareInterno(hardware *models.HardwareInterno) error {
	if hardware.ID == 0 {
		return errors.New("ID de hardware interno no válido")
	}
	if hardware.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if hardware.Componente == "" {
		return errors.New("el tipo de componente es obligatorio")
	}
	if hardware.Tecnologia == "" {
		return errors.New("la tecnología del componente es obligatoria")
	}
	if hardware.Capacidad == "" {
		return errors.New("la capacidad del componente es obligatoria")
	}
	return s.hardwareRepo.Update(hardware)
}

// DeleteHardwareInterno elimina un componente de hardware interno por su ID
func (s *hardwareInternoService) DeleteHardwareInterno(id uint) error {
	if id == 0 {
		return errors.New("ID de hardware interno no válido")
	}
	return s.hardwareRepo.Delete(id)
}

// GetAllHardwareInterno obtiene todos los componentes de hardware interno
func (s *hardwareInternoService) GetAllHardwareInterno() ([]models.HardwareInterno, error) {
	return s.hardwareRepo.FindAll()
}

// GetHardwareInternoByEquipoID obtiene todos los componentes de hardware interno asociados a un equipo
func (s *hardwareInternoService) GetHardwareInternoByEquipoID(equipoID uint) ([]models.HardwareInterno, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.hardwareRepo.FindByEquipoID(equipoID)
}