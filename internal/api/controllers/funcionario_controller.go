package controllers

import (
	"net/http"
	"strconv"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/domain/services"

	"github.com/labstack/echo/v4"
)

// FuncionarioController maneja las solicitudes HTTP relacionadas con funcionarios
type FuncionarioController struct {
	funcionarioService services.FuncionarioService
}

// NewFuncionarioController crea una nueva instancia de FuncionarioController
func NewFuncionarioController(funcionarioService services.FuncionarioService) *FuncionarioController {
	return &FuncionarioController{
		funcionarioService: funcionarioService,
	}
}

// CreateFuncionario maneja la creación de un nuevo funcionario
func (c *FuncionarioController) CreateFuncionario(ctx echo.Context) error {
	funcionario := new(models.Funcionario)
	if err := ctx.Bind(funcionario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	if err := c.funcionarioService.CreateFuncionario(funcionario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, funcionario)
}

// GetFuncionario obtiene un funcionario por su ID
func (c *FuncionarioController) GetFuncionario(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	funcionario, err := c.funcionarioService.GetFuncionarioByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Funcionario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, funcionario)
}

// UpdateFuncionario actualiza un funcionario existente
func (c *FuncionarioController) UpdateFuncionario(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	funcionario := new(models.Funcionario)
	if err := ctx.Bind(funcionario); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	funcionario.ID = uint(id)
	if err := c.funcionarioService.UpdateFuncionario(funcionario); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, funcionario)
}

// DeleteFuncionario elimina un funcionario por su ID
func (c *FuncionarioController) DeleteFuncionario(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := c.funcionarioService.DeleteFuncionario(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"mensaje": "Funcionario eliminado correctamente"})
}

// GetAllFuncionarios obtiene todos los funcionarios
func (c *FuncionarioController) GetAllFuncionarios(ctx echo.Context) error {
	funcionarios, err := c.funcionarioService.GetAllFuncionarios()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, funcionarios)
}

// GetFuncionarioByCedula obtiene un funcionario por su cédula
func (c *FuncionarioController) GetFuncionarioByCedula(ctx echo.Context) error {
	cedula := ctx.QueryParam("cedula")
	if cedula == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Cédula no proporcionada"})
	}

	funcionario, err := c.funcionarioService.GetFuncionarioByCedula(cedula)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Funcionario no encontrado"})
	}

	return ctx.JSON(http.StatusOK, funcionario)
}

// GetFuncionariosByReporte obtiene todos los funcionarios asociados a un reporte
func (c *FuncionarioController) GetFuncionariosByReporte(ctx echo.Context) error {
	reporteID, err := strconv.ParseUint(ctx.Param("reporteId"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de reporte inválido"})
	}

	funcionarios, err := c.funcionarioService.GetFuncionariosByReporteID(uint(reporteID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, funcionarios)
}