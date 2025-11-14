package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/models/dto"
	"tum_inv_backend/internal/domain/repositories"
)

// EquipoService define las operaciones del servicio para Equipo
type EquipoService interface {
	CreateEquipo(equipo *models.Equipo) error
	GetEquipoByID(id uint) (*models.Equipo, error)
	UpdateEquipo(equipo *models.Equipo) error
	DeleteEquipo(id uint) error
	GetAllEquipos() ([]models.Equipo, error)
	GetEquiposByDependenciaID(dependenciaID uint) ([]models.Equipo, error)
	GetEquipoUsuDepByID(equipoID uint) (dto.EquipoConResponsableDTO, error)
	GetAllEquiposDetalle() ([]dto.EquipoConResponsableDTO, error)
}

// equipoService implementa EquipoService
type equipoService struct {
	equipoRepo repositories.EquipoRepository
}

// NewEquipoService crea una nueva instancia de EquipoService
func NewEquipoService(equipoRepo repositories.EquipoRepository) EquipoService {
	return &equipoService{equipoRepo: equipoRepo}
}

// CreateEquipo crea un nuevo equipo
func (s *equipoService) CreateEquipo(equipo *models.Equipo) error {
	if equipo.Serial == "" {
		return errors.New("el número de serie es obligatorio")
	}
	if equipo.Marca == "" {
		return errors.New("la marca es obligatoria")
	}
	return s.equipoRepo.Create(equipo)
}

// GetEquipoByID obtiene un equipo por su ID
func (s *equipoService) GetEquipoByID(id uint) (*models.Equipo, error) {
	return s.equipoRepo.FindByID(id)
}

// UpdateEquipo actualiza un equipo existente
func (s *equipoService) UpdateEquipo(equipo *models.Equipo) error {
	if equipo.ID == 0 {
		return errors.New("ID de equipo no válido")
	}
	return s.equipoRepo.Update(equipo)
}

// DeleteEquipo elimina un equipo por su ID
func (s *equipoService) DeleteEquipo(id uint) error {
	if id == 0 {
		return errors.New("ID de equipo no válido")
	}
	return s.equipoRepo.Delete(id)
}

// GetAllEquipos obtiene todos los equipos
func (s *equipoService) GetAllEquipos() ([]models.Equipo, error) {
	return s.equipoRepo.FindAll()
}

// GetEquiposByDependenciaID obtiene todos los equipos de una dependencia
func (s *equipoService) GetEquiposByDependenciaID(dependenciaID uint) ([]models.Equipo, error) {
	if dependenciaID == 0 {
		return nil, errors.New("ID de dependencia no válido")
	}
	return s.equipoRepo.FindByDependenciaID(dependenciaID)
}

// GetEquiposByDependenciaID obtiene todos los equipos de una dependencia
func (s *equipoService) GetEquipoUsuDepByID(equipoID uint) (dto.EquipoConResponsableDTO, error) {

	return s.equipoRepo.FindEquiUsuDepByID(equipoID)
}

func (s *equipoService) GetAllEquiposDetalle() ([]dto.EquipoConResponsableDTO, error) {
	return s.equipoRepo.FindAllEquiposDetalle()
}
