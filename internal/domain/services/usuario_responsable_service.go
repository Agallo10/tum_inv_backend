package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// UsuarioResponsableService define las operaciones del servicio para UsuarioResponsable
type UsuarioResponsableService interface {
	CreateUsuarioResponsable(usuario *models.UsuarioResponsable) error
	GetUsuarioResponsableByID(id uint) (*models.UsuarioResponsable, error)
	UpdateUsuarioResponsable(usuario *models.UsuarioResponsable) error
	DeleteUsuarioResponsable(id uint) error
	GetAllUsuariosResponsables() ([]models.UsuarioResponsable, error)
	GetUsuarioResponsableByCedula(cedula string) (*models.UsuarioResponsable, error)
	GetUsuariosByDependenciaID(dependenciaID uint) ([]models.UsuarioResponsable, error)
}

// usuarioResponsableService implementa UsuarioResponsableService
type usuarioResponsableService struct {
	usuarioRepo repositories.UsuarioResponsableRepository
}

// NewUsuarioResponsableService crea una nueva instancia de UsuarioResponsableService
func NewUsuarioResponsableService(usuarioRepo repositories.UsuarioResponsableRepository) UsuarioResponsableService {
	return &usuarioResponsableService{usuarioRepo: usuarioRepo}
}

// CreateUsuarioResponsable crea un nuevo usuario responsable
func (s *usuarioResponsableService) CreateUsuarioResponsable(usuario *models.UsuarioResponsable) error {
	if usuario.NombresApellidos == "" {
		return errors.New("los nombres y apellidos son obligatorios")
	}
	if usuario.Cedula == "" {
		return errors.New("la cédula es obligatoria")
	}

	// Verificar si ya existe un usuario con la misma cédula
	existingUser, err := s.usuarioRepo.FindByCedula(usuario.Cedula)
	if err == nil && existingUser != nil {
		return errors.New("ya existe un usuario con esta cédula")
	}

	return s.usuarioRepo.Create(usuario)
}

// GetUsuarioResponsableByID obtiene un usuario responsable por su ID
func (s *usuarioResponsableService) GetUsuarioResponsableByID(id uint) (*models.UsuarioResponsable, error) {
	return s.usuarioRepo.FindByID(id)
}

// UpdateUsuarioResponsable actualiza un usuario responsable existente
func (s *usuarioResponsableService) UpdateUsuarioResponsable(usuario *models.UsuarioResponsable) error {
	if usuario.ID == 0 {
		return errors.New("ID de usuario no válido")
	}
	if usuario.NombresApellidos == "" {
		return errors.New("los nombres y apellidos son obligatorios")
	}
	if usuario.Cedula == "" {
		return errors.New("la cédula es obligatoria")
	}
	// if usuario.CorreoInstitucional == "" {
	// 	return errors.New("el correo institucional es obligatorio")
	// }

	// Verificar si ya existe otro usuario con la misma cédula
	existingUser, err := s.usuarioRepo.FindByCedula(usuario.Cedula)
	if err == nil && existingUser != nil && existingUser.ID != usuario.ID {
		return errors.New("ya existe otro usuario con esta cédula")
	}

	return s.usuarioRepo.Update(usuario)
}

// DeleteUsuarioResponsable elimina un usuario responsable por su ID
func (s *usuarioResponsableService) DeleteUsuarioResponsable(id uint) error {
	if id == 0 {
		return errors.New("ID de usuario no válido")
	}

	// Verificar si el usuario tiene equipos asociados
	// usuario, err := s.usuarioRepo.FindByID(id)
	// if err != nil {
	// 	return err
	// }

	// if len(usuario.Equipos) > 0 {
	// 	return errors.New("no se puede eliminar el usuario porque tiene equipos asociados")
	// }

	return s.usuarioRepo.Delete(id)
}

// GetAllUsuariosResponsables obtiene todos los usuarios responsables
func (s *usuarioResponsableService) GetAllUsuariosResponsables() ([]models.UsuarioResponsable, error) {
	return s.usuarioRepo.FindAll()
}

// GetUsuarioResponsableByCedula obtiene un usuario responsable por su cédula
func (s *usuarioResponsableService) GetUsuarioResponsableByCedula(cedula string) (*models.UsuarioResponsable, error) {
	if cedula == "" {
		return nil, errors.New("la cédula es obligatoria")
	}
	return s.usuarioRepo.FindByCedula(cedula)
}

// GetUsuariosByDependenciaID obtiene todos los usuarios responsables de una dependencia
func (s *usuarioResponsableService) GetUsuariosByDependenciaID(dependenciaID uint) ([]models.UsuarioResponsable, error) {
	if dependenciaID == 0 {
		return nil, errors.New("ID de dependencia no válido")
	}
	return s.usuarioRepo.FindByDependenciaID(dependenciaID)
}
