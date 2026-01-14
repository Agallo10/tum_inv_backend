package services

import (
	"bytes"
	"fmt"
	"strings"
	"tum_inv_backend/internal/domain/models"

	"github.com/go-pdf/fpdf"
	"gorm.io/gorm"
)

// PDFReporteService servicio para generar PDFs de reportes
type PDFReporteService struct {
	db *gorm.DB
}

// NewPDFReporteService crea una nueva instancia del servicio
func NewPDFReporteService(db *gorm.DB) *PDFReporteService {
	return &PDFReporteService{db: db}
}

// ReportePDFData contiene todos los datos necesarios para generar el PDF
type ReportePDFData struct {
	Reporte            models.ReporteServicio
	Equipo             models.Equipo
	UsuarioResponsable *models.UsuarioResponsable
	UsuarioSistema     models.Usuario
	Repuestos          []models.Repuesto
	TipoMantenimiento  models.TipoMantenimiento
}

// Constantes de diseño
const (
	pageWidth    = 216.0 // Letter width in mm
	pageHeight   = 279.0 // Letter height in mm
	marginLeft   = 15.0
	marginRight  = 15.0
	contentWidth = pageWidth - marginLeft - marginRight // 186mm

	// Rutas de logos (relativas al ejecutable)
	logoIzquierda = "assets/logos/logo-004.png"        // Logo Alcaldía
	logoDerecha   = "assets/logos/logo-escudo.png"     // Escudo Colombia (con transparencia)
	logoMarcaAgua = "assets/logos/logo-marca-agua.png" // Marca de agua central (con transparencia)
)

// GenerarPDFReporte genera el PDF del reporte de servicio técnico
func (s *PDFReporteService) GenerarPDFReporte(reporteID uint, usuarioSistemaID uint) ([]byte, error) {
	var reporte models.ReporteServicio
	if err := s.db.Preload("Equipo").
		Preload("Equipo.UsuarioResponsable").
		Preload("TipoMantenimiento").
		Preload("Repuestos").
		First(&reporte, reporteID).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo reporte: %w", err)
	}

	var usuarioSistema models.Usuario
	if err := s.db.First(&usuarioSistema, usuarioSistemaID).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo usuario del sistema: %w", err)
	}

	data := ReportePDFData{
		Reporte:            reporte,
		Equipo:             reporte.Equipo,
		UsuarioResponsable: reporte.Equipo.UsuarioResponsable,
		UsuarioSistema:     usuarioSistema,
		Repuestos:          reporte.Repuestos,
		TipoMantenimiento:  reporte.TipoMantenimiento,
	}

	return s.crearPDF(data)
}

func (s *PDFReporteService) crearPDF(data ReportePDFData) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(marginLeft, 15, marginRight)
	pdf.SetAutoPageBreak(true, 15)

	// ========== PÁGINA 1 ==========
	pdf.AddPage()
	s.agregarMarcaAgua(pdf)
	s.agregarEncabezado(pdf)
	s.agregarTitulo(pdf, "REPORTE DE SERVICIO TECNICO")
	s.agregarDatosReporte(pdf, data)
	s.agregarTrabajoRealizado(pdf, data.TipoMantenimiento)
	s.agregarSeccionesTexto(pdf, data.Reporte)
	s.agregarPiePagina(pdf)

	// ========== PÁGINA 2 ==========
	pdf.AddPage()
	s.agregarMarcaAgua(pdf)
	s.agregarEncabezado(pdf)
	s.agregarTitulo(pdf, "REPUESTOS EMPLEADOS Y/O REEMPLAZADO")
	s.agregarTablaRepuestos(pdf, data.Repuestos)
	s.agregarFirmas(pdf, data)
	s.agregarPiePagina(pdf)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("error generando PDF: %w", err)
	}

	return buf.Bytes(), nil
}

