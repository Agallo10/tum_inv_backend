package services

import (
	"errors"
	"strings"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

type EstadoEquipoService struct {
	repo *repositories.EstadoEquipoRepository
}

func NewEstadoEquipoService(repo *repositories.EstadoEquipoRepository) *EstadoEquipoService {
	return &EstadoEquipoService{repo: repo}
}

// GetAllEstados obtiene todos los estados de equipo
func (s *EstadoEquipoService) GetAllEstados() ([]models.EstadoEquipo, error) {
	return s.repo.GetAll()
}

// GetEstadoByID obtiene un estado de equipo por su ID
func (s *EstadoEquipoService) GetEstadoByID(id uint) (*models.EstadoEquipo, error) {
	if id == 0 {
		return nil, errors.New("ID no válido")
	}
	return s.repo.GetByID(id)
}

// GetActiveEstados obtiene todos los estados de equipo activos
func (s *EstadoEquipoService) GetActiveEstados() ([]models.EstadoEquipo, error) {
	return s.repo.GetActive()
}

// CreateEstado crea un nuevo estado de equipo con validaciones
func (s *EstadoEquipoService) CreateEstado(estado *models.EstadoEquipo) error {
	// Validaciones
	if err := s.validateEstado(estado); err != nil {
		return err
	}

	// Verificar que no exista otro estado con el mismo nombre
	exists, err := s.repo.ExistsByName(estado.Nombre)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ya existe un estado de equipo con ese nombre")
	}

	// Limpiar y formatear datos
	estado.Nombre = strings.TrimSpace(estado.Nombre)
	estado.Descripcion = strings.TrimSpace(estado.Descripcion)

	return s.repo.Create(estado)
}

// UpdateEstado actualiza un estado de equipo existente
func (s *EstadoEquipoService) UpdateEstado(id uint, estado *models.EstadoEquipo) error {
	if id == 0 {
		return errors.New("ID no válido")
	}

	// Verificar que el estado existe
	existingEstado, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("estado de equipo no encontrado")
	}

	// Validaciones
	if err := s.validateEstado(estado); err != nil {
		return err
	}

	// Verificar que no exista otro estado con el mismo nombre (excluyendo el actual)
	exists, err := s.repo.ExistsByNameExcludingID(estado.Nombre, id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ya existe otro estado de equipo con ese nombre")
	}

	// Actualizar campos
	existingEstado.Nombre = strings.TrimSpace(estado.Nombre)
	existingEstado.Descripcion = strings.TrimSpace(estado.Descripcion)
	existingEstado.Activo = estado.Activo

	return s.repo.Update(existingEstado)
}

// DeleteEstado elimina un estado de equipo
func (s *EstadoEquipoService) DeleteEstado(id uint) error {
	if id == 0 {
		return errors.New("ID no válido")
	}

	// Verificar que el estado existe
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("estado de equipo no encontrado")
	}

	// Verificar si hay equipos asociados a este estado
	equipos, err := s.repo.GetEquiposByEstado(id)
	if err != nil {
		return err
	}
	if len(equipos) > 0 {
		return errors.New("no se puede eliminar el estado porque hay equipos asociados a él")
	}

	return s.repo.Delete(id)
}

// ToggleActivo cambia el estado activo/inactivo de un estado de equipo
func (s *EstadoEquipoService) ToggleActivo(id uint) error {
	if id == 0 {
		return errors.New("ID no válido")
	}

	estado, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("estado de equipo no encontrado")
	}

	// Si se va a desactivar, verificar que no haya equipos activos con este estado
	if estado.Activo {
		equipos, err := s.repo.GetEquiposByEstado(id)
		if err != nil {
			return err
		}
		if len(equipos) > 0 {
			return errors.New("no se puede desactivar el estado porque hay equipos asociados a él")
		}
	}

	estado.Activo = !estado.Activo
	return s.repo.Update(estado)
}

// GetEquiposByEstado obtiene todos los equipos que tienen un estado específico
func (s *EstadoEquipoService) GetEquiposByEstado(estadoID uint) ([]models.Equipo, error) {
	if estadoID == 0 {
		return nil, errors.New("ID de estado no válido")
	}

	// Verificar que el estado existe
	_, err := s.repo.GetByID(estadoID)
	if err != nil {
		return nil, errors.New("estado de equipo no encontrado")
	}

	return s.repo.GetEquiposByEstado(estadoID)
}

// validateEstado valida los datos del estado de equipo
func (s *EstadoEquipoService) validateEstado(estado *models.EstadoEquipo) error {
	if strings.TrimSpace(estado.Nombre) == "" {
		return errors.New("el nombre del estado es obligatorio")
	}

	if len(strings.TrimSpace(estado.Nombre)) < 3 {
		return errors.New("el nombre del estado debe tener al menos 3 caracteres")
	}

	if len(strings.TrimSpace(estado.Nombre)) > 50 {
		return errors.New("el nombre del estado no puede tener más de 50 caracteres")
	}

	if strings.TrimSpace(estado.Descripcion) == "" {
		return errors.New("la descripción del estado es obligatoria")
	}

	if len(strings.TrimSpace(estado.Descripcion)) < 5 {
		return errors.New("la descripción del estado debe tener al menos 5 caracteres")
	}

	if len(strings.TrimSpace(estado.Descripcion)) > 255 {
		return errors.New("la descripción del estado no puede tener más de 255 caracteres")
	}

	return nil
}