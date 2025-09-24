package main

import (
	"time"
	"tum_inv_backend/internal/api/routes"
	"tum_inv_backend/internal/infrastructure/config"
	"tum_inv_backend/internal/infrastructure/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()
	// Inicializar Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	// Conectar a la base de datos
	database.ConnectDB(cfg)

	// Configurar rutas usando la variable global DB y la configuración
	routes.SetupRoutes(e, database.DB, cfg)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
