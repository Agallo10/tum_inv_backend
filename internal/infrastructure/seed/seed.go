package seed

import (
	"log"
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// Seeder contiene la instancia de la base de datos para realizar el seeding
type Seeder struct {
	DB *gorm.DB
}

// NewSeeder crea una nueva instancia del seeder
func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{DB: db}
}

// SeedAll ejecuta todas las funciones de seeding
func (s *Seeder) SeedAll() error {
	log.Println("Iniciando proceso de seeding...")

	// Ejecutar seeds en orden de dependencias
	if err := s.SeedSecretarias(); err != nil {
		return err
	}

	if err := s.SeedDependencias(); err != nil {
		return err
	}

	if err := s.SeedEstadosEquipo(); err != nil {
		return err
	}

	log.Println("Proceso de seeding completado exitosamente")
	return nil
}

// SeedSecretarias inserta las secretarías iniciales
func (s *Seeder) SeedSecretarias() error {
	log.Println("Insertando secretarías iniciales...")

	secretarias := []models.Secretaria{
		{
			Nombre:      "Secretaría de Educación",
			Descripcion: "Secretaría encargada de la administración y supervisión del sistema educativo municipal",
			Ubicacion:   "Edificio Central - Piso 3",
			Secretario:  "Alexis Erazo",
			Telefono:    "",
		},
		{
			Nombre:      "Secretaría de Salud",
			Descripcion: "Secretaría responsable de la gestión de servicios de salud pública municipal",
			Ubicacion:   "Centro Administrativo - Piso 2",
			Secretario:  "Dr. Carlos Alberto Hernández",
			Telefono:    "123-456-7891",
		},
		{
			Nombre:      "Secretaría de Infraestructura",
			Descripcion: "Secretaría encargada del desarrollo y mantenimiento de la infraestructura municipal",
			Ubicacion:   "Edificio Técnico - Piso 1",
			Secretario:  "Ing. Ana Patricia López",
			Telefono:    "123-456-7892",
		},
		{
			Nombre:      "Secretaría de Cultura y Deportes",
			Descripcion: "Secretaría responsable de promover actividades culturales y deportivas",
			Ubicacion:   "Casa de la Cultura - Piso 2",
			Secretario:  "Lic. Roberto José Martínez",
			Telefono:    "123-456-7893",
		},
		{
			Nombre:      "Secretaría de Tecnología e Innovación",
			Descripcion: "Secretaría encargada de la modernización tecnológica del municipio",
			Ubicacion:   "Centro de Innovación - Piso 4",
			Secretario:  "Ing. Laura Beatriz Gómez",
			Telefono:    "123-456-7894",
		},
	}

	for _, secretaria := range secretarias {
		// Verificar si ya existe la secretaría
		var existing models.Secretaria
		if err := s.DB.Where("nombre = ?", secretaria.Nombre).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// No existe, crear nueva
				if err := s.DB.Create(&secretaria).Error; err != nil {
					log.Printf("Error al crear secretaría %s: %v", secretaria.Nombre, err)
					return err
				}
				log.Printf("Secretaría '%s' creada exitosamente", secretaria.Nombre)
			} else {
				return err
			}
		} else {
			log.Printf("Secretaría '%s' ya existe, omitiendo...", secretaria.Nombre)
		}
	}

	return nil
}