func (s *PDFReporteService) agregarMarcaAgua(pdf *fpdf.Fpdf) {
	// Marca de agua central - posición: x=61, y=73, tamaño: 95x135mm
	// Se coloca detrás del contenido con transparencia
	pdf.ImageOptions(logoMarcaAgua, 61, 73, 95, 135, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
}

func (s *PDFReporteService) agregarEncabezado(pdf *fpdf.Fpdf) {
	// Guardar posición Y inicial
	inicioY := pdf.GetY()

	// Logo izquierda (Alcaldía) - posición: x=16, y=7, tamaño: 25x20mm
	pdf.ImageOptions(logoIzquierda, 16, 7, 25, 20, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	// Logo derecha (Escudo Colombia) - posición: x=186, y=7, tamaño: 15x22mm
	pdf.ImageOptions(logoDerecha, 186, 7, 15, 22, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	// Texto del encabezado centrado
	pdf.SetY(inicioY)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(contentWidth, 5, "REPUBLICA DE COLOMBIA", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 5, "ALCALDIA DISTRITAL DE TUMACO", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 5, "SECRETARIA GENERAL", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 5, "OFICINA DE SISTEMAS", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(contentWidth, 5, "NIT. 891.200.916-2", "", 1, "C", false, 0, "")
	pdf.Ln(8)
}

func (s *PDFReporteService) agregarTitulo(pdf *fpdf.Fpdf, titulo string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(contentWidth, 8, titulo, "", 1, "C", false, 0, "")
	pdf.Ln(3)
}

func (s *PDFReporteService) agregarDatosReporte(pdf *fpdf.Fpdf, data ReportePDFData) {
	colWidth := contentWidth / 2 // 93mm cada columna
	rowHeight := 8.0

	fechaInicio := data.Reporte.FechaInicio.Format("02/01/2006")
	fechaFin := ""
	if data.Reporte.FechaFinalizacion != nil {
		fechaFin = data.Reporte.FechaFinalizacion.Format("02/01/2006")
	}

	// Fila 1: Fechas
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "FECHA INICIO:", fechaInicio)
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "FECHA FINALIZACION:", fechaFin)
	pdf.Ln(rowHeight)

	// Fila 2: Dependencia y Ubicación
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "DEPENDENCIA:", data.Reporte.Dependencia)
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "UBICACION:", data.Reporte.Ubicacion)
	pdf.Ln(rowHeight)

	// Fila 3: Equipo y Marca
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "EQUIPO:", data.Equipo.TipoDispositivo)
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "MARCA:", data.Equipo.Marca)
	pdf.Ln(rowHeight)

	// Fila 4: Modelo y Serie
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "MODELO:", data.Equipo.Modelo)
	s.dibujarCeldaConBorde(pdf, colWidth, rowHeight, "SERIE:", data.Equipo.Serial)
	pdf.Ln(rowHeight + 5)
}

func (s *PDFReporteService) dibujarCeldaConBorde(pdf *fpdf.Fpdf, width, height float64, label, value string) {
	x, y := pdf.GetXY()

	// Dibujar borde de la celda
	pdf.Rect(x, y, width, height, "D")

	// Escribir label en negrita
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(x+2, y+2)
	pdf.Cell(35, height-4, label)

	// Escribir valor
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(width-40, height-4, value)

	// Mover X al final de la celda
	pdf.SetXY(x+width, y)
}

