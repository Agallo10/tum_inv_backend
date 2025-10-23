package dto

import (
	"time"
	"tum_inv_backend/internal/domain/models"
)

// CrearReporteCompletoDTO representa la estructura para crear un reporte completo con todas sus relaciones
type CrearReporteCompletoDTO struct {
	// Datos del reporte principal
	EquipoID           *uint      `json:"equipo_id,omitempty"`
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

	// IDs de los funcionarios involucrados
	FuncionarioIDs []uint `json:"funcionario_ids" validate:"required,min=1"`
}

// TipoMantenimientoDTO representa los datos para crear un tipo de mantenimiento
type TipoMantenimientoDTO struct {
	Tipo            string `json:"tipo" validate:"required,oneof=PREVENTIVO CORRECTIVO"`
	Revision        bool   `json:"revision"`
	Instalacion     bool   `json:"instalacion"`
	Configuracion   bool   `json:"configuracion"`
	Ingreso         bool   `json:"ingreso"`
	Salida          bool   `json:"salida"`
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