// SeedDependencias inserta las dependencias iniciales para cada secretaría
func (s *Seeder) SeedDependencias() error {
	log.Println("Insertando dependencias iniciales...")

	// Obtener las secretarías para obtener sus IDs
	var secretarias []models.Secretaria
	if err := s.DB.Find(&secretarias).Error; err != nil {
		return err
	}

	// Crear un mapa para facilitar la búsqueda por nombre
	secretariaMap := make(map[string]uint)
	for _, sec := range secretarias {
		secretariaMap[sec.Nombre] = sec.ID
	}

	dependencias := []struct {
		SecretariaNombre    string
		Nombre              string
		Descripcion         string
		UbicacionOficina    string
		JefeOficina         string
		CorreoInstitucional string
		Telefono            string
	}{
		// Dependencias de Secretaría de Educación
		{
			SecretariaNombre:    "Secretaría de Educación",
			Nombre:              "Casa de la Cultura",
			Descripcion:         "Dependencia encargada de supervisar y mejorar la calidad educativa",
			UbicacionOficina:    "Edificio Central - Oficina 301",
			JefeOficina:         "Lic. Carmen Rosa Díaz",
			CorreoInstitucional: "calidad.educativa@municipio.gov.co",
			Telefono:            "123-456-7801",
		},
		{
			SecretariaNombre:    "Secretaría de Educación",
			Nombre:              "Dirección de Infraestructura Educativa",
			Descripcion:         "Dependencia responsable del mantenimiento de instalaciones educativas",
			UbicacionOficina:    "Edificio Central - Oficina 302",
			JefeOficina:         "Ing. Miguel Ángel Torres",
			CorreoInstitucional: "infraestructura.educativa@municipio.gov.co",
			Telefono:            "123-456-7802",
		},
		// Dependencias de Secretaría de Salud
		{
			SecretariaNombre:    "Secretaría de Salud",
			Nombre:              "Dirección de Atención Primaria",
			Descripcion:         "Dependencia encargada de la atención primaria en salud",
			UbicacionOficina:    "Centro Administrativo - Oficina 201",
			JefeOficina:         "Dr. Sandra Milena Vargas",
			CorreoInstitucional: "atencion.primaria@salud.municipio.gov.co",
			Telefono:            "123-456-7811",
		},
		{
			SecretariaNombre:    "Secretaría de Salud",
			Nombre:              "Dirección de Vigilancia Epidemiológica",
			Descripcion:         "Dependencia responsable del control epidemiológico municipal",
			UbicacionOficina:    "Centro Administrativo - Oficina 202",
			JefeOficina:         "Dr. Fernando Javier Ruiz",
			CorreoInstitucional: "epidemiologia@salud.municipio.gov.co",
			Telefono:            "123-456-7812",
		},
		// Dependencias de Secretaría de Infraestructura
		{
			SecretariaNombre:    "Secretaría de Infraestructura",
			Nombre:              "Dirección de Obras Públicas",
			Descripcion:         "Dependencia encargada de la ejecución de obras públicas",
			UbicacionOficina:    "Edificio Técnico - Oficina 101",
			JefeOficina:         "Ing. Pablo César Moreno",
			CorreoInstitucional: "obras.publicas@infraestructura.municipio.gov.co",
			Telefono:            "123-456-7821",
		},
		{
			SecretariaNombre:    "Secretaría de Infraestructura",
			Nombre:              "Dirección de Servicios Públicos",
			Descripcion:         "Dependencia responsable de la supervisión de servicios públicos",
			UbicacionOficina:    "Edificio Técnico - Oficina 102",
			JefeOficina:         "Ing. Gloria Esperanza Silva",
			CorreoInstitucional: "servicios.publicos@infraestructura.municipio.gov.co",
			Telefono:            "123-456-7822",
		},
		// Dependencias de Secretaría de Cultura y Deportes
		{
			SecretariaNombre:    "Secretaría de Cultura y Deportes",
			Nombre:              "Dirección de Promoción Cultural",
			Descripcion:         "Dependencia encargada de promover actividades culturales",
			UbicacionOficina:    "Casa de la Cultura - Oficina 201",
			JefeOficina:         "Lic. Claudia Patricia Ramírez",
			CorreoInstitucional: "promocion.cultural@cultura.municipio.gov.co",
			Telefono:            "123-456-7831",
		},
		{
			SecretariaNombre:    "Secretaría de Cultura y Deportes",
			Nombre:              "Dirección de Deportes y Recreación",
			Descripcion:         "Dependencia responsable de actividades deportivas y recreativas",
			UbicacionOficina:    "Casa de la Cultura - Oficina 202",
			JefeOficina:         "Lic. Andrés Felipe Castro",
			CorreoInstitucional: "deportes.recreacion@cultura.municipio.gov.co",
			Telefono:            "123-456-7832",
		},
		// Dependencias de Secretaría de Tecnología e Innovación
		{
			SecretariaNombre:    "Secretaría de Tecnología e Innovación",
			Nombre:              "Dirección de Sistemas de Información",
			Descripcion:         "Dependencia encargada de los sistemas de información municipales",
			UbicacionOficina:    "Centro de Innovación - Oficina 401",
			JefeOficina:         "Ing. Diana Carolina Pérez",
			CorreoInstitucional: "sistemas.informacion@tecnologia.municipio.gov.co",
			Telefono:            "123-456-7841",
		},
		{
			SecretariaNombre:    "Secretaría de Tecnología e Innovación",
			Nombre:              "Dirección de Innovación Digital",
			Descripcion:         "Dependencia responsable de proyectos de innovación digital",
			UbicacionOficina:    "Centro de Innovación - Oficina 402",
			JefeOficina:         "Ing. Julián Camilo Mendoza",
			CorreoInstitucional: "innovacion.digital@tecnologia.municipio.gov.co",
			Telefono:            "123-456-7842",
		},
	}

	for _, dep := range dependencias {
		secretariaID, exists := secretariaMap[dep.SecretariaNombre]
		if !exists {
			log.Printf("Secretaría '%s' no encontrada para dependencia '%s'", dep.SecretariaNombre, dep.Nombre)
			continue
		}

		dependencia := models.Dependencia{
			SecretariaID:        secretariaID,
			Nombre:              dep.Nombre,
			Descripcion:         dep.Descripcion,
			UbicacionOficina:    dep.UbicacionOficina,
			JefeOficina:         dep.JefeOficina,
			CorreoInstitucional: dep.CorreoInstitucional,
			Telefono:            dep.Telefono,
		}

		// Verificar si ya existe la dependencia
		var existing models.Dependencia
		if err := s.DB.Where("nombre = ? AND secretaria_id = ?", dependencia.Nombre, dependencia.SecretariaID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// No existe, crear nueva
				if err := s.DB.Create(&dependencia).Error; err != nil {
					log.Printf("Error al crear dependencia %s: %v", dependencia.Nombre, err)
					return err
				}
				log.Printf("Dependencia '%s' creada exitosamente en secretaría '%s'", dependencia.Nombre, dep.SecretariaNombre)
			} else {
				return err
			}
		} else {
			log.Printf("Dependencia '%s' ya existe en secretaría '%s', omitiendo...", dependencia.Nombre, dep.SecretariaNombre)
		}
	}

	return nil
}

