package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// DependenciaService define la interfaz para operaciones de servicio de Dependencia
type DependenciaService interface {
	CreateDependencia(dependencia *models.Dependencia) error
	GetDependenciaByID(id uint) (*models.Dependencia, error)
	GetAllDependencias() ([]models.Dependencia, error)
	UpdateDependencia(dependencia *models.Dependencia) error
	DeleteDependencia(id uint) error
	GetDependenciasBySecretariaID(secretariaID uint) ([]models.Dependencia, error)
}

// dependenciaService implementa DependenciaService
type dependenciaService struct {
	repo repositories.DependenciaRepository
}

// NewDependenciaService crea una nueva instancia de DependenciaService
func NewDependenciaService(repo repositories.DependenciaRepository) DependenciaService {
	return &dependenciaService{repo: repo}
}

// CreateDependencia crea una nueva dependencia
func (s *dependenciaService) CreateDependencia(dependencia *models.Dependencia) error {
	if dependencia.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if dependencia.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}
	if dependencia.UbicacionOficina == "" {
		return errors.New("la ubicación de la oficina es obligatoria")
	}
	if dependencia.JefeOficina == "" {
		return errors.New("el nombre del jefe de oficina es obligatorio")
	}
	if dependencia.SecretariaID == 0 {
		return errors.New("la secretaría es obligatoria")
	}

	if dependencia.CorreoInstitucional == "" {
		return errors.New("el correo institucional es obligatorio")
	}

	return s.repo.CreateDependencia(dependencia)
}

// GetDependenciaByID obtiene una dependencia por su ID
func (s *dependenciaService) GetDependenciaByID(id uint) (*models.Dependencia, error) {
	return s.repo.GetDependenciaByID(id)
}

// GetAllDependencias obtiene todas las dependencias
func (s *dependenciaService) GetAllDependencias() ([]models.Dependencia, error) {
	return s.repo.GetAllDependencias()
}

// UpdateDependencia actualiza una dependencia existente
func (s *dependenciaService) UpdateDependencia(dependencia *models.Dependencia) error {
	if dependencia.ID == 0 {
		return errors.New("ID de dependencia no válido")
	}
	if dependencia.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if dependencia.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}
	if dependencia.UbicacionOficina == "" {
		return errors.New("la ubicación de la oficina es obligatoria")
	}
	if dependencia.SecretariaID == 0 {
		return errors.New("la secretaría es obligatoria")
	}

	return s.repo.UpdateDependencia(dependencia)
}

// DeleteDependencia elimina una dependencia por su ID
func (s *dependenciaService) DeleteDependencia(id uint) error {
	if id == 0 {
		return errors.New("ID de dependencia no válido")
	}

	// Verificar si la dependencia tiene usuarios responsables asociados
	dependencia, err := s.repo.GetDependenciaByID(id)
	if err != nil {
		return err
	}

	if len(dependencia.UsuarioResponsables) > 0 {
		return errors.New("no se puede eliminar la dependencia porque tiene usuarios responsables asociados")
	}

	return s.repo.DeleteDependencia(id)
}

// GetDependenciasBySecretariaID obtiene todas las dependencias de una secretaría
func (s *dependenciaService) GetDependenciasBySecretariaID(secretariaID uint) ([]models.Dependencia, error) {
	if secretariaID == 0 {
		return nil, errors.New("ID de secretaría no válido")
	}
	return s.repo.GetDependenciasBySecretariaID(secretariaID)
}
