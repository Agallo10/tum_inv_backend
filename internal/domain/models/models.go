package models

import (
	"time"

	"gorm.io/gorm"
)

// Secretaria representa a la secretaria a cargo de la dependencia
type Secretaria struct {
	gorm.Model
	Nombre      string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	Ubicacion   string `gorm:"not null"`
	Secretario  string `gorm:"not null"`
	Telefono    string

	//Relaciones
	Dependencias []Dependencia `gorm:"foreignKey:SecretariaID"`
}

// Dependencia representa a la dependencia a cargo del usuario responsable
type Dependencia struct {
	gorm.Model
	SecretariaID        uint   `gorm:"not null"`
	Nombre              string `gorm:"not null"`
	Descripcion         string `gorm:"not null"`
	UbicacionOficina    string `gorm:"not null"`
	JefeOficina         string `gorm:"not null"`
	CorreoInstitucional string `gorm:"not null"`
	Telefono            string

	//Relaciones
	UsuarioResponsables []UsuarioResponsable `gorm:"foreignKey:DependenciaID"`
	// Equipos             []Equipo             `gorm:"foreignKey:DependenciaID"`
}

// UsuarioResponsable representa al usuario a cargo del equipo
type UsuarioResponsable struct {
	gorm.Model
	DependenciaID    uint   `gorm:"not null"`
	NombresApellidos string `gorm:"not null"`
	Cedula           string `gorm:"unique;not null"`
	CorreoPersonal   string
	TipoVinculacion  string `gorm:"check:tipo_vinculacion IN ('Planta', 'Contratista', 'Otro')"`
	Celular          string

	// Relaciones
	Equipo Equipo `gorm:"foreignKey:UsuarioResponsableID"`
}

// Equipo representa un dispositivo tecnológico
type Equipo struct {
	gorm.Model
	UsuarioResponsableID   *uint
	TipoDispositivo        string `gorm:"check:tipo_dispositivo IN ('Todo en Uno', 'Escritorio', 'Portátil', 'Impresora', 'Escáner', 'Otro')"`
	PlacaInventario        string `gorm:"unique"`
	Marca                  string `gorm:"not null"`
	Serial                 string `gorm:"unique;not null"`
	Modelo                 string
	FechaDiligenciamiento  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ObservacionesGenerales string

	// Relaciones
	Perifericos      []Periferico      `gorm:"foreignKey:EquipoID"`
	HardwareInterno  []HardwareInterno `gorm:"foreignKey:EquipoID"`
	Software         []Software        `gorm:"foreignKey:EquipoID"`
	ConfiguracionRed ConfiguracionRed  `gorm:"foreignKey:EquipoID"`
	UsuariosSistema  []UsuarioSistema  `gorm:"foreignKey:EquipoID"`
	AccesosRemotos   []AccesoRemoto    `gorm:"foreignKey:EquipoID"`
	Backups          []Backup          `gorm:"foreignKey:EquipoID"`
	Reportes         []ReporteServicio `gorm:"foreignKey:EquipoID"`
}

// Periferico representa dispositivos conectados al equipo
type Periferico struct {
	gorm.Model
	EquipoID       uint
	TipoPeriferico string `gorm:"check:tipo_periferico IN ('Teclado', 'Mouse', 'Monitor', 'Otros')"`
	// TipoPeriferico  string `gorm:"check:tipo_periferico IN ('Teclado', 'Mouse', 'Monitor', 'Impresora', 'Escáner', 'Otros')"`
	PlacaInventario string
	Marca           string
	Serial          string
}

// HardwareInterno representa componentes internos del equipo
type HardwareInterno struct {
	gorm.Model
	EquipoID   uint   `gorm:"not null"`
	Componente string `gorm:"check:componente IN ('Disco Duro', 'Memoria RAM', 'Procesador')"`
	Tecnologia string `gorm:"not null"`
	Capacidad  string `gorm:"not null"`
}

