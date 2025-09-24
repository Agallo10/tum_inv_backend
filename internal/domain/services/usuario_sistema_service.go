package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// UsuarioSistemaService define las operaciones del servicio para UsuarioSistema
type UsuarioSistemaService interface {
	CreateUsuarioSistema(usuario *models.UsuarioSistema) error
	GetUsuarioSistemaByID(id uint) (*models.UsuarioSistema, error)
	UpdateUsuarioSistema(usuario *models.UsuarioSistema) error
	DeleteUsuarioSistema(id uint) error
	GetAllUsuariosSistema() ([]models.UsuarioSistema, error)
	GetUsuariosSistemaByEquipoID(equipoID uint) ([]models.UsuarioSistema, error)
	GetUsuarioSistemaByNombreUsuario(nombreUsuario string, equipoID uint) (*models.UsuarioSistema, error)
}

// usuarioSistemaService implementa UsuarioSistemaService
type usuarioSistemaService struct {
	usuarioRepo repositories.UsuarioSistemaRepository
}

// NewUsuarioSistemaService crea una nueva instancia de UsuarioSistemaService
func NewUsuarioSistemaService(usuarioRepo repositories.UsuarioSistemaRepository) UsuarioSistemaService {
	return &usuarioSistemaService{usuarioRepo: usuarioRepo}
}

// CreateUsuarioSistema crea un nuevo usuario del sistema
func (s *usuarioSistemaService) CreateUsuarioSistema(usuario *models.UsuarioSistema) error {
	if usuario.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if usuario.NombreUsuario == "" {
		return errors.New("el nombre de usuario es obligatorio")
	}
	
	// Verificar si ya existe un usuario con el mismo nombre en el mismo equipo
	existente, err := s.usuarioRepo.FindByNombreUsuario(usuario.NombreUsuario, usuario.EquipoID)
	if err == nil && existente != nil {
		return errors.New("ya existe un usuario con este nombre en el equipo")
	}
	
	return s.usuarioRepo.Create(usuario)
}

// GetUsuarioSistemaByID obtiene un usuario del sistema por su ID
func (s *usuarioSistemaService) GetUsuarioSistemaByID(id uint) (*models.UsuarioSistema, error) {
	return s.usuarioRepo.FindByID(id)
}

// UpdateUsuarioSistema actualiza un usuario del sistema existente
func (s *usuarioSistemaService) UpdateUsuarioSistema(usuario *models.UsuarioSistema) error {
	if usuario.ID == 0 {
		return errors.New("ID de usuario no válido")
	}
	if usuario.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if usuario.NombreUsuario == "" {
		return errors.New("el nombre de usuario es obligatorio")
	}
	
	// Verificar si existe el usuario
	existente, err := s.usuarioRepo.FindByID(usuario.ID)
	if err != nil {
		return errors.New("usuario no encontrado")
	}
	
	// Verificar si al cambiar el nombre de usuario no se genera conflicto con otro usuario en el mismo equipo
	if existente.NombreUsuario != usuario.NombreUsuario || existente.EquipoID != usuario.EquipoID {
		otro, err := s.usuarioRepo.FindByNombreUsuario(usuario.NombreUsuario, usuario.EquipoID)
		if err == nil && otro != nil && otro.ID != usuario.ID {
			return errors.New("ya existe un usuario con este nombre en el equipo")
		}
	}
	
	return s.usuarioRepo.Update(usuario)
}

// DeleteUsuarioSistema elimina un usuario del sistema por su ID
func (s *usuarioSistemaService) DeleteUsuarioSistema(id uint) error {
	if id == 0 {
		return errors.New("ID de usuario no válido")
	}
	return s.usuarioRepo.Delete(id)
}

// GetAllUsuariosSistema obtiene todos los usuarios del sistema
func (s *usuarioSistemaService) GetAllUsuariosSistema() ([]models.UsuarioSistema, error) {
	return s.usuarioRepo.FindAll()
}

// GetUsuariosSistemaByEquipoID obtiene todos los usuarios del sistema asociados a un equipo
func (s *usuarioSistemaService) GetUsuariosSistemaByEquipoID(equipoID uint) ([]models.UsuarioSistema, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.usuarioRepo.FindByEquipoID(equipoID)
}

// GetUsuarioSistemaByNombreUsuario obtiene un usuario del sistema por su nombre de usuario y equipo
func (s *usuarioSistemaService) GetUsuarioSistemaByNombreUsuario(nombreUsuario string, equipoID uint) (*models.UsuarioSistema, error) {
	if nombreUsuario == "" {
		return nil, errors.New("el nombre de usuario es obligatorio")
	}
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.usuarioRepo.FindByNombreUsuario(nombreUsuario, equipoID)
}