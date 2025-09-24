package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/repositories"
)

// BackupService define las operaciones del servicio para Backup
type BackupService interface {
	CreateBackup(backup *models.Backup) error
	GetBackupByID(id uint) (*models.Backup, error)
	UpdateBackup(backup *models.Backup) error
	DeleteBackup(id uint) error
	GetAllBackups() ([]models.Backup, error)
	GetBackupsByEquipoID(equipoID uint) ([]models.Backup, error)
}

// backupService implementa BackupService
type backupService struct {
	backupRepo repositories.BackupRepository
}

// NewBackupService crea una nueva instancia de BackupService
func NewBackupService(backupRepo repositories.BackupRepository) BackupService {
	return &backupService{backupRepo: backupRepo}
}

// CreateBackup crea un nuevo backup
func (s *backupService) CreateBackup(backup *models.Backup) error {
	if backup.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if backup.RutaBackup == "" {
		return errors.New("la ruta del backup es obligatoria")
	}

	return s.backupRepo.Create(backup)
}

// GetBackupByID obtiene un backup por su ID
func (s *backupService) GetBackupByID(id uint) (*models.Backup, error) {
	return s.backupRepo.FindByID(id)
}

// UpdateBackup actualiza un backup existente
func (s *backupService) UpdateBackup(backup *models.Backup) error {
	if backup.ID == 0 {
		return errors.New("ID de backup no válido")
	}
	if backup.EquipoID == 0 {
		return errors.New("el ID del equipo es obligatorio")
	}
	if backup.RutaBackup == "" {
		return errors.New("la ruta del backup es obligatoria")
	}

	// Verificar si existe el backup
	existente, err := s.backupRepo.FindByID(backup.ID)
	if err != nil && existente != nil {
		return errors.New("backup no encontrado")
	}

	return s.backupRepo.Update(backup)
}

// DeleteBackup elimina un backup por su ID
func (s *backupService) DeleteBackup(id uint) error {
	if id == 0 {
		return errors.New("ID de backup no válido")
	}
	return s.backupRepo.Delete(id)
}

// GetAllBackups obtiene todos los backups
func (s *backupService) GetAllBackups() ([]models.Backup, error) {
	return s.backupRepo.FindAll()
}

// GetBackupsByEquipoID obtiene todos los backups asociados a un equipo
func (s *backupService) GetBackupsByEquipoID(equipoID uint) ([]models.Backup, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.backupRepo.FindByEquipoID(equipoID)
}
