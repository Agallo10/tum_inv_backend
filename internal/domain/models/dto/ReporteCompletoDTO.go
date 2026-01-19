package dto

import (
	"strconv"
	"time"
	"tum_inv_backend/internal/domain/models"
)

// CrearReporteCompletoDTO representa la estructura para crear un reporte completo con todas sus relaciones
type CrearReporteCompletoDTO struct {
	// Datos del reporte principal
	CreadoPorID        uint       `json:"creado_por_id" validate:"required"`
	EquipoID           uint       `json:"equipo_id" validate:"required"`
	FechaInicio        time.Time  `json:"fecha_inicio" validate:"required"`
	FechaFinalizacion  *time.Time `json:"fecha_finalizacion,omitempty"`
	Dependencia        string     `json:"dependencia" validate:"required"`
	Ubicacion          string     `json:"ubicacion" validate:"required"`
	DiagnosticoFalla   string     `json:"diagnostico_falla,omitempty"`
	ActividadRealizada string     `json:"actividad_realizada" validate:"required"`
	Observaciones      string     `json:"observaciones,omitempty"`

	// Datos del tipo de mantenimiento
	TipoMantenimiento TipoMantenimientoDTO `json:"tipo_mantenimiento" validate:"required"`

	// Datos de los repuestos utilizados
	Repuestos []RepuestoDTO `json:"repuestos,omitempty"`
}

// TipoMantenimientoDTO representa los datos para crear un tipo de mantenimiento
type TipoMantenimientoDTO struct {
	Tipo            string `json:"tipo" validate:"required,oneof=PREVENTIVO CORRECTIVO"`
	Revision        bool   `json:"revision"`
	Instalacion     bool   `json:"instalacion"`
	Configuracion   bool   `json:"configuracion"`
	Ingreso         bool   `json:"ingreso"`
	Salida          bool   `json:"salida"`
	ConceptoBaja    bool   `json:"concepto_baja"`
	Otro            bool   `json:"otro"`
	DescripcionOtro string `json:"descripcion_otro,omitempty"`
}

// RepuestoDTO representa los datos para crear un repuesto
type RepuestoDTO struct {
	Cantidad          int       `json:"cantidad" validate:"required,min=1"`
	SerialNumeroParte string    `json:"serial_numero_parte" validate:"required"`
	Marca             string    `json:"marca,omitempty"`
	Tecnologia        string    `json:"tecnologia,omitempty"`
	Capacidad         string    `json:"capacidad,omitempty"`
	Descripcion       string    `json:"descripcion" validate:"required"`
	FechaUtilizacion  time.Time `json:"fecha_utilizacion"`
}

// ToReporteServicio convierte el DTO a modelo ReporteServicio
func (dto *CrearReporteCompletoDTO) ToReporteServicio() *models.ReporteServicio {
	return &models.ReporteServicio{
		CreadoPorID:        dto.CreadoPorID,
		EquipoID:           dto.EquipoID,
		FechaInicio:        dto.FechaInicio,
		FechaFinalizacion:  dto.FechaFinalizacion,
		Dependencia:        dto.Dependencia,
		Ubicacion:          dto.Ubicacion,
		DiagnosticoFalla:   dto.DiagnosticoFalla,
		ActividadRealizada: dto.ActividadRealizada,
		Observaciones:      dto.Observaciones,
	}
}

// ToTipoMantenimiento convierte el DTO a modelo TipoMantenimiento
func (dto *CrearReporteCompletoDTO) ToTipoMantenimiento(reporteID uint) models.TipoMantenimiento {
	return models.TipoMantenimiento{
		ReporteID:       reporteID,
		Tipo:            dto.TipoMantenimiento.Tipo,
		Revision:        dto.TipoMantenimiento.Revision,
		Instalacion:     dto.TipoMantenimiento.Instalacion,
		Configuracion:   dto.TipoMantenimiento.Configuracion,
		Ingreso:         dto.TipoMantenimiento.Ingreso,
		Salida:          dto.TipoMantenimiento.Salida,
		ConceptoBaja:    dto.TipoMantenimiento.ConceptoBaja,
		Otro:            dto.TipoMantenimiento.Otro,
		DescripcionOtro: dto.TipoMantenimiento.DescripcionOtro,
	}
}

// ToRepuestos convierte los DTOs a modelos Repuesto
func (dto *CrearReporteCompletoDTO) ToRepuestos(reporteID uint) []models.Repuesto {
	repuestos := make([]models.Repuesto, len(dto.Repuestos))
	for i, repuestoDTO := range dto.Repuestos {
		repuestosReporteID := reporteID
		repuestos[i] = models.Repuesto{
			ReporteID:         &repuestosReporteID,
			Cantidad:          repuestoDTO.Cantidad,
			SerialNumeroParte: repuestoDTO.SerialNumeroParte,
			Marca:             repuestoDTO.Marca,
			Tecnologia:        repuestoDTO.Tecnologia,
			Capacidad:         repuestoDTO.Capacidad,
			Descripcion:       repuestoDTO.Descripcion,
			FechaUtilizacion:  repuestoDTO.FechaUtilizacion,
		}
	}
	return repuestos
}

// ReporteResumenDTO representa un resumen de reporte para listados
type ReporteResumenDTO struct {
	ID                 uint    `json:"id"`
	CreadoPorID        uint    `json:"creado_por_id"`
	CreadoPorNombre    string  `json:"creado_por_nombre"`
	DiagnosticoFalla   string  `json:"diagnostico_falla"`
	ActividadRealizada string  `json:"actividad_realizada"`
	TipoMantenimiento  string  `json:"tipo_mantenimiento"`
	Repuestos          string  `json:"repuestos"`
	FechaInicio        string  `json:"fecha_inicio"`
	FechaFinalizacion  *string `json:"fecha_finalizacion"`
}

// ReportesToResumenDTO convierte una lista de reportes a DTOs de resumen
func ReportesToResumenDTO(reportes []models.ReporteServicio) []ReporteResumenDTO {
	resumen := make([]ReporteResumenDTO, len(reportes))
	for i, reporte := range reportes {
		// Determinar texto de repuestos
		repuestosTexto := "No"
		if len(reporte.Repuestos) > 0 {
			repuestosTexto = "Sí (" + strconv.Itoa(len(reporte.Repuestos)) + ")"
		}

		// Formato de fecha
		fechaInicio := reporte.FechaInicio.Format("2006-01-02 15:04")
		var fechaFinalizacion *string
		if reporte.FechaFinalizacion != nil {
			formatted := reporte.FechaFinalizacion.Format("2006-01-02 15:04")
			fechaFinalizacion = &formatted
		}

		// Nombre del creador
		creadoPorNombre := reporte.CreadoPor.Nombre + " " + reporte.CreadoPor.Apellido

		resumen[i] = ReporteResumenDTO{
			ID:                 reporte.ID,
			CreadoPorID:        reporte.CreadoPorID,
			CreadoPorNombre:    creadoPorNombre,
			DiagnosticoFalla:   reporte.DiagnosticoFalla,
			ActividadRealizada: reporte.ActividadRealizada,
			TipoMantenimiento:  reporte.TipoMantenimiento.Tipo,
			Repuestos:          repuestosTexto,
			FechaInicio:        fechaInicio,
			FechaFinalizacion:  fechaFinalizacion,
		}
	}
	return resumen
}
