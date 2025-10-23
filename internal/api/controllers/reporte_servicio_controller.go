package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/models/dto"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// ReporteServicioController maneja las solicitudes HTTP relacionadas con reportes de servicio
type ReporteServicioController struct {
	reporteService services.ReporteServicioService
}

// NewReporteServicioController crea una nueva instancia de ReporteServicioController
func NewReporteServicioController(reporteService services.ReporteServicioService) *ReporteServicioController {
	return &ReporteServicioController{
		reporteService: reporteService,
	}
}

// CreateReporteServicio maneja la creación de un nuevo reporte de servicio
func (c *ReporteServicioController) CreateReporteServicio(ctx echo.Context) error {
	reporte := new(models.ReporteServicio)
	if err := ctx.Bind(reporte); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.reporteService.CreateReporteServicio(reporte); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, reporte)
}

// GetReporteServicio obtiene un reporte de servicio por su ID
func (c *ReporteServicioController) GetReporteServicio(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	reporte, err := c.reporteService.GetReporteServicioByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Reporte no encontrado"})
	}

	return ctx.JSON(http.StatusOK, reporte)
}

// UpdateReporteServicio actualiza un reporte de servicio existente
func (c *ReporteServicioController) UpdateReporteServicio(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	reporte := new(models.ReporteServicio)
	if err := ctx.Bind(reporte); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	reporte.ID = uint(id)
	if err := c.reporteService.UpdateReporteServicio(reporte); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reporte)
}

// DeleteReporteServicio elimina un reporte de servicio por su ID
func (c *ReporteServicioController) DeleteReporteServicio(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.reporteService.DeleteReporteServicio(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Reporte eliminado correctamente"})
}

// GetAllReportesServicio obtiene todos los reportes de servicio
func (c *ReporteServicioController) GetAllReportesServicio(ctx echo.Context) error {
	reportes, err := c.reporteService.GetAllReportesServicio()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reportes)
}

// GetReportesServicioByEquipo obtiene todos los reportes de servicio asociados a un equipo
func (c *ReporteServicioController) GetReportesServicioByEquipo(ctx echo.Context) error {
	equipoID, err := strconv.ParseUint(ctx.Param("equipoId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de equipo inválido"})
	}

	reportes, err := c.reporteService.GetReportesServicioByEquipoID(uint(equipoID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, reportes)
}

// CrearReporteConTipo crea un reporte de servicio completo con tipos de mantenimiento, repuestos y funcionarios
func (c *ReporteServicioController) CrearReporteConTipo(ctx echo.Context) error {
	reporteData := new(dto.CrearReporteCompletoDTO)
	if err := ctx.Bind(reporteData); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	// Validar que los campos requeridos estén presentes
	if reporteData.Dependencia == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "La dependencia es obligatoria",
		})
	}
	if reporteData.Ubicacion == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "La ubicación es obligatoria",
		})
	}
	if reporteData.ActividadRealizada == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "La actividad realizada es obligatoria",
		})
	}
	if len(reporteData.TipoMantenimiento.Tipo) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Debe especificar el tipo de mantenimiento",
		})
	}
	if len(reporteData.FuncionarioIDs) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Debe especificar al menos un funcionario",
		})
	}

	// Validar tipo de mantenimiento
	if reporteData.TipoMantenimiento.Tipo != "PREVENTIVO" && reporteData.TipoMantenimiento.Tipo != "CORRECTIVO" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "El tipo de mantenimiento debe ser 'PREVENTIVO' o 'CORRECTIVO'",
		})
	}

	// Validar repuestos (si existen)
	for i, repuesto := range reporteData.Repuestos {
		if repuesto.Cantidad <= 0 {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "La cantidad del repuesto en la posición " + strconv.Itoa(i) + " debe ser mayor a 0",
			})
		}
		if repuesto.SerialNumeroParte == "" {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "El serial/número de parte del repuesto en la posición " + strconv.Itoa(i) + " es obligatorio",
			})
		}
		if repuesto.Descripcion == "" {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "La descripción del repuesto en la posición " + strconv.Itoa(i) + " es obligatoria",
			})
		}
	}

	// Crear el reporte completo
	reporte, err := c.reporteService.CrearReporteConTipo(reporteData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Reporte creado exitosamente",
		"reporte": reporte,
	})
}
