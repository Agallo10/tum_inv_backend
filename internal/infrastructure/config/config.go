package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config almacena todas las configuraciones de la aplicación
type Config struct {
	DatabaseURL string // Railway provee esta variable directamente
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	DBTimeout   time.Duration
	AppPort     string
	AppEnv      string
	JWTSecret   string
	FrontendURL string
}

// LoadConfig carga la configuración desde variables de entorno
func LoadConfig() *Config {
	// Cargar archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno del sistema")
	}

	// Obtener tiempo de espera para DB
	timeout, err := strconv.Atoi(getEnv("DB_TIMEOUT", "10"))
	if err != nil {
		log.Fatalf("Valor inválido para DB_TIMEOUT: %v", err)
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""), // Railway provee esta variable
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "inventario"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
		DBTimeout:   time.Duration(timeout) * time.Second,
		AppPort:     getEnv("APP_PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "tu_clave_secreta_jwt_super_segura"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
	}
}

// getEnv obtiene una variable de entorno o devuelve un valor predeterminado
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
