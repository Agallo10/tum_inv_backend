package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// FuncionarioService define las operaciones del servicio para Funcionario
type FuncionarioService interface {
	CreateFuncionario(funcionario *models.Funcionario) error
	GetFuncionarioByID(id uint) (*models.Funcionario, error)
	UpdateFuncionario(funcionario *models.Funcionario) error
	DeleteFuncionario(id uint) error
	GetAllFuncionarios() ([]models.Funcionario, error)
	GetFuncionarioByCedula(cedula string) (*models.Funcionario, error)
	GetFuncionariosByReporteID(reporteID uint) ([]models.Funcionario, error)
}

// funcionarioService implementa FuncionarioService
type funcionarioService struct {
	funcionarioRepo repositories.FuncionarioRepository
}

// NewFuncionarioService crea una nueva instancia de FuncionarioService
func NewFuncionarioService(funcionarioRepo repositories.FuncionarioRepository) FuncionarioService {
	return &funcionarioService{funcionarioRepo: funcionarioRepo}
}

// CreateFuncionario crea un nuevo funcionario
func (s *funcionarioService) CreateFuncionario(funcionario *models.Funcionario) error {
	if funcionario.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if funcionario.Cargo == "" {
		return errors.New("el cargo es obligatorio")
	}
	if funcionario.Cedula == "" {
		return errors.New("la cédula es obligatoria")
	}
	if funcionario.Tipo == "" {
		return errors.New("el tipo es obligatorio")
	}
	if funcionario.Area == "" {
		return errors.New("el área es obligatoria")
	}
	
	// Verificar si ya existe un funcionario con la misma cédula
	existente, err := s.funcionarioRepo.FindByCedula(funcionario.Cedula)
	if err == nil && existente != nil {
		return errors.New("ya existe un funcionario con esta cédula")
	}
	
	return s.funcionarioRepo.Create(funcionario)
}

// GetFuncionarioByID obtiene un funcionario por su ID
func (s *funcionarioService) GetFuncionarioByID(id uint) (*models.Funcionario, error) {
	return s.funcionarioRepo.FindByID(id)
}

// UpdateFuncionario actualiza un funcionario existente
func (s *funcionarioService) UpdateFuncionario(funcionario *models.Funcionario) error {
	if funcionario.ID == 0 {
		return errors.New("ID de funcionario no válido")
	}
	if funcionario.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if funcionario.Cargo == "" {
		return errors.New("el cargo es obligatorio")
	}
	if funcionario.Cedula == "" {
		return errors.New("la cédula es obligatoria")
	}
	if funcionario.Tipo == "" {
		return errors.New("el tipo es obligatorio")
	}
	if funcionario.Area == "" {
		return errors.New("el área es obligatoria")
	}
	
	// Verificar si existe el funcionario
	existente, err := s.funcionarioRepo.FindByID(funcionario.ID)
	if err != nil {
		return errors.New("funcionario no encontrado")
	}
	
	// Verificar si al cambiar la cédula no se genera conflicto con otro funcionario
	if existente.Cedula != funcionario.Cedula {
		otro, err := s.funcionarioRepo.FindByCedula(funcionario.Cedula)
		if err == nil && otro != nil && otro.ID != funcionario.ID {
			return errors.New("ya existe un funcionario con esta cédula")
		}
	}
	
	return s.funcionarioRepo.Update(funcionario)
}

// DeleteFuncionario elimina un funcionario por su ID
func (s *funcionarioService) DeleteFuncionario(id uint) error {
	if id == 0 {
		return errors.New("ID de funcionario no válido")
	}
	return s.funcionarioRepo.Delete(id)
}

// GetAllFuncionarios obtiene todos los funcionarios
func (s *funcionarioService) GetAllFuncionarios() ([]models.Funcionario, error) {
	return s.funcionarioRepo.FindAll()
}

// GetFuncionarioByCedula obtiene un funcionario por su cédula
func (s *funcionarioService) GetFuncionarioByCedula(cedula string) (*models.Funcionario, error) {
	if cedula == "" {
		return nil, errors.New("la cédula es obligatoria")
	}
	return s.funcionarioRepo.FindByCedula(cedula)
}

// GetFuncionariosByReporteID obtiene todos los funcionarios asociados a un reporte
func (s *funcionarioService) GetFuncionariosByReporteID(reporteID uint) ([]models.Funcionario, error) {
	if reporteID == 0 {
		return nil, errors.New("ID de reporte no válido")
	}
	return s.funcionarioRepo.FindByReporteID(reporteID)
}