// Software representa programas instalados en el equipo
type Software struct {
	gorm.Model
	EquipoID     uint   `gorm:"not null"`
	Nombre       string `gorm:"not null"`
	Version      string
	TipoLicencia string
	Categoria    string `gorm:"check:categoria IN ('Sistema Operativo', 'Paquete de Oficina', 'Navegador Web', 'Otro')"`
}

// ConfiguracionRed representa la configuración de red del equipo
type ConfiguracionRed struct {
	gorm.Model
	EquipoID          uint   `gorm:"unique;not null"`
	DireccionIP       string `gorm:"not null"`
	AsignacionIP      string `gorm:"check:asignacion_ip IN ('Manual', 'Automatica', 'Dinamica')"`
	NombreDispositivo string `gorm:"not null"`
	Conectividad      string
	// Estado            string `gorm:"check:estado IN ('Activo', 'Inactivo', 'Mantenimiento')"`
}

// UsuarioSistema representa usuarios locales del equipo
type UsuarioSistema struct {
	gorm.Model
	EquipoID        uint   `gorm:"not null"`
	NombreUsuario   string `gorm:"not null"`
	Contrasena      string
	EsAdministrador bool `gorm:"default:false"`
}

// AccesoRemoto representa credenciales de acceso remoto
type AccesoRemoto struct {
	gorm.Model
	EquipoID   uint   `gorm:"not null"`
	Plataforma string `gorm:"default:'AnyDesk'"`
	Usuario    string `gorm:"not null"`
	Contrasena string
	IDConexion string `gorm:"not null"`
}

// Backup representa copias de seguridad realizadas
type Backup struct {
	gorm.Model
	EquipoID          uint      `gorm:"not null"`
	Fecha             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	NumCarpetas       int
	PesoTotalArchivos string
	RutaBackup        string `gorm:"not null"`
	SeRealizoBackup   bool   `gorm:"not null"`
}

// ReporteServicio representa los reportes técnicos
type ReporteServicio struct {
	gorm.Model
	EquipoID           *uint
	FechaInicio        time.Time `gorm:"not null"`
	FechaFinalizacion  *time.Time
	Dependencia        string `gorm:"not null"`
	Ubicacion          string `gorm:"not null"`
	DiagnosticoFalla   string
	ActividadRealizada string `gorm:"not null"`
	Observaciones      string

	// Relaciones
	TiposMantenimiento []TipoMantenimiento `gorm:"foreignKey:ReporteID"`
	Repuestos          []Repuesto          `gorm:"foreignKey:ReporteID"`
	Funcionarios       []Funcionario       `gorm:"many2many:reportes_funcionarios;"`
}

// TipoMantenimiento representa el tipo de mantenimiento realizado
type TipoMantenimiento struct {
	gorm.Model
	ReporteID       uint   `gorm:"not null"`
	Tipo            string `gorm:"check:tipo IN ('PREVENTIVO', 'CORRECTIVO')"`
	Revision        bool   `gorm:"default:false"`
	Instalacion     bool   `gorm:"default:false"`
	Configuracion   bool   `gorm:"default:false"`
	Ingreso         bool   `gorm:"default:false"`
	Salida          bool   `gorm:"default:false"`
	Otro            bool   `gorm:"default:false"`
	DescripcionOtro string
}

// Repuesto representa repuestos utilizados en mantenimientos
type Repuesto struct {
	gorm.Model
	ReporteID         *uint
	Cantidad          int    `gorm:"check:cantidad > 0"`
	SerialNumeroParte string `gorm:"not null"`
	Marca             string
	Tecnologia        string
	Capacidad         string
	Descripcion       string    `gorm:"not null"`
	FechaUtilizacion  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Funcionario representa al personal técnico
type Funcionario struct {
	gorm.Model
	Nombre string `gorm:"not null"`
	Cargo  string `gorm:"not null"`
	Cedula string `gorm:"unique;not null"`
	Tipo   string `gorm:"check:tipo IN ('FUNCIONARIO', 'CONTRATISTA')"`
	Area   string `gorm:"check:area IN ('SERVICIO', 'SISTEMAS')"`

	// Relación muchos a muchos con reportes
	Reportes []ReporteServicio `gorm:"many2many:reportes_funcionarios;"`
}
