package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// FuncionarioRepository define las operaciones del repositorio para Funcionario
type FuncionarioRepository interface {
	Create(funcionario *models.Funcionario) error
	FindByID(id uint) (*models.Funcionario, error)
	Update(funcionario *models.Funcionario) error
	Delete(id uint) error
	FindAll() ([]models.Funcionario, error)
	FindByCedula(cedula string) (*models.Funcionario, error)
	FindByReporteID(reporteID uint) ([]models.Funcionario, error)
}

// funcionarioRepository implementa FuncionarioRepository
type funcionarioRepository struct {
	db *gorm.DB
}

// NewFuncionarioRepository crea una nueva instancia de FuncionarioRepository
func NewFuncionarioRepository(db *gorm.DB) FuncionarioRepository {
	return &funcionarioRepository{db: db}
}

// Create crea un nuevo funcionario en la base de datos
func (r *funcionarioRepository) Create(funcionario *models.Funcionario) error {
	return r.db.Create(funcionario).Error
}

// FindByID busca un funcionario por su ID
func (r *funcionarioRepository) FindByID(id uint) (*models.Funcionario, error) {
	var funcionario models.Funcionario
	err := r.db.Preload("Reportes").First(&funcionario, id).Error
	if err != nil {
		return nil, err
	}
	return &funcionario, nil
}

// Update actualiza un funcionario existente
func (r *funcionarioRepository) Update(funcionario *models.Funcionario) error {
	return r.db.Save(funcionario).Error
}

// Delete elimina un funcionario por su ID
func (r *funcionarioRepository) Delete(id uint) error {
	return r.db.Delete(&models.Funcionario{}, id).Error
}

// FindAll retorna todos los funcionarios
func (r *funcionarioRepository) FindAll() ([]models.Funcionario, error) {
	var funcionarios []models.Funcionario
	err := r.db.Find(&funcionarios).Error
	return funcionarios, err
}

// FindByCedula busca un funcionario por su cédula
func (r *funcionarioRepository) FindByCedula(cedula string) (*models.Funcionario, error) {
	var funcionario models.Funcionario
	err := r.db.Where("cedula = ?", cedula).First(&funcionario).Error
	if err != nil {
		return nil, err
	}
	return &funcionario, nil
}

// FindByReporteID retorna todos los funcionarios asociados a un reporte
func (r *funcionarioRepository) FindByReporteID(reporteID uint) ([]models.Funcionario, error) {
	var funcionarios []models.Funcionario
	err := r.db.Joins("JOIN reportes_funcionarios ON reportes_funcionarios.funcionario_id = funcionarios.id").Where("reportes_funcionarios.reporte_servicio_id = ?", reporteID).Find(&funcionarios).Error
	return funcionarios, err
}