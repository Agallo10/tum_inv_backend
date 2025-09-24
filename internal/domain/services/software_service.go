package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// SoftwareService define las operaciones del servicio para Software
type SoftwareService interface {
	CreateSoftware(software *models.Software) error
	GetSoftwareByID(id uint) (*models.Software, error)
	UpdateSoftware(software *models.Software) error
	DeleteSoftware(id uint) error
	GetAllSoftware() ([]models.Software, error)
	GetAllSoftwareByEquipoID(equipoID uint) ([]models.Software, error)
}

// softwareService implementa SoftwareService
type softwareService struct {
	softwareRepo repositories.SoftwareRepository
}

// NewSoftwareService crea una nueva instancia de SoftwareService
func NewSoftwareService(softwareRepo repositories.SoftwareRepository) SoftwareService {
	return &softwareService{softwareRepo: softwareRepo}
}

// CreateSoftware crea un nuevo software
func (s *softwareService) CreateSoftware(software *models.Software) error {
	if software.EquipoID == 0 {
		return errors.New("el Id del equipo es obligatorio")
	}
	if software.Nombre == "" {
		return errors.New("el nombre del software es obligatorio")
	}
	return s.softwareRepo.Create(software)
}

// GetSoftwareByID obtiene un software por ID
func (s *softwareService) GetSoftwareByID(id uint) (*models.Software, error) {
	return s.softwareRepo.FindByID(id)
}

// UpdateSoftware actualiza un software existente
func (s *softwareService) UpdateSoftware(software *models.Software) error {
	if software.ID == 0 {
		return errors.New(("ID de software no válido"))
	}
	return s.softwareRepo.Update(software)
}

// DeleteSoftware elimina un software por su ID
func (s *softwareService) DeleteSoftware(id uint) error {
	if id == 0 {
		return errors.New("ID de software no válido")
	}
	return s.softwareRepo.Delete(id)
}

// GetAllSoftware obtiene todos los software
func (s *softwareService) GetAllSoftware() ([]models.Software, error) {
	return s.softwareRepo.FindAll()
}

// GetAllSoftwareByEquipoID obtiene todos los software asociados a un equipo
func (s *softwareService) GetAllSoftwareByEquipoID(equipoID uint) ([]models.Software, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no valido")
	}
	return s.softwareRepo.FindByEquipoID(equipoID)
}
