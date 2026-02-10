package controllers

import (
	"net/http"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// DashboardController maneja las peticiones del dashboard
type DashboardController struct {
	service services.DashboardService
}

// NewDashboardController crea una nueva instancia del controlador
func NewDashboardController(service services.DashboardService) *DashboardController {
	return &DashboardController{service: service}
}

// GetDashboardStats retorna todas las estadísticas del dashboard en una sola petición
func (dc *DashboardController) GetDashboardStats(c echo.Context) error {
	stats, err := dc.service.GetDashboardStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al obtener estadísticas del dashboard: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, stats)
}
