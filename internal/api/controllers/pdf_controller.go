package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// PDFController maneja las solicitudes HTTP para generar PDFs
type PDFController struct {
	pdfService *services.PDFReporteService
}

// NewPDFController crea una nueva instancia de PDFController
func NewPDFController(pdfService *services.PDFReporteService) *PDFController {
	return &PDFController{
		pdfService: pdfService,
	}
}

// GenerarReportePDF genera el PDF de un reporte de servicio técnico
// GET /api/reportes-servicio/:id/pdf
func (c *PDFController) GenerarReportePDF(ctx echo.Context) error {
	// Obtener ID del reporte
	reporteID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de reporte inválido"})
	}

	// Obtener el usuario de la sesión (del contexto JWT)
	usuarioID, ok := ctx.Get("usuario_id").(uint)
	if !ok {
		// Si no hay usuario en contexto, intentar obtenerlo del query param (para pruebas)
		userIDParam := ctx.QueryParam("usuario_id")
		if userIDParam != "" {
			parsedID, err := strconv.ParseUint(userIDParam, 10, 32)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
			}
			usuarioID = uint(parsedID)
		} else {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Usuario no autenticado"})
		}
	}

	// Generar el PDF
	pdfBytes, err := c.pdfService.GenerarPDFReporte(uint(reporteID), usuarioID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Configurar headers para descarga del PDF
	ctx.Response().Header().Set("Content-Type", "application/pdf")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=reporte_servicio_%d.pdf", reporteID))
	ctx.Response().Header().Set("Content-Length", strconv.Itoa(len(pdfBytes)))

	return ctx.Blob(http.StatusOK, "application/pdf", pdfBytes)
}

// VisualizarReportePDF genera el PDF para visualización en el navegador
// GET /api/reportes-servicio/:id/pdf/view
func (c *PDFController) VisualizarReportePDF(ctx echo.Context) error {
	// Obtener ID del reporte
	reporteID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de reporte inválido"})
	}

	// Obtener el usuario de la sesión (del contexto JWT)
	usuarioID, ok := ctx.Get("usuario_id").(uint)
	if !ok {
		// Si no hay usuario en contexto, intentar obtenerlo del query param (para pruebas)
		userIDParam := ctx.QueryParam("usuario_id")
		if userIDParam != "" {
			parsedID, err := strconv.ParseUint(userIDParam, 10, 32)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
			}
			usuarioID = uint(parsedID)
		} else {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Usuario no autenticado"})
		}
	}

	// Generar el PDF
	pdfBytes, err := c.pdfService.GenerarPDFReporte(uint(reporteID), usuarioID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Configurar headers para visualización inline
	ctx.Response().Header().Set("Content-Type", "application/pdf")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=reporte_servicio_%d.pdf", reporteID))

	return ctx.Blob(http.StatusOK, "application/pdf", pdfBytes)
}
