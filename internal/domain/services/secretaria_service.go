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
	repo    repositories.SecretariaRepository
	depRepo repositories.DependenciaRepository
}

// NewSecretariaService crea un nuevo servicio de Secretaria
func NewSecretariaService(repo repositories.SecretariaRepository, depRepo repositories.DependenciaRepository) SecretariaService {
	return &secretariaService{repo: repo, depRepo: depRepo}
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

// DeleteSecretaria elimina una Secretaria por su ID, liberando usuarios y eliminando dependencias
func (s *secretariaService) DeleteSecretaria(id uint) error {
	// Verificar que la secretaría existe y cargar sus dependencias
	secretaria, err := s.repo.GetSecretariaByID(id)
	if err != nil {
		return errors.New("secretaría no encontrada")
	}

	// Para cada dependencia, liberar usuarios responsables y luego eliminar la dependencia
	for _, dep := range secretaria.Dependencias {
		// Liberar usuarios responsables (poner dependencia_id en NULL no es posible porque es NOT NULL)
		// Así que simplemente eliminamos la dependencia — los usuarios quedan intactos
		// porque GORM usa soft delete y no hay ON DELETE CASCADE en la FK
		if err := s.depRepo.LiberarUsuariosDeDependencia(dep.ID); err != nil {
			return errors.New("error al liberar usuarios de la dependencia: " + dep.Nombre)
		}
		if err := s.depRepo.DeleteDependencia(dep.ID); err != nil {
			return errors.New("error al eliminar la dependencia: " + dep.Nombre)
		}
	}

	return s.repo.DeleteSecretaria(id)
}
