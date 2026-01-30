package database

import (
	"fmt"
	"log"
	"time"
	"tum_inv_backend/internal/domain/models"
	"tum_inv_backend/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB establece la conexión con la base de datos y devuelve la instancia de DB
func ConnectDB(cfg *config.Config) *gorm.DB {
	var dsn string

	// Railway provee DATABASE_URL directamente
	if cfg.DatabaseURL != "" {
		dsn = cfg.DatabaseURL
		log.Println("Usando DATABASE_URL para conexión (Railway)")
	} else {
		// Fallback a variables individuales (desarrollo local)
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
			cfg.DBSSLMode,
		)
		log.Println("Usando variables individuales para conexión")
	}

	// Configurar logger de GORM según el entorno
	gormConfig := &gorm.Config{}
	if cfg.AppEnv == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	// Intentar conexión con reintentos
	var err error
	for i := 0; i < 3; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err == nil {
			break
		}

		log.Printf("Intento %d: Error al conectar a la base de datos: %v", i+1, err)
		time.Sleep(cfg.DBTimeout / 3)
	}

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos después de 3 intentos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida")

	// Configurar conexión de la base de datos subyacente
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión de base de datos: %v", err)
	}

	// Configurar pool de conexiones
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrar modelos (mantén el mismo código de migración que antes)
	log.Println("Ejecutando migraciones...")
	err = DB.AutoMigrate(
		&models.Secretaria{},
		&models.Dependencia{},
		&models.UsuarioResponsable{},
		&models.Equipo{},
		&models.Periferico{},
		&models.HardwareInterno{},
		&models.Software{},
		&models.ConfiguracionRed{},
		&models.UsuarioSistema{},
		&models.AccesoRemoto{},
		&models.Backup{},
		&models.ReporteServicio{},
		&models.TipoMantenimiento{},
		&models.Repuesto{},
		&models.EstadoEquipo{},
		&models.Usuario{},
	)

	if err != nil {
		log.Fatal("Error al migrar modelos: ", err)
	}

	log.Println("Migraciones completadas con éxito")

	return DB
}
