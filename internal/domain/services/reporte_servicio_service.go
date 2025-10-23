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
	CrearReporteConTipo(reporteData *dto.CrearReporteCompletoDTO) (*models.ReporteServicio, error)
}

// reporteServicioService implementa ReporteServicioService
type reporteServicioService struct {
	reporteRepo     repositories.ReporteServicioRepository
	funcionarioRepo repositories.FuncionarioRepository
}

// NewReporteServicioService crea una nueva instancia de ReporteServicioService
func NewReporteServicioService(
	reporteRepo repositories.ReporteServicioRepository,
	funcionarioRepo repositories.FuncionarioRepository,
) ReporteServicioService {
	return &reporteServicioService{
		reporteRepo:     reporteRepo,
		funcionarioRepo: funcionarioRepo,
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

// CrearReporteConTipo crea un reporte de servicio completo con tipos de mantenimiento, repuestos y funcionarios
func (s *reporteServicioService) CrearReporteConTipo(reporteData *dto.CrearReporteCompletoDTO) (*models.ReporteServicio, error) {
	// Validaciones iniciales
	if reporteData == nil {
		return nil, errors.New("los datos del reporte son obligatorios")
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
	if reporteData.TipoMantenimiento.Tipo == "" {
		return nil, errors.New("debe especificar el tipo de mantenimiento")
	}
	if len(reporteData.FuncionarioIDs) == 0 {
		return nil, errors.New("debe especificar al menos un funcionario")
	}

	// Validar que el tipo de mantenimiento sea válido
	if reporteData.TipoMantenimiento.Tipo != "PREVENTIVO" && reporteData.TipoMantenimiento.Tipo != "CORRECTIVO" {
		return nil, errors.New("el tipo de mantenimiento debe ser 'PREVENTIVO' o 'CORRECTIVO'")
	}

	// Verificar que todos los funcionarios existen y obtenerlos
	var funcionarios []models.Funcionario
	for _, funcionarioID := range reporteData.FuncionarioIDs {
		funcionario, err := s.funcionarioRepo.FindByID(funcionarioID)
		if err != nil {
			return nil, errors.New("funcionario con ID " + string(rune(funcionarioID)) + " no encontrado")
		}
		funcionarios = append(funcionarios, *funcionario)
	}

	// Preparar datos para el repositorio
	reporte := reporteData.ToReporteServicio()
	tipoMantenimiento := reporteData.ToTipoMantenimiento(0) // El ID se asignará en el repositorio
	repuestos := reporteData.ToRepuestos(0)                 // El ID se asignará en el repositorio

	// Crear el reporte completo usando el repositorio
	if err := s.reporteRepo.CreateReporteCompleto(reporte, &tipoMantenimiento, repuestos, funcionarios); err != nil {
		return nil, errors.New("error al crear el reporte completo: " + err.Error())
	}

	// Cargar el reporte completo con todas sus relaciones
	reporteCompleto, err := s.reporteRepo.FindByID(reporte.ID)
	if err != nil {
		return nil, errors.New("error al cargar el reporte completo: " + err.Error())
	}

	return reporteCompleto, nil
}
