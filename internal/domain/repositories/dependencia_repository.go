package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// DependenciaRepository define las operaciones del repositorio para Dependencia
type DependenciaRepository interface {
	CreateDependencia(dependencia *models.Dependencia) error
	GetDependenciaByID(id uint) (*models.Dependencia, error)
	GetAllDependencias() ([]models.Dependencia, error)
	UpdateDependencia(dependencia *models.Dependencia) error
	DeleteDependencia(id uint) error
	GetDependenciasBySecretariaID(secretariaID uint) ([]models.Dependencia, error)
}

// dependenciaRepository implementa DependenciaRepository
type dependenciaRepository struct {
	db *gorm.DB
}

// NewDependenciaRepository crea una nueva instancia de DependenciaRepository
func NewDependenciaRepository(db *gorm.DB) DependenciaRepository {
	return &dependenciaRepository{db: db}
}

// CreateDependencia crea una nueva dependencia en la base de datos
func (r *dependenciaRepository) CreateDependencia(dependencia *models.Dependencia) error {
	return r.db.Create(dependencia).Error
}

// GetDependenciaByID busca una dependencia por su ID
func (r *dependenciaRepository) GetDependenciaByID(id uint) (*models.Dependencia, error) {
	var dependencia models.Dependencia
	err := r.db.Preload("UsuarioResponsables").Preload("Equipos").First(&dependencia, id).Error
	if err != nil {
		return nil, err
	}
	return &dependencia, nil
}

// GetAllDependencias retorna todas las dependencias
func (r *dependenciaRepository) GetAllDependencias() ([]models.Dependencia, error) {
	var dependencias []models.Dependencia
	err := r.db.Find(&dependencias).Error
	return dependencias, err
}

// UpdateDependencia actualiza una dependencia existente
func (r *dependenciaRepository) UpdateDependencia(dependencia *models.Dependencia) error {
	return r.db.Save(dependencia).Error
}

// DeleteDependencia elimina una dependencia por su ID
func (r *dependenciaRepository) DeleteDependencia(id uint) error {
	return r.db.Delete(&models.Dependencia{}, id).Error
}

// GetDependenciasBySecretariaID retorna todas las dependencias de una secretaría
func (r *dependenciaRepository) GetDependenciasBySecretariaID(secretariaID uint) ([]models.Dependencia, error) {
	var dependencias []models.Dependencia
	err := r.db.Where("secretaria_id = ?", secretariaID).Find(&dependencias).Error
	return dependencias, err
}