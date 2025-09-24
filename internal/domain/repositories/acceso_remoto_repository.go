package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// AccesoRemotoRepository define las operaciones del repositorio para AccesoRemoto
type AccesoRemotoRepository interface {
	Create(acceso *models.AccesoRemoto) error
	FindByID(id uint) (*models.AccesoRemoto, error)
	Update(acceso *models.AccesoRemoto) error
	Delete(id uint) error
	FindAll() ([]models.AccesoRemoto, error)
	FindByEquipoID(equipoID uint) ([]models.AccesoRemoto, error)
}

// accesoRemotoRepository implementa AccesoRemotoRepository
type accesoRemotoRepository struct {
	db *gorm.DB
}

// NewAccesoRemotoRepository crea una nueva instancia de AccesoRemotoRepository
func NewAccesoRemotoRepository(db *gorm.DB) AccesoRemotoRepository {
	return &accesoRemotoRepository{db: db}
}

// Create crea un nuevo acceso remoto en la base de datos
func (r *accesoRemotoRepository) Create(acceso *models.AccesoRemoto) error {
	return r.db.Create(acceso).Error
}

// FindByID busca un acceso remoto por su ID
func (r *accesoRemotoRepository) FindByID(id uint) (*models.AccesoRemoto, error) {
	var acceso models.AccesoRemoto
	err := r.db.First(&acceso, id).Error
	if err != nil {
		return nil, err
	}
	return &acceso, nil
}

// Update actualiza un acceso remoto existente
func (r *accesoRemotoRepository) Update(acceso *models.AccesoRemoto) error {
	return r.db.Save(acceso).Error
}

// Delete elimina un acceso remoto por su ID
func (r *accesoRemotoRepository) Delete(id uint) error {
	return r.db.Delete(&models.AccesoRemoto{}, id).Error
}

// FindAll retorna todos los accesos remotos
func (r *accesoRemotoRepository) FindAll() ([]models.AccesoRemoto, error) {
	var accesos []models.AccesoRemoto
	err := r.db.Find(&accesos).Error
	return accesos, err
}

// FindByEquipoID retorna todos los accesos remotos asociados a un equipo
func (r *accesoRemotoRepository) FindByEquipoID(equipoID uint) ([]models.AccesoRemoto, error) {
	var accesos []models.AccesoRemoto
	err := r.db.Where("equipo_id = ?", equipoID).Find(&accesos).Error
	return accesos, err
}