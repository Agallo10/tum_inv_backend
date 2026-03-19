package services

import (
	"errors"
	"fmt"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/models/dto"
	"tum_inv_backend/internal/domain/repositories"
	"tum_inv_backend/internal/infrastructure/storage"
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
	SubirFirmado(reporteID uint, fileData []byte, contentType string) (*models.ReporteServicio, error)
	ObtenerURLFirmado(reporteID uint) (string, error)
	ReabrirReporte(reporteID uint) error
}

// reporteServicioService implementa ReporteServicioService
type reporteServicioService struct {
	reporteRepo repositories.ReporteServicioRepository
	storage     *storage.SupabaseStorage
}

// NewReporteServicioService crea una nueva instancia de ReporteServicioService
func NewReporteServicioService(
	reporteRepo repositories.ReporteServicioRepository,
	storageSvc ...*storage.SupabaseStorage,
) ReporteServicioService {
	s := &reporteServicioService{
		reporteRepo: reporteRepo,
	}
	if len(storageSvc) > 0 {
		s.storage = storageSvc[0]
	}
	return s
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

// SubirFirmado sube un PDF firmado a Supabase Storage y cierra el reporte
func (s *reporteServicioService) SubirFirmado(reporteID uint, fileData []byte, contentType string) (*models.ReporteServicio, error) {
	if reporteID == 0 {
		return nil, errors.New("ID de reporte no válido")
	}
	if s.storage == nil {
		return nil, errors.New("servicio de almacenamiento no configurado")
	}

	// Verificar que el reporte existe y no está cerrado
	reporte, err := s.reporteRepo.FindByID(reporteID)
	if err != nil {
		return nil, errors.New("reporte no encontrado")
	}
	if reporte.FechaCierre != nil {
		return nil, errors.New("el reporte ya está cerrado")
	}

	// Subir archivo a Supabase Storage
	fileName := fmt.Sprintf("reporte_%d_firmado.pdf", reporteID)
	objectPath, err := s.storage.Upload(fileName, fileData, contentType)
	if err != nil {
		return nil, fmt.Errorf("error al subir archivo: %w", err)
	}

	// Cerrar el reporte en la BD
	if err := s.reporteRepo.CerrarReporte(reporteID, objectPath); err != nil {
		return nil, fmt.Errorf("error al cerrar el reporte: %w", err)
	}

	// Retornar el reporte actualizado
	return s.reporteRepo.FindByID(reporteID)
}

// ObtenerURLFirmado genera una URL firmada temporal para descargar el PDF firmado
func (s *reporteServicioService) ObtenerURLFirmado(reporteID uint) (string, error) {
	if reporteID == 0 {
		return "", errors.New("ID de reporte no válido")
	}
	if s.storage == nil {
		return "", errors.New("servicio de almacenamiento no configurado")
	}

	reporte, err := s.reporteRepo.FindByID(reporteID)
	if err != nil {
		return "", errors.New("reporte no encontrado")
	}
	if reporte.FechaCierre == nil || reporte.ArchivoFirmadoURL == "" {
		return "", errors.New("este reporte no tiene un PDF firmado")
	}

	// Generar URL firmada (válida por 1 hora = 3600 segundos)
	fileName := fmt.Sprintf("reporte_%d_firmado.pdf", reporteID)
	signedURL, err := s.storage.GetSignedURL(fileName, 3600)
	if err != nil {
		return "", fmt.Errorf("error al generar URL de descarga: %w", err)
	}

	return signedURL, nil
}

// ReabrirReporte elimina el PDF firmado del storage y reabre el reporte
func (s *reporteServicioService) ReabrirReporte(reporteID uint) error {
	if reporteID == 0 {
		return errors.New("ID de reporte no válido")
	}
	if s.storage == nil {
		return errors.New("servicio de almacenamiento no configurado")
	}

	// Verificar que el reporte existe y está cerrado
	reporte, err := s.reporteRepo.FindByID(reporteID)
	if err != nil {
		return errors.New("reporte no encontrado")
	}
	if reporte.FechaCierre == nil {
		return errors.New("el reporte no está cerrado")
	}

	// Eliminar el archivo del storage
	fileName := fmt.Sprintf("reporte_%d_firmado.pdf", reporteID)
	if err := s.storage.Delete(fileName); err != nil {
		return fmt.Errorf("error al eliminar el archivo: %w", err)
	}

	// Reabrir el reporte en la BD
	return s.reporteRepo.ReabrirReporte(reporteID)
}