// SeedEstadosEquipo inserta los estados de equipo básicos
func (s *Seeder) SeedEstadosEquipo() error {
	log.Println("Insertando estados de equipo iniciales...")

	estados := []models.EstadoEquipo{
		{
			Nombre:      "Activo",
			Descripcion: "Equipo en funcionamiento normal",
			Activo:      true,
		},
		{
			Nombre:      "Inactivo",
			Descripcion: "Equipo fuera de servicio temporalmente",
			Activo:      false,
		},
		{
			Nombre:      "En Mantenimiento",
			Descripcion: "Equipo en proceso de mantenimiento",
			Activo:      false,
		},
		{
			Nombre:      "Dañado",
			Descripcion: "Equipo con fallas que requieren reparación",
			Activo:      false,
		},
		{
			Nombre:      "Dado de Baja",
			Descripcion: "Equipo retirado definitivamente del inventario",
			Activo:      false,
		},
	}

	for _, estado := range estados {
		// Verificar si ya existe el estado
		var existing models.EstadoEquipo
		if err := s.DB.Where("nombre = ?", estado.Nombre).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// No existe, crear nuevo
				if err := s.DB.Create(&estado).Error; err != nil {
					log.Printf("Error al crear estado de equipo %s: %v", estado.Nombre, err)
					return err
				}
				log.Printf("Estado de equipo '%s' creado exitosamente", estado.Nombre)
			} else {
				return err
			}
		} else {
			log.Printf("Estado de equipo '%s' ya existe, omitiendo...", estado.Nombre)
		}
	}

	return nil
}
