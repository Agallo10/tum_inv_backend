package services

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// DashboardStats representa las estadísticas del dashboard
type DashboardStats struct {
	TotalSecretarias  int64                  `json:"totalSecretarias"`
	TotalDependencias int64                  `json:"totalDependencias"`
	TotalEquipos      int64                  `json:"totalEquipos"`
	EquiposSinAsignar int64                  `json:"equiposSinAsignar"`
	UsuariosLibres    int64                  `json:"usuariosLibres"`
	EquiposPorEstado  []EstadoCount          `json:"equiposPorEstado"`
	EquiposPorTipo    []TipoCount            `json:"equiposPorTipo"`
	Secretarias       []SecretariaConEquipos `json:"secretarias"`
}

// EstadoCount conteo por estado
type EstadoCount struct {
	Estado   string `json:"estado"`
	Cantidad int    `json:"cantidad"`
}

// TipoCount conteo por tipo de dispositivo
type TipoCount struct {
	Tipo     string `json:"tipo"`
	Cantidad int    `json:"cantidad"`
}

// SecretariaConEquipos secretaría con su conteo de equipos
type SecretariaConEquipos struct {
	ID           uint                    `json:"ID"`
	Nombre       string                  `json:"Nombre"`
	Secretario   string                  `json:"Secretario"`
	Ubicacion    string                  `json:"Ubicacion"`
	TotalEquipos int                     `json:"totalEquipos"`
	Dependencias []DependenciaConEquipos `json:"dependencias"`
}

// DependenciaConEquipos dependencia con su conteo de equipos
type DependenciaConEquipos struct {
	ID           uint   `json:"ID"`
	Nombre       string `json:"Nombre"`
	TotalEquipos int    `json:"totalEquipos"`
}

// EquipoSinSecretaria representa un equipo sin secretaría asignada
type EquipoSinSecretaria struct {
	ID               uint    `json:"ID"`
	TipoDispositivo  string  `json:"TipoDispositivo"`
	PlacaInventario  string  `json:"PlacaInventario"`
	Marca            string  `json:"Marca"`
	Serial           string  `json:"Serial"`
	Modelo           string  `json:"Modelo"`
	Estado           string  `json:"Estado"`
	NombresApellidos *string `json:"NombresApellidos"`
	Cedula           *string `json:"Cedula"`
}

// UsuarioSinSecretaria representa un usuario responsable sin secretaría
type UsuarioSinSecretaria struct {
	ID               uint   `json:"ID"`
	NombresApellidos string `json:"NombresApellidos"`
	Cedula           string `json:"Cedula"`
	CorreoPersonal   string `json:"CorreoPersonal"`
	TipoVinculacion  string `json:"TipoVinculacion"`
	Celular          string `json:"Celular"`
	TotalEquipos     int    `json:"TotalEquipos"`
}

// SinSecretariaResponse respuesta con equipos y usuarios sin secretaría
type SinSecretariaResponse struct {
	Equipos  []EquipoSinSecretaria  `json:"equipos"`
	Usuarios []UsuarioSinSecretaria `json:"usuarios"`
}

// DashboardService interfaz del servicio de dashboard
type DashboardService interface {
	GetDashboardStats() (*DashboardStats, error)
	GetSinSecretaria() (*SinSecretariaResponse, error)
}

type dashboardService struct {
	db *gorm.DB
}

// NewDashboardService crea una nueva instancia del servicio
func NewDashboardService(db *gorm.DB) DashboardService {
	return &dashboardService{db: db}
}

