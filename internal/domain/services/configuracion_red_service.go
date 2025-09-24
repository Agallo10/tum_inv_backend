package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// ConfiguracionRedService define las operaciones del servicio para ConfiguracionRed
type ConfiguracionRedService interface {
	CreateConfiguracionRed(configuracion *models.ConfiguracionRed) error
	GetConfiguracionRedByID(id uint) (*models.ConfiguracionRed, error)
	UpdateConfiguracionRed(configuracion *models.ConfiguracionRed) error
	DeleteConfiguracionRed(id uint) error
	GetAllConfiguracionesRed() ([]models.ConfiguracionRed, error)
	GetConfiguracionRedByEquipoID(equipoID uint) (*models.ConfiguracionRed, error)
}

// configuracionRedService implementa ConfiguracionRedService
type configuracionRedService struct {
	configuracionRepo repositories.ConfiguracionRedRepository
}

// NewConfiguracionRedService crea una nueva instancia de ConfiguracionRedService
func NewConfiguracionRedService(configuracionRepo repositories.ConfiguracionRedRepository) ConfiguracionRedService {
	return &configuracionRedService{configuracionRepo: configuracionRepo}
}

// CreateConfiguracionRed crea una nueva configuración de red
func (s *configuracionRedService) CreateConfiguracionRed(configuracion *models.ConfiguracionRed) error {
	if configuracion.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if configuracion.DireccionIP == "" {
		return errors.New("la dirección IP es obligatoria")
	}
	if configuracion.NombreDispositivo == "" {
		return errors.New("el nombre del dispositivo es obligatorio")
	}
	
	// Verificar si ya existe una configuración para este equipo
	existente, err := s.configuracionRepo.FindByEquipoID(configuracion.EquipoID)
	if err == nil && existente != nil {
		return errors.New("ya existe una configuración de red para este equipo")
	}
	
	return s.configuracionRepo.Create(configuracion)
}

// GetConfiguracionRedByID obtiene una configuración de red por su ID
func (s *configuracionRedService) GetConfiguracionRedByID(id uint) (*models.ConfiguracionRed, error) {
	return s.configuracionRepo.FindByID(id)
}

// UpdateConfiguracionRed actualiza una configuración de red existente
func (s *configuracionRedService) UpdateConfiguracionRed(configuracion *models.ConfiguracionRed) error {
	if configuracion.ID == 0 {
		return errors.New("ID de configuración de red no válido")
	}
	if configuracion.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if configuracion.DireccionIP == "" {
		return errors.New("la dirección IP es obligatoria")
	}
	if configuracion.NombreDispositivo == "" {
		return errors.New("el nombre del dispositivo es obligatorio")
	}
	
	// Verificar si existe la configuración
	existente, err := s.configuracionRepo.FindByID(configuracion.ID)
	if err != nil {
		return errors.New("configuración de red no encontrada")
	}
	
	// Verificar si al cambiar el EquipoID no se genera conflicto con otra configuración
	if existente.EquipoID != configuracion.EquipoID {
		otro, err := s.configuracionRepo.FindByEquipoID(configuracion.EquipoID)
		if err == nil && otro != nil && otro.ID != configuracion.ID {
			return errors.New("ya existe una configuración de red para el equipo destino")
		}
	}
	
	return s.configuracionRepo.Update(configuracion)
}

// DeleteConfiguracionRed elimina una configuración de red por su ID
func (s *configuracionRedService) DeleteConfiguracionRed(id uint) error {
	if id == 0 {
		return errors.New("ID de configuración de red no válido")
	}
	return s.configuracionRepo.Delete(id)
}

// GetAllConfiguracionesRed obtiene todas las configuraciones de red
func (s *configuracionRedService) GetAllConfiguracionesRed() ([]models.ConfiguracionRed, error) {
	return s.configuracionRepo.FindAll()
}

// GetConfiguracionRedByEquipoID obtiene la configuración de red asociada a un equipo
func (s *configuracionRedService) GetConfiguracionRedByEquipoID(equipoID uint) (*models.ConfiguracionRed, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.configuracionRepo.FindByEquipoID(equipoID)
}