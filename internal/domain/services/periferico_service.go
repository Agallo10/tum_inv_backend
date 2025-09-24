package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// PerifericoService define las operaciones del servicio para Periferico
type PerifericoService interface {
	CreatePeriferico(periferico *models.Periferico) error
	GetPerifericoByID(id uint) (*models.Periferico, error)
	UpdatePeriferico(periferico *models.Periferico) error
	DeletePeriferico(id uint) error
	GetAllPerifericos() ([]models.Periferico, error)
	GetPerifericosByEquipoID(equipoID uint) ([]models.Periferico, error)
}

// perifericoService implementa PerifericoService
type perifericoService struct {
	perifericoRepo repositories.PerifericoRepository
}

// NewPerifericoService crea una nueva instancia de PerifericoService
func NewPerifericoService(perifericoRepo repositories.PerifericoRepository) PerifericoService {
	return &perifericoService{perifericoRepo: perifericoRepo}
}

// CreatePeriferico crea un nuevo periférico
func (s *perifericoService) CreatePeriferico(periferico *models.Periferico) error {
	if periferico.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if periferico.TipoPeriferico == "" {
		return errors.New("el tipo de periférico es obligatorio")
	}
	return s.perifericoRepo.Create(periferico)
}

// GetPerifericoByID obtiene un periférico por su ID
func (s *perifericoService) GetPerifericoByID(id uint) (*models.Periferico, error) {
	return s.perifericoRepo.FindByID(id)
}

// UpdatePeriferico actualiza un periférico existente
func (s *perifericoService) UpdatePeriferico(periferico *models.Periferico) error {
	if periferico.ID == 0 {
		return errors.New("ID de periférico no válido")
	}
	return s.perifericoRepo.Update(periferico)
}

// DeletePeriferico elimina un periférico por su ID
func (s *perifericoService) DeletePeriferico(id uint) error {
	if id == 0 {
		return errors.New("ID de periférico no válido")
	}
	return s.perifericoRepo.Delete(id)
}

// GetAllPerifericos obtiene todos los periféricos
func (s *perifericoService) GetAllPerifericos() ([]models.Periferico, error) {
	return s.perifericoRepo.FindAll()
}

// GetPerifericosByEquipoID obtiene todos los periféricos asociados a un equipo
func (s *perifericoService) GetPerifericosByEquipoID(equipoID uint) ([]models.Periferico, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.perifericoRepo.FindByEquipoID(equipoID)
}