func (s *dashboardService) GetDashboardStats() (*DashboardStats, error) {
	stats := &DashboardStats{}

	// Conteos totales usando los modelos GORM (resuelve nombres de tabla automáticamente)
	s.db.Model(&models.Secretaria{}).Count(&stats.TotalSecretarias)
	s.db.Model(&models.Dependencia{}).Count(&stats.TotalDependencias)
	s.db.Model(&models.Equipo{}).Count(&stats.TotalEquipos)

	// Equipos sin asignar: sin usuario responsable O cuyo usuario no tiene dependencia
	s.db.Raw(`
		SELECT COUNT(*) FROM equipos e
		WHERE e.deleted_at IS NULL
		AND (
			e.usuario_responsable_id IS NULL
			OR e.usuario_responsable_id NOT IN (
				SELECT ur.id FROM usuario_responsables ur
				WHERE ur.deleted_at IS NULL AND ur.dependencia_id IS NOT NULL
			)
		)
	`).Scan(&stats.EquiposSinAsignar)

	// Usuarios responsables libres (sin dependencia asignada)
	s.db.Raw(`
		SELECT COUNT(*) FROM usuario_responsables
		WHERE deleted_at IS NULL AND dependencia_id IS NULL
	`).Scan(&stats.UsuariosLibres)

	// Equipos por estado (1 query con JOIN)
	s.db.Raw(`
		SELECT es.nombre as estado, COUNT(e.id) as cantidad
		FROM equipos e
		JOIN estado_equipos es ON es.id = e.estado_equipo_id
		WHERE e.deleted_at IS NULL AND es.deleted_at IS NULL
		GROUP BY es.nombre
		ORDER BY cantidad DESC
	`).Scan(&stats.EquiposPorEstado)

	// Equipos por tipo de dispositivo (1 query)
	s.db.Raw(`
		SELECT tipo_dispositivo as tipo, COUNT(*) as cantidad
		FROM equipos
		WHERE deleted_at IS NULL AND tipo_dispositivo IS NOT NULL AND tipo_dispositivo != ''
		GROUP BY tipo_dispositivo
		ORDER BY cantidad DESC
	`).Scan(&stats.EquiposPorTipo)

	// Cargar secretarías con sus dependencias usando Preload de GORM
	var secretarias []models.Secretaria
	s.db.Preload("Dependencias").Find(&secretarias)

	// Cargar conteo de equipos por dependencia en una sola query
	type depEquipoCount struct {
		DependenciaID uint
		TotalEquipos  int
	}
	var depCounts []depEquipoCount
	s.db.Raw(`
		SELECT ur.dependencia_id, COUNT(e.id) as total_equipos
		FROM equipos e
		JOIN usuario_responsables ur ON ur.id = e.usuario_responsable_id
		WHERE e.deleted_at IS NULL AND ur.deleted_at IS NULL
		GROUP BY ur.dependencia_id
	`).Scan(&depCounts)

	// Crear mapa de dependencia_id → total_equipos para acceso rápido
	depEquiposMap := make(map[uint]int)
	for _, dc := range depCounts {
		depEquiposMap[dc.DependenciaID] = dc.TotalEquipos
	}

	// Armar estructura de secretarías con dependencias
	for _, sec := range secretarias {
		secConEquipos := SecretariaConEquipos{
			ID:         sec.ID,
			Nombre:     sec.Nombre,
			Secretario: sec.Secretario,
			Ubicacion:  sec.Ubicacion,
		}

		totalEquiposSec := 0
		for _, dep := range sec.Dependencias {
			totalEquiposDep := depEquiposMap[dep.ID]
			secConEquipos.Dependencias = append(secConEquipos.Dependencias, DependenciaConEquipos{
				ID:           dep.ID,
				Nombre:       dep.Nombre,
				TotalEquipos: totalEquiposDep,
			})
			totalEquiposSec += totalEquiposDep
		}

		secConEquipos.TotalEquipos = totalEquiposSec
		stats.Secretarias = append(stats.Secretarias, secConEquipos)
	}

	return stats, nil
}

func (s *dashboardService) GetSinSecretaria() (*SinSecretariaResponse, error) {
	response := &SinSecretariaResponse{}

	// Equipos sin secretaría: sin usuario responsable, o cuyo usuario no tiene dependencia
	s.db.Raw(`
		SELECT e.id, e.tipo_dispositivo, e.placa_inventario, e.marca, e.serial, e.modelo,
			es.nombre as estado,
			ur.nombres_apellidos, ur.cedula
		FROM equipos e
		LEFT JOIN estado_equipos es ON es.id = e.estado_equipo_id
		LEFT JOIN usuario_responsables ur ON ur.id = e.usuario_responsable_id AND ur.deleted_at IS NULL
		WHERE e.deleted_at IS NULL
		AND (
			e.usuario_responsable_id IS NULL
			OR ur.dependencia_id IS NULL
		)
		ORDER BY e.id DESC
	`).Scan(&response.Equipos)

	// Usuarios responsables sin dependencia (libres)
	s.db.Raw(`
		SELECT ur.id, ur.nombres_apellidos, ur.cedula, ur.correo_personal,
			ur.tipo_vinculacion, ur.celular,
			COUNT(e.id) as total_equipos
		FROM usuario_responsables ur
		LEFT JOIN equipos e ON e.usuario_responsable_id = ur.id AND e.deleted_at IS NULL
		WHERE ur.deleted_at IS NULL AND ur.dependencia_id IS NULL
		GROUP BY ur.id, ur.nombres_apellidos, ur.cedula, ur.correo_personal,
			ur.tipo_vinculacion, ur.celular
		ORDER BY ur.nombres_apellidos ASC
	`).Scan(&response.Usuarios)

	return response, nil
}
