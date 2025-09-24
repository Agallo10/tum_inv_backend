package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// SecretariaRepository define las operaciones del repositorio para Secretaria
type SecretariaRepository interface {
	CreateSecretaria(secretaria *models.Secretaria) error
	GetSecretariaByID(id uint) (*models.Secretaria, error)
	GetAllSecretarias() ([]models.Secretaria, error)
	UpdateSecretaria(secretaria *models.Secretaria) error
	DeleteSecretaria(id uint) error
}

// secretariaRepository implementa SecretariaRepository
type secretariaRepository struct {
	db *gorm.DB
}

// NewSecretariaRepository crea una nueva instancia de SecretariaRepository
func NewSecretariaRepository(db *gorm.DB) SecretariaRepository {
	return &secretariaRepository{db: db}
}

// CreateSecretaria crea una nueva secretaría/ en la base de datos
func (r *secretariaRepository) CreateSecretaria(secretaria *models.Secretaria) error {
	return r.db.Create(secretaria).Error
}

// GetSecretariaByID busca una secretaría/ por su ID
func (r *secretariaRepository) GetSecretariaByID(id uint) (*models.Secretaria, error) {
	var secretaria models.Secretaria
	err := r.db.Preload("Dependencias").First(&secretaria, id).Error
	if err != nil {
		return nil, err
	}
	return &secretaria, nil
}

// GetAllSecretarias retorna todas las secretaríass
func (r *secretariaRepository) GetAllSecretarias() ([]models.Secretaria, error) {
	var secretarias []models.Secretaria
	err := r.db.Find(&secretarias).Error
	return secretarias, err
}

// UpdateSecretaria actualiza una secretaría existente
func (r *secretariaRepository) UpdateSecretaria(secretaria *models.Secretaria) error {
	return r.db.Save(secretaria).Error
}

// DeleteSecretaria elimina una secretaría por su ID
func (r *secretariaRepository) DeleteSecretaria(id uint) error {
	return r.db.Delete(&models.Secretaria{}, id).Error
}
