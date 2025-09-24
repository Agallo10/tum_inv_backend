package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// BackupRepository define las operaciones del repositorio para Backup
type BackupRepository interface {
	Create(backup *models.Backup) error
	FindByID(id uint) (*models.Backup, error)
	Update(backup *models.Backup) error
	Delete(id uint) error
	FindAll() ([]models.Backup, error)
	FindByEquipoID(equipoID uint) ([]models.Backup, error)
}

// backupRepository implementa BackupRepository
type backupRepository struct {
	db *gorm.DB
}

// NewBackupRepository crea una nueva instancia de BackupRepository
func NewBackupRepository(db *gorm.DB) BackupRepository {
	return &backupRepository{db: db}
}

// Create crea un nuevo backup en la base de datos
func (r *backupRepository) Create(backup *models.Backup) error {
	return r.db.Create(backup).Error
}

// FindByID busca un backup por su ID
func (r *backupRepository) FindByID(id uint) (*models.Backup, error) {
	var backup models.Backup
	err := r.db.First(&backup, id).Error
	if err != nil {
		return nil, err
	}
	return &backup, nil
}

// Update actualiza un backup existente
func (r *backupRepository) Update(backup *models.Backup) error {
	return r.db.Save(backup).Error
}

// Delete elimina un backup por su ID
func (r *backupRepository) Delete(id uint) error {
	return r.db.Delete(&models.Backup{}, id).Error
}

// FindAll retorna todos los backups
func (r *backupRepository) FindAll() ([]models.Backup, error) {
	var backups []models.Backup
	err := r.db.Find(&backups).Error
	return backups, err
}

// FindByEquipoID retorna todos los backups asociados a un equipo
func (r *backupRepository) FindByEquipoID(equipoID uint) ([]models.Backup, error) {
	var backups []models.Backup
	err := r.db.Where("equipo_id = ?", equipoID).Find(&backups).Error
	return backups, err
}