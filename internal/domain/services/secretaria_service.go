package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// SecretariaService define la interfaz para operaciones de servicio de Secretaria
type SecretariaService interface {
	CreateSecretaria(secretaria *models.Secretaria) error
	GetSecretariaByID(id uint) (*models.Secretaria, error)
	GetAllSecretarias() ([]models.Secretaria, error)
	UpdateSecretaria(secretaria *models.Secretaria) error
	DeleteSecretaria(id uint) error
}

// secretariaService implementa SecretariaService
type secretariaService struct {
	repo repositories.SecretariaRepository
}

// NewSecretariaService crea un nuevo servicio de Secretaria
func NewSecretariaService(repo repositories.SecretariaRepository) SecretariaService {
	return &secretariaService{repo: repo}
}

// CreateSecretaria crea una nueva Secretaria
func (s *secretariaService) CreateSecretaria(secretaria *models.Secretaria) error {
	if secretaria.Nombre == "" {
		return errors.New("el nombre de la secretaría/ es obligatorio")
	}
	if secretaria.Descripcion == "" {
		return errors.New("la descripción de la secretaría/ es obligatoria")
	}
	if secretaria.Ubicacion == "" {
		return errors.New("la ubicación de la secretaria es obligatoria")
	}
	if secretaria.Secretario == "" {
		return errors.New("el nombre del secretario es obligatorio")
	}

	return s.repo.CreateSecretaria(secretaria)
}

// GetSecretariaByID obtiene una Secretaria por su ID
func (s *secretariaService) GetSecretariaByID(id uint) (*models.Secretaria, error) {
	return s.repo.GetSecretariaByID(id)
}

// GetAllSecretarias obtiene todas las Secretarias
func (s *secretariaService) GetAllSecretarias() ([]models.Secretaria, error) {
	return s.repo.GetAllSecretarias()
}

// UpdateSecretaria actualiza una Secretaria existente
func (s *secretariaService) UpdateSecretaria(secretaria *models.Secretaria) error {
	if secretaria.ID == 0 {
		return errors.New("ID de secretaría/ no válido")
	}
	if secretaria.Nombre == "" {
		return errors.New("el nombre de la secretaría/ es obligatorio")
	}
	if secretaria.Descripcion == "" {
		return errors.New("la descripción de la secretaría/ es obligatoria")
	}
	if secretaria.Ubicacion == "" {
		return errors.New("la ubicación de la secretaria es obligatoria")
	}
	if secretaria.Secretario == "" {
		return errors.New("el nombre del secretario es obligatorio")
	}

	return s.repo.UpdateSecretaria(secretaria)
}

// DeleteSecretaria elimina una Secretaria por su ID
func (s *secretariaService) DeleteSecretaria(id uint) error {
	return s.repo.DeleteSecretaria(id)
}
