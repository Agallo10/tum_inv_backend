package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// BackupController maneja las solicitudes HTTP relacionadas con backups
type BackupController struct {
	backupService services.BackupService
}

// NewBackupController crea una nueva instancia de BackupController
func NewBackupController(backupService services.BackupService) *BackupController {
	return &BackupController{
		backupService: backupService,
	}
}

// CreateBackup maneja la creación de un nuevo backup
func (c *BackupController) CreateBackup(ctx echo.Context) error {
	backup := new(models.Backup)
	if err := ctx.Bind(backup); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.backupService.CreateBackup(backup); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, backup)
}

// GetBackup obtiene un backup por su ID
func (c *BackupController) GetBackup(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	backup, err := c.backupService.GetBackupByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Backup no encontrado"})
	}

	return ctx.JSON(http.StatusOK, backup)
}

// UpdateBackup actualiza un backup existente
func (c *BackupController) UpdateBackup(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	backup := new(models.Backup)
	if err := ctx.Bind(backup); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	backup.ID = uint(id)
	if err := c.backupService.UpdateBackup(backup); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, backup)
}

// DeleteBackup elimina un backup por su ID
func (c *BackupController) DeleteBackup(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.backupService.DeleteBackup(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Backup eliminado correctamente"})
}

// GetAllBackups obtiene todos los backups
func (c *BackupController) GetAllBackups(ctx echo.Context) error {
	backups, err := c.backupService.GetAllBackups()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, backups)
}

// GetBackupsByEquipo obtiene todos los backups asociados a un equipo
func (c *BackupController) GetBackupsByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	backups, err := c.backupService.GetBackupsByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, backups)
}