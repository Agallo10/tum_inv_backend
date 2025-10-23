package routes

import (
	"tum_inv_backend/internal/api/controllers"
	"tum_inv_backend/internal/api/middleware"
	"tum_inv_backend/internal/domain/repositories"
	"tum_inv_backend/internal/domain/services"
	"tum_inv_backend/internal/infrastructure/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	// Repositorios
	equipoRepo := repositories.NewEquipoRepository(db)
	perifericoRepo := repositories.NewPerifericoRepository(db)
	softwareRepo := repositories.NewSoftwareRepository(db)
	usuarioResponsableRepo := repositories.NewUsuarioResponsableRepository(db)
	hardwareInternoRepo := repositories.NewHardwareInternoRepository(db)
	configuracionRedRepo := repositories.NewConfiguracionRedRepository(db)
	usuarioSistemaRepo := repositories.NewUsuarioSistemaRepository(db)
	accesoRemotoRepo := repositories.NewAccesoRemotoRepository(db)
	backupRepo := repositories.NewBackupRepository(db)
	reporteServicioRepo := repositories.NewReporteServicioRepository(db)
	tipoMantenimientoRepo := repositories.NewTipoMantenimientoRepository(db)
	repuestoRepo := repositories.NewRepuestoRepository(db)
	funcionarioRepo := repositories.NewFuncionarioRepository(db)
	usuarioRepo := repositories.NewUsuarioRepository(db)
	secretariaRepo := repositories.NewSecretariaRepository(db)
	dependenciaRepo := repositories.NewDependenciaRepository(db)
	estadoEquipoRepo := repositories.NewEstadoEquipoRepository(db)

	// Servicios
	equipoService := services.NewEquipoService(equipoRepo)
	perifericoService := services.NewPerifericoService(perifericoRepo)
	softwareService := services.NewSoftwareService((softwareRepo))
	usuarioResponsableService := services.NewUsuarioResponsableService(usuarioResponsableRepo)
	hardwareInternoService := services.NewHardwareInternoService(hardwareInternoRepo)
	configuracionRedService := services.NewConfiguracionRedService(configuracionRedRepo)
	usuarioSistemaService := services.NewUsuarioSistemaService(usuarioSistemaRepo)
	accesoRemotoService := services.NewAccesoRemotoService(accesoRemotoRepo)
	backupService := services.NewBackupService(backupRepo)
	reporteServicioService := services.NewReporteServicioService(reporteServicioRepo, funcionarioRepo)
	tipoMantenimientoService := services.NewTipoMantenimientoService(tipoMantenimientoRepo)
	repuestoService := services.NewRepuestoService(repuestoRepo)
	funcionarioService := services.NewFuncionarioService(funcionarioRepo)
	authService := services.NewAuthService(usuarioRepo, cfg)
	secretariaService := services.NewSecretariaService(secretariaRepo)
	dependenciaService := services.NewDependenciaService(dependenciaRepo)
	estadoEquipoService := services.NewEstadoEquipoService(estadoEquipoRepo)

	// Controladores
	equipoController := controllers.NewEquipoController(equipoService)
	perifericoController := controllers.NewPerifericoController(perifericoService)
	softwareController := controllers.NewSoftwareController(softwareService)
	usuarioResponsableController := controllers.NewUsuarioResponsableController(usuarioResponsableService)
	hardwareInternoController := controllers.NewHardwareInternoController(hardwareInternoService)
	configuracionRedController := controllers.NewConfiguracionRedController(configuracionRedService)
	usuarioSistemaController := controllers.NewUsuarioSistemaController(usuarioSistemaService)
	accesoRemotoController := controllers.NewAccesoRemotoController(accesoRemotoService)
	backupController := controllers.NewBackupController(backupService)
	reporteServicioController := controllers.NewReporteServicioController(reporteServicioService)
	tipoMantenimientoController := controllers.NewTipoMantenimientoController(tipoMantenimientoService)
	repuestoController := controllers.NewRepuestoController(repuestoService)
	funcionarioController := controllers.NewFuncionarioController(funcionarioService)
	authController := controllers.NewAuthController(authService)
	secretariaController := controllers.NewSecretariaController(secretariaService)
	dependenciaController := controllers.NewDependenciaController(dependenciaService)
	estadoEquipoController := controllers.NewEstadoEquipoController(estadoEquipoService)

	// Middleware
	jwtMiddleware := middleware.NewJWTMiddleware(authService)

	// Grupo de rutas para API
	api := e.Group("/api")

	// Rutas de autenticación (públicas)
	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.POST("/refresh", authController.RefreshToken)

	// Ruta protegida para obtener perfil de usuario
	auth.GET("/profile", authController.GetProfile, jwtMiddleware.Authenticate)

	// Rutas para Equipos
	equipos := api.Group("/equipos")
	equipos.POST("", equipoController.CreateEquipo)
	equipos.GET("", equipoController.GetAllEquipos)
	equipos.GET("/:id", equipoController.GetEquipo)
	equipos.PUT("/:id", equipoController.UpdateEquipo)
	equipos.DELETE("/:id", equipoController.DeleteEquipo)
	// Ruta para obtener equipos dpor dependencia
	equipos.GET("/:dependenciaId/dependencia", equipoController.GetEquiposByDependencia)
	// Ruta para obtener la hoja de vida del equipo
	equipos.GET("/:equipoId/hv", equipoController.GetEquipoUsuDepByID)

	// Rutas para Periféricos
	perifericos := api.Group("/perifericos")
	perifericos.POST("", perifericoController.CreatePeriferico)
	perifericos.GET("", perifericoController.GetAllPerifericos)
	perifericos.GET("/:id", perifericoController.GetPeriferico)
	perifericos.PUT("/:id", perifericoController.UpdatePeriferico)
	perifericos.DELETE("/:id", perifericoController.DeletePeriferico)

	// Ruta para obtener periféricos por equipo
	equipos.GET("/:equipoId/perifericos", perifericoController.GetPerifericosByEquipo)

	// Rutas para Software
	software := api.Group("/software")
	software.POST("", softwareController.CreateSoftware)
	software.GET("", softwareController.GetAllSoftware)
	software.GET("/:id", softwareController.GetSoftware)
	software.PUT("/:id", softwareController.UpdateSoftware)
	software.DELETE("/:id", softwareController.DeleteSoftware)

	// Ruta para obtener periféricos por equipo
	equipos.GET("/:equipoId/software", softwareController.GetAllSoftwareByEquipo)

	// Rutas para Usuarios Responsables
	usuariosResponsables := api.Group("/usuarios-responsables")
	usuariosResponsables.POST("", usuarioResponsableController.CreateUsuarioResponsable)
	usuariosResponsables.GET("", usuarioResponsableController.GetAllUsuariosResponsables)
	usuariosResponsables.GET("/buscar", usuarioResponsableController.GetUsuarioResponsableByCedula)
	usuariosResponsables.GET("/:id", usuarioResponsableController.GetUsuarioResponsable)
	usuariosResponsables.PUT("/:id", usuarioResponsableController.UpdateUsuarioResponsable)
	usuariosResponsables.DELETE("/:id", usuarioResponsableController.DeleteUsuarioResponsable)
	// usuariosResponsables.GET("/:id/equipos", usuarioResponsableController.GetEquiposByUsuarioResponsable)
	usuariosResponsables.GET("/:dependenciaId/dependencia", usuarioResponsableController.GetUsuariosByDependencia)

	// Rutas para Hardware Interno
	hardwareInterno := api.Group("/hardware-interno")
	hardwareInterno.POST("", hardwareInternoController.CreateHardwareInterno)
	hardwareInterno.GET("", hardwareInternoController.GetAllHardwareInterno)
	hardwareInterno.GET("/:id", hardwareInternoController.GetHardwareInterno)
	hardwareInterno.PUT("/:id", hardwareInternoController.UpdateHardwareInterno)
	hardwareInterno.DELETE("/:id", hardwareInternoController.DeleteHardwareInterno)

	// Ruta para obtener hardware interno por equipo
	equipos.GET("/:equipoId/hardware-interno", hardwareInternoController.GetHardwareInternoByEquipo)

	// Rutas para Configuración de Red
	configuracionesRed := api.Group("/configuraciones-red")
	configuracionesRed.POST("", configuracionRedController.CreateConfiguracionRed)
	configuracionesRed.GET("", configuracionRedController.GetAllConfiguracionesRed)
	configuracionesRed.GET("/:id", configuracionRedController.GetConfiguracionRed)
	configuracionesRed.PUT("/:id", configuracionRedController.UpdateConfiguracionRed)
	configuracionesRed.DELETE("/:id", configuracionRedController.DeleteConfiguracionRed)

	// Ruta para obtener configuración de red por equipo
	equipos.GET("/:equipoId/configuracion-red", configuracionRedController.GetConfiguracionRedByEquipo)

	// Rutas para Usuarios del Sistema
	usuariosSistema := api.Group("/usuarios-sistema")
	usuariosSistema.POST("", usuarioSistemaController.CreateUsuarioSistema)
	usuariosSistema.GET("", usuarioSistemaController.GetAllUsuariosSistema)
	usuariosSistema.GET("/buscar", usuarioSistemaController.GetUsuarioSistemaByNombreUsuario)
	usuariosSistema.GET("/:id", usuarioSistemaController.GetUsuarioSistema)
	usuariosSistema.PUT("/:id", usuarioSistemaController.UpdateUsuarioSistema)
	usuariosSistema.DELETE("/:id", usuarioSistemaController.DeleteUsuarioSistema)

	// Ruta para obtener usuarios del sistema por equipo
	equipos.GET("/:equipoId/usuarios-sistema", usuarioSistemaController.GetUsuariosSistemaByEquipo)

	// Rutas para Accesos Remotos
	accesosRemotos := api.Group("/accesos-remotos")
	accesosRemotos.POST("", accesoRemotoController.CreateAccesoRemoto)
	accesosRemotos.GET("", accesoRemotoController.GetAllAccesosRemotos)
	accesosRemotos.GET("/:id", accesoRemotoController.GetAccesoRemoto)
	accesosRemotos.PUT("/:id", accesoRemotoController.UpdateAccesoRemoto)
	accesosRemotos.DELETE("/:id", accesoRemotoController.DeleteAccesoRemoto)

	// Ruta para obtener accesos remotos por equipo
	equipos.GET("/:equipoId/accesos-remotos", accesoRemotoController.GetAccesosRemotosByEquipo)

	// Rutas para Backups
	backups := api.Group("/backups")
	backups.POST("", backupController.CreateBackup)
	backups.GET("", backupController.GetAllBackups)
	backups.GET("/:id", backupController.GetBackup)
	backups.PUT("/:id", backupController.UpdateBackup)
	backups.DELETE("/:id", backupController.DeleteBackup)

	// Ruta para obtener backups por equipo
	equipos.GET("/:equipoId/backups", backupController.GetBackupsByEquipo)

	// Rutas para Reportes de Servicio
	reportesServicio := api.Group("/reportes-servicio")
	reportesServicio.POST("", reporteServicioController.CreateReporteServicio)
	reportesServicio.POST("/completo", reporteServicioController.CrearReporteConTipo)
	reportesServicio.GET("", reporteServicioController.GetAllReportesServicio)
	reportesServicio.GET("/:id", reporteServicioController.GetReporteServicio)
	reportesServicio.PUT("/:id", reporteServicioController.UpdateReporteServicio)
	reportesServicio.DELETE("/:id", reporteServicioController.DeleteReporteServicio)

	// Ruta para obtener reportes de servicio por equipo
	equipos.GET("/:equipoId/reportes-servicio", reporteServicioController.GetReportesServicioByEquipo)

	// Rutas para Tipos de Mantenimiento
	tiposMantenimiento := api.Group("/tipos-mantenimiento")
	tiposMantenimiento.POST("", tipoMantenimientoController.CreateTipoMantenimiento)
	tiposMantenimiento.GET("", tipoMantenimientoController.GetAllTiposMantenimiento)
	tiposMantenimiento.GET("/:id", tipoMantenimientoController.GetTipoMantenimiento)
	tiposMantenimiento.PUT("/:id", tipoMantenimientoController.UpdateTipoMantenimiento)
	tiposMantenimiento.DELETE("/:id", tipoMantenimientoController.DeleteTipoMantenimiento)

	// Ruta para obtener tipos de mantenimiento por reporte
	reportesServicio.GET("/:reporteId/tipos-mantenimiento", tipoMantenimientoController.GetTiposMantenimientoByReporte)

	// Rutas para Repuestos
	repuestos := api.Group("/repuestos")
	repuestos.POST("", repuestoController.CreateRepuesto)
	repuestos.GET("", repuestoController.GetAllRepuestos)
	repuestos.GET("/:id", repuestoController.GetRepuesto)
	repuestos.PUT("/:id", repuestoController.UpdateRepuesto)
	repuestos.DELETE("/:id", repuestoController.DeleteRepuesto)

	// Ruta para obtener repuestos por reporte
	reportesServicio.GET("/:reporteId/repuestos", repuestoController.GetRepuestosByReporte)

	// Rutas para Funcionarios
	funcionarios := api.Group("/funcionarios")
	funcionarios.POST("", funcionarioController.CreateFuncionario)
	funcionarios.GET("", funcionarioController.GetAllFuncionarios)
	funcionarios.GET("/buscar", funcionarioController.GetFuncionarioByCedula)
	funcionarios.GET("/:id", funcionarioController.GetFuncionario)
	funcionarios.PUT("/:id", funcionarioController.UpdateFuncionario)
	funcionarios.DELETE("/:id", funcionarioController.DeleteFuncionario)

	// Ruta para obtener funcionarios por reporte
	reportesServicio.GET("/:reporteId/funcionarios", funcionarioController.GetFuncionariosByReporte)

	// Rutas para Secretarías/s
	secretarias := api.Group("/secretarias")
	secretarias.POST("", secretariaController.CreateSecretaria)
	secretarias.GET("", secretariaController.GetAllSecretarias)
	secretarias.GET("/:id", secretariaController.GetSecretaria)
	secretarias.PUT("/:id", secretariaController.UpdateSecretaria)
	secretarias.DELETE("/:id", secretariaController.DeleteSecretaria)
	secretarias.GET("/:id/dependencias", secretariaController.GetDependenciasBySecretaria)

	// Rutas para Dependencias
	dependencias := api.Group("/dependencias")
	dependencias.POST("", dependenciaController.CreateDependencia)
	dependencias.GET("", dependenciaController.GetAllDependencias)
	dependencias.GET("/:id", dependenciaController.GetDependencia)
	dependencias.PUT("/:id", dependenciaController.UpdateDependencia)
	dependencias.DELETE("/:id", dependenciaController.DeleteDependencia)
	dependencias.GET("/:id/usuarios", dependenciaController.GetUsuariosByDependencia)
	// dependencias.GET("/:id/equipos", dependenciaController.GetEquiposByDependencia)

	// Ruta para obtener dependencias por secretaría
	secretarias.GET("/:secretariaId/dependencias", dependenciaController.GetDependenciasBySecretaria)

	// Rutas para Estados de Equipo
	estadosEquipo := api.Group("/estados-equipo")
	estadosEquipo.POST("", estadoEquipoController.CreateEstado)
	estadosEquipo.GET("", estadoEquipoController.GetAllEstados)
	estadosEquipo.GET("/activos", estadoEquipoController.GetActiveEstados)
	estadosEquipo.GET("/:id", estadoEquipoController.GetEstadoByID)
	estadosEquipo.PUT("/:id", estadoEquipoController.UpdateEstado)
	estadosEquipo.DELETE("/:id", estadoEquipoController.DeleteEstado)
	estadosEquipo.PATCH("/:id/toggle-activo", estadoEquipoController.ToggleActivo)
	estadosEquipo.GET("/:id/equipos", estadoEquipoController.GetEquiposByEstado)
}
