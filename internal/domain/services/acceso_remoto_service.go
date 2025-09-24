package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// AccesoRemotoService define las operaciones del servicio para AccesoRemoto
type AccesoRemotoService interface {
	CreateAccesoRemoto(acceso *models.AccesoRemoto) error
	GetAccesoRemotoByID(id uint) (*models.AccesoRemoto, error)
	UpdateAccesoRemoto(acceso *models.AccesoRemoto) error
	DeleteAccesoRemoto(id uint) error
	GetAllAccesosRemotos() ([]models.AccesoRemoto, error)
	GetAccesosRemotosByEquipoID(equipoID uint) ([]models.AccesoRemoto, error)
}

// accesoRemotoService implementa AccesoRemotoService
type accesoRemotoService struct {
	accesoRepo repositories.AccesoRemotoRepository
}

// NewAccesoRemotoService crea una nueva instancia de AccesoRemotoService
func NewAccesoRemotoService(accesoRepo repositories.AccesoRemotoRepository) AccesoRemotoService {
	return &accesoRemotoService{accesoRepo: accesoRepo}
}

// CreateAccesoRemoto crea un nuevo acceso remoto
func (s *accesoRemotoService) CreateAccesoRemoto(acceso *models.AccesoRemoto) error {
	if acceso.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if acceso.Usuario == "" {
		return errors.New("el usuario es obligatorio")
	}
	if acceso.IDConexion == "" {
		return errors.New("el ID de conexión es obligatorio")
	}

	return s.accesoRepo.Create(acceso)
}

// GetAccesoRemotoByID obtiene un acceso remoto por su ID
func (s *accesoRemotoService) GetAccesoRemotoByID(id uint) (*models.AccesoRemoto, error) {
	return s.accesoRepo.FindByID(id)
}

// UpdateAccesoRemoto actualiza un acceso remoto existente
func (s *accesoRemotoService) UpdateAccesoRemoto(acceso *models.AccesoRemoto) error {
	if acceso.ID == 0 {
		return errors.New("ID de acceso remoto no válido")
	}
	if acceso.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if acceso.Usuario == "" {
		return errors.New("el usuario es obligatorio")
	}
	if acceso.IDConexion == "" {
		return errors.New("el ID de conexión es obligatorio")
	}

	// Verificar si existe el acceso remoto
	existente, err := s.accesoRepo.FindByID(acceso.ID)
	if err != nil && existente != nil {
		return errors.New("acceso remoto no encontrado")
	}

	return s.accesoRepo.Update(acceso)
}

// DeleteAccesoRemoto elimina un acceso remoto por su ID
func (s *accesoRemotoService) DeleteAccesoRemoto(id uint) error {
	if id == 0 {
		return errors.New("ID de acceso remoto no válido")
	}
	return s.accesoRepo.Delete(id)
}

// GetAllAccesosRemotos obtiene todos los accesos remotos
func (s *accesoRemotoService) GetAllAccesosRemotos() ([]models.AccesoRemoto, error) {
	return s.accesoRepo.FindAll()
}

// GetAccesosRemotosByEquipoID obtiene todos los accesos remotos asociados a un equipo
func (s *accesoRemotoService) GetAccesosRemotosByEquipoID(equipoID uint) ([]models.AccesoRemoto, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.accesoRepo.FindByEquipoID(equipoID)
}
