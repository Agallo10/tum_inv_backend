package services

import (
	"errors"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/models/dto"
	"tum_inv_backend/internal/domain/repositories"
)

// ReporteServicioService define las operaciones del servicio para ReporteServicio
type ReporteServicioService interface {
	CreateReporteServicio(reporte *models.ReporteServicio) error
	GetReporteServicioByID(id uint) (*models.ReporteServicio, error)
	UpdateReporteServicio(reporte *models.ReporteServicio) error
	DeleteReporteServicio(id uint) error
	GetAllReportesServicio() ([]models.ReporteServicio, error)
	GetReportesServicioByEquipoID(equipoID uint) ([]models.ReporteServicio, error)
	GetReportesResumenByEquipoID(equipoID uint) ([]dto.ReporteResumenDTO, error)
	CrearReporteConTipo(reporteData *dto.CrearReporteCompletoDTO) (*models.ReporteServicio, error)
}

// reporteServicioService implementa ReporteServicioService
type reporteServicioService struct {
	reporteRepo repositories.ReporteServicioRepository
}

// NewReporteServicioService crea una nueva instancia de ReporteServicioService
func NewReporteServicioService(
	reporteRepo repositories.ReporteServicioRepository,
) ReporteServicioService {
	return &reporteServicioService{
		reporteRepo: reporteRepo,
	}
}

// CreateReporteServicio crea un nuevo reporte de servicio
func (s *reporteServicioService) CreateReporteServicio(reporte *models.ReporteServicio) error {
	if reporte.Dependencia == "" {
		return errors.New("la dependencia es obligatoria")
	}
	if reporte.Ubicacion == "" {
		return errors.New("la ubicación es obligatoria")
	}
	if reporte.ActividadRealizada == "" {
		return errors.New("la actividad realizada es obligatoria")
	}

	return s.reporteRepo.Create(reporte)
}

// GetReporteServicioByID obtiene un reporte de servicio por su ID
func (s *reporteServicioService) GetReporteServicioByID(id uint) (*models.ReporteServicio, error) {
	return s.reporteRepo.FindByID(id)
}

// UpdateReporteServicio actualiza un reporte de servicio existente
func (s *reporteServicioService) UpdateReporteServicio(reporte *models.ReporteServicio) error {
	if reporte.ID == 0 {
		return errors.New("ID de reporte no válido")
	}
	if reporte.Dependencia == "" {
		return errors.New("la dependencia es obligatoria")
	}
	if reporte.Ubicacion == "" {
		return errors.New("la ubicación es obligatoria")
	}
	if reporte.ActividadRealizada == "" {
		return errors.New("la actividad realizada es obligatoria")
	}

	// Verificar si existe el reporte
	existente, err := s.reporteRepo.FindByID(reporte.ID)
	if err != nil && existente != nil {
		return errors.New("reporte no encontrado")
	}

	return s.reporteRepo.Update(reporte)
}

// DeleteReporteServicio elimina un reporte de servicio por su ID
func (s *reporteServicioService) DeleteReporteServicio(id uint) error {
	if id == 0 {
		return errors.New("ID de reporte no válido")
	}
	return s.reporteRepo.Delete(id)
}

// GetAllReportesServicio obtiene todos los reportes de servicio
func (s *reporteServicioService) GetAllReportesServicio() ([]models.ReporteServicio, error) {
	return s.reporteRepo.FindAll()
}

// GetReportesServicioByEquipoID obtiene todos los reportes de servicio asociados a un equipo
func (s *reporteServicioService) GetReportesServicioByEquipoID(equipoID uint) ([]models.ReporteServicio, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	return s.reporteRepo.FindByEquipoID(equipoID)
}

// GetReportesResumenByEquipoID obtiene un resumen de los reportes de servicio de un equipo
func (s *reporteServicioService) GetReportesResumenByEquipoID(equipoID uint) ([]dto.ReporteResumenDTO, error) {
	if equipoID == 0 {
		return nil, errors.New("ID de equipo no válido")
	}
	reportes, err := s.reporteRepo.FindByEquipoID(equipoID)
	if err != nil {
		return nil, err
	}
	return dto.ReportesToResumenDTO(reportes), nil
}

// CrearReporteConTipo crea un reporte de servicio completo con tipos de mantenimiento y repuestos
func (s *reporteServicioService) CrearReporteConTipo(reporteData *dto.CrearReporteCompletoDTO) (*models.ReporteServicio, error) {
	// Validaciones iniciales
	if reporteData == nil {
		return nil, errors.New("los datos del reporte son obligatorios")
	}
	if reporteData.EquipoID == 0 {
		return nil, errors.New("el equipo es obligatorio")
	}
	if reporteData.Dependencia == "" {
		return nil, errors.New("la dependencia es obligatoria")
	}
	if reporteData.Ubicacion == "" {
		return nil, errors.New("la ubicación es obligatoria")
	}
	if reporteData.ActividadRealizada == "" {
		return nil, errors.New("la actividad realizada es obligatoria")
	}

	// Validar tipo de mantenimiento: debe ser PREVENTIVO, CORRECTIVO, o vacío si Otro es true
	tipoValido := reporteData.TipoMantenimiento.Tipo == "PREVENTIVO" || reporteData.TipoMantenimiento.Tipo == "CORRECTIVO"
	otroValido := reporteData.TipoMantenimiento.Tipo == "" && reporteData.TipoMantenimiento.Otro

	if !tipoValido && !otroValido {
		return nil, errors.New("debe especificar el tipo de mantenimiento (PREVENTIVO, CORRECTIVO) o marcar 'otro'")
	}

	// Preparar datos para el repositorio
	reporte := reporteData.ToReporteServicio()
	tipoMantenimiento := reporteData.ToTipoMantenimiento(0) // El ID se asignará en el repositorio
	repuestos := reporteData.ToRepuestos(0)                 // El ID se asignará en el repositorio

	// Crear el reporte completo usando el repositorio
	if err := s.reporteRepo.CreateReporteCompleto(reporte, &tipoMantenimiento, repuestos); err != nil {
		return nil, errors.New("error al crear el reporte completo: " + err.Error())
	}

	// Cargar el reporte completo con todas sus relaciones
	reporteCompleto, err := s.reporteRepo.FindByID(reporte.ID)
	if err != nil {
		return nil, errors.New("error al cargar el reporte completo: " + err.Error())
	}

	return reporteCompleto, nil
}