func (s *PDFReporteService) agregarTrabajoRealizado(pdf *fpdf.Fpdf, tm models.TipoMantenimiento) {
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(contentWidth, 7, "TRABAJO REALIZADO", "1", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Primera fila: MANTENIMIENTO PREVENTIVO | REVISIÓN | INSTALACIÓN | CONFIGURACION:
	rowY := pdf.GetY()
	pdf.SetFont("Arial", "", 8)

	// Posiciones X ajustadas según el formato original
	x1 := marginLeft       // MANTENIMIENTO PREVENTIVO
	x2 := marginLeft + 55  // REVISIÓN
	x3 := marginLeft + 95  // INSTALACIÓN
	x4 := marginLeft + 140 // CONFIGURACION

	pdf.SetXY(x1, rowY)
	s.dibujarCheckbox(pdf, tm.Tipo == "PREVENTIVO")
	pdf.Cell(2, 5, "")
	pdf.Cell(50, 5, "MANTENIMIENTO PREVENTIVO")

	pdf.SetXY(x2, rowY)
	s.dibujarCheckbox(pdf, tm.Revision)
	pdf.Cell(2, 5, "")
	pdf.Cell(35, 5, "REVISION")

	pdf.SetXY(x3, rowY)
	s.dibujarCheckbox(pdf, tm.Instalacion)
	pdf.Cell(2, 5, "")
	pdf.Cell(40, 5, "INSTALACION")

	pdf.SetXY(x4, rowY)
	s.dibujarCheckbox(pdf, tm.Configuracion)
	pdf.Cell(2, 5, "")
	pdf.Cell(40, 5, "CONFIGURACION:")

	pdf.Ln(10)

	// Segunda fila: MANTENIMIENTO CORRECTIVO | INGRESO | SALIDA | OTRO:
	rowY = pdf.GetY()

	pdf.SetXY(x1, rowY)
	s.dibujarCheckbox(pdf, tm.Tipo == "CORRECTIVO")
	pdf.Cell(2, 5, "")
	pdf.Cell(50, 5, "MANTENIMIENTO CORRECTIVO")

	pdf.SetXY(x2, rowY)
	s.dibujarCheckbox(pdf, tm.Ingreso)
	pdf.Cell(2, 5, "")
	pdf.Cell(35, 5, "INGRESO")

	pdf.SetXY(x3, rowY)
	s.dibujarCheckbox(pdf, tm.Salida)
	pdf.Cell(2, 5, "")
	pdf.Cell(25, 5, "SALIDA")

	pdf.SetXY(x4, rowY)
	s.dibujarCheckbox(pdf, tm.Otro)
	pdf.Cell(2, 5, "")
	pdf.Cell(12, 5, "OTRO:")
	if tm.DescripcionOtro != "" {
		pdf.Cell(30, 5, tm.DescripcionOtro)
	} else {
		pdf.Cell(30, 5, "____________________")
	}

	pdf.Ln(12)
}

func (s *PDFReporteService) dibujarCheckboxConLabel(pdf *fpdf.Fpdf, x, y float64, label string, checked bool) {
	pdf.SetXY(x, y)
	s.dibujarCheckbox(pdf, checked)
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(2, 5, "")
	pdf.Cell(40, 5, label)
}

func (s *PDFReporteService) dibujarCheckbox(pdf *fpdf.Fpdf, checked bool) {
	x, y := pdf.GetXY()
	size := 4.0

	// Dibujar cuadro
	pdf.Rect(x, y, size, size, "D")

	// Si está marcado, dibujar X
	if checked {
		pdf.Line(x, y, x+size, y+size)
		pdf.Line(x+size, y, x, y+size)
	}

	pdf.SetXY(x+size, y)
}

func (s *PDFReporteService) agregarSeccionesTexto(pdf *fpdf.Fpdf, reporte models.ReporteServicio) {
	boxHeight := 25.0

	// DIAGNOSTICO Y/O FALLA REPORTADA
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(contentWidth, 6, "DIAGNOSTICO Y/O FALLA REPORTADA:", "LTR", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	x, y := pdf.GetXY()
	pdf.Rect(x, y, contentWidth, boxHeight, "D")
	pdf.SetXY(x+2, y+2)
	pdf.MultiCell(contentWidth-4, 5, reporte.DiagnosticoFalla, "", "L", false)
	pdf.SetY(y + boxHeight)
	pdf.Ln(3)

	// ACTIVIDAD REALIZADA
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(contentWidth, 6, "ACTIVIDAD REALIZADA:", "LTR", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	x, y = pdf.GetXY()
	pdf.Rect(x, y, contentWidth, boxHeight, "D")
	pdf.SetXY(x+2, y+2)
	pdf.MultiCell(contentWidth-4, 5, reporte.ActividadRealizada, "", "L", false)
	pdf.SetY(y + boxHeight)
	pdf.Ln(3)

	// OBSERVACIONES
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(contentWidth, 6, "OBSERVACIONES:", "LTR", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	x, y = pdf.GetXY()
	pdf.Rect(x, y, contentWidth, boxHeight, "D")
	pdf.SetXY(x+2, y+2)
	observaciones := reporte.Observaciones
	if observaciones == "" {
		observaciones = ""
	}
	pdf.MultiCell(contentWidth-4, 5, observaciones, "", "L", false)
	pdf.SetY(y + boxHeight)
}

func (s *PDFReporteService) agregarTablaRepuestos(pdf *fpdf.Fpdf, repuestos []models.Repuesto) {
	// Anchos de columnas
	col1 := 25.0                       // CANTIDAD
	col2 := 50.0                       // SERIAL O NUMERO DE PARTE
	col3 := contentWidth - col1 - col2 // MARCA, TECNOLOGÍA, CAPACIDAD, DESCRIPCIÓN

	// Encabezados
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(col1, 8, "CANTIDAD", "1", 0, "C", true, 0, "")
	pdf.CellFormat(col2, 8, "SERIAL O NUMERO DE PARTE", "1", 0, "C", true, 0, "")
	pdf.CellFormat(col3, 8, "MARCA, TECNOLOGIA, CAPACIDAD, DESCRIPCION", "1", 1, "C", true, 0, "")

	// Datos
	pdf.SetFont("Arial", "", 8)
	if len(repuestos) == 0 {
		// Filas vacías para el formato
		for i := 0; i < 5; i++ {
			pdf.CellFormat(col1, 8, "", "1", 0, "C", false, 0, "")
			pdf.CellFormat(col2, 8, "", "1", 0, "C", false, 0, "")
			pdf.CellFormat(col3, 8, "", "1", 1, "L", false, 0, "")
		}
	} else {
		for _, r := range repuestos {
			descripcion := fmt.Sprintf("%s, %s, %s, %s", r.Marca, r.Tecnologia, r.Capacidad, r.Descripcion)
			pdf.CellFormat(col1, 8, fmt.Sprintf("%d", r.Cantidad), "1", 0, "C", false, 0, "")
			pdf.CellFormat(col2, 8, r.SerialNumeroParte, "1", 0, "C", false, 0, "")
			pdf.CellFormat(col3, 8, s.truncarTexto(descripcion, 55), "1", 1, "L", false, 0, "")
		}
		// Rellenar con filas vacías hasta tener al menos 5
		for i := len(repuestos); i < 5; i++ {
			pdf.CellFormat(col1, 8, "", "1", 0, "C", false, 0, "")
			pdf.CellFormat(col2, 8, "", "1", 0, "C", false, 0, "")
			pdf.CellFormat(col3, 8, "", "1", 1, "L", false, 0, "")
		}
	}
	pdf.Ln(10)
}

func (s *PDFReporteService) truncarTexto(texto string, maxLen int) string {
	if len(texto) > maxLen {
		return texto[:maxLen-3] + "..."
	}
	return texto
}

func (s *PDFReporteService) agregarFirmas(pdf *fpdf.Fpdf, data ReportePDFData) {
	colWidth := contentWidth / 2
	rowHeight := 7.0

	// Títulos de las columnas
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(colWidth, rowHeight, "FUNCIONARIO Y/O CONTRATISTA DEL SERVICIO", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidth, rowHeight, "FUNCIONARIO Y/O CONTRATISTA DE SISTEMAS", "1", 1, "C", false, 0, "")

	// Obtener datos
	nombreResponsable := ""
	cedulaResponsable := ""
	cargoResponsable := ""
	if data.UsuarioResponsable != nil {
		nombreResponsable = data.UsuarioResponsable.NombresApellidos
		cedulaResponsable = data.UsuarioResponsable.Cedula
		cargoResponsable = data.UsuarioResponsable.TipoVinculacion
	}

	nombreSistema := fmt.Sprintf("%s %s", data.UsuarioSistema.Nombre, data.UsuarioSistema.Apellido)
	cedulaSistema := data.UsuarioSistema.Cedula
	cargoSistema := strings.ToUpper(data.UsuarioSistema.Rol)

	// Fila NOMBRE
	pdf.SetFont("Arial", "B", 8)
	x := pdf.GetX()
	y := pdf.GetY()
	pdf.Rect(x, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+2, y+1)
	pdf.Cell(18, 5, "NOMBRE:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-22, 5, nombreResponsable)

	pdf.SetXY(x+colWidth, y)
	pdf.Rect(x+colWidth, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+colWidth+2, y+1)
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(18, 5, "NOMBRE:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-22, 5, nombreSistema)
	pdf.Ln(rowHeight)

	// Fila CARGO
	x = pdf.GetX()
	y = pdf.GetY()
	pdf.Rect(x, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+2, y+1)
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(18, 5, "CARGO:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-22, 5, cargoResponsable)

	pdf.SetXY(x+colWidth, y)
	pdf.Rect(x+colWidth, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+colWidth+2, y+1)
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(18, 5, "CARGO:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-22, 5, cargoSistema)
	pdf.Ln(rowHeight)

	// Espacio para firma
	x = pdf.GetX()
	y = pdf.GetY()
	firmaHeight := 20.0
	pdf.Rect(x, y, colWidth, firmaHeight, "D")
	pdf.Rect(x+colWidth, y, colWidth, firmaHeight, "D")
	pdf.Ln(firmaHeight)

	// Fila FIRMA y C.C.
	x = pdf.GetX()
	y = pdf.GetY()
	pdf.Rect(x, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+2, y+1)
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(15, 5, "FIRMA")
	pdf.Cell(20, 5, "")
	pdf.Cell(10, 5, "C.C:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-50, 5, cedulaResponsable)

	pdf.SetXY(x+colWidth, y)
	pdf.Rect(x+colWidth, y, colWidth, rowHeight, "D")
	pdf.SetXY(x+colWidth+2, y+1)
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(15, 5, "FIRMA")
	pdf.Cell(20, 5, "")
	pdf.Cell(10, 5, "C.C:")
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(colWidth-50, 5, cedulaSistema)
	pdf.Ln(rowHeight)
}

func (s *PDFReporteService) agregarPiePagina(pdf *fpdf.Fpdf) {
	pdf.SetY(-35)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(contentWidth, 4, "Calle 11 con Carrera 9a esquina - Edificio Municipal, Telefax (2) 727 12 01", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 4, "Pagina Web: www.tumaco-narino.gov.co", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 4, "Correo electronico: contactenos@tumaco-narino.gov.co", "", 1, "C", false, 0, "")
	pdf.CellFormat(contentWidth, 4, "Tumaco - Narino", "", 1, "C", false, 0, "")
}
