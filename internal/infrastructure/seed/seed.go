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
	if err := s.SeedUsuarios(); err != nil {
		return err
	}

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
			Nombre:      "Secretaría de General",
			Descripcion: "Secretaría encargada de contratación",
			Ubicacion:   "Alcaldía, Edificio Central - Piso 3",
			Secretario:  "Marcos Castillo",
			Telefono:    "",
		},
		{
			Nombre:      "Secretaría de Salud",
			Descripcion: "Secretaría responsable de la gestión de servicios de salud pública municipal",
			Ubicacion:   "Centro Administrativo - Piso 2",
			Secretario:  "Maricel Rodriguez Ortega",
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

		// Dependencias de Secretaria General

		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Talento Humano",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con talento humano",
			UbicacionOficina:    "",
			JefeOficina:         "Jehyms Silva Vallecilla",
			CorreoInstitucional: "talentohumano@tumaco-narino.gov.co",
			Telefono:            "3183653577",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Contratación",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con contratación",
			UbicacionOficina:    "",
			JefeOficina:         "Javier Mayolo",
			CorreoInstitucional: "juridica@tumaco-narino.gov.co",
			Telefono:            "3188603186",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Archivo Central",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con archivo central",
			UbicacionOficina:    "",
			JefeOficina:         "Víctor Hugo Holguín",
			CorreoInstitucional: "victorhholguin@hotmail.com",
			Telefono:            "3147186598",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Sistemas",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con sistemas",
			UbicacionOficina:    "",
			JefeOficina:         "Francisco Rodríguez",
			CorreoInstitucional: "oficinadesistema@tumaco-narino.gov.co",
			Telefono:            "3208371910",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "CIS",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con cis",
			UbicacionOficina:    "",
			JefeOficina:         "Jorge Torres",
			CorreoInstitucional: "jorgetorres201800@gmail.com",
			Telefono:            "3102366002",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Gestión Documental",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con gestión documental",
			UbicacionOficina:    "",
			JefeOficina:         "Dolores Castillo",
			CorreoInstitucional: "gestiondocumental@tumaco-narino.gov.co",
			Telefono:            "3137540974",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Almacén",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con almacén",
			UbicacionOficina:    "",
			JefeOficina:         "Holger Burbano",
			CorreoInstitucional: "almacen@tumaco-narino.gov.co",
			Telefono:            "3153812112",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Atención al Ciudadano",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con atención al ciudadano",
			UbicacionOficina:    "",
			JefeOficina:         "Milton Quiñonez",
			CorreoInstitucional: "atencionalciudadano@tumaco-narino.gov.co",
			Telefono:            "3153812112",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Seguridad y Salud en el Trabajo",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con seguridad y salud en el trabajo",
			UbicacionOficina:    "",
			JefeOficina:         "Kleiver Vidal",
			CorreoInstitucional: "sstumaco2020@gmail.com",
			Telefono:            "3105120462",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Prensa",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con prensa",
			UbicacionOficina:    "",
			JefeOficina:         "Jesús Enrique",
			CorreoInstitucional: "comunicaciones@tumaco-narino.gov.co",
			Telefono:            "3007076578",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Control Interno",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con control interno",
			UbicacionOficina:    "",
			JefeOficina:         "Jerson Valencia",
			CorreoInstitucional: "",
			Telefono:            "",
		},
		{
			SecretariaNombre:    "Secretaría de General",
			Nombre:              "Gestión del Riesgo de Desastres",
			Descripcion:         "Dependencia encargada de las funciones relacionadas con gestión del riesgo de desastres",
			UbicacionOficina:    "",
			JefeOficina:         "Luz Mirian Quiñones",
			CorreoInstitucional: "",
			Telefono:            "",
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

// SeedUsuarios inserta los usuarios iniciales del sistema
func (s *Seeder) SeedUsuarios() error {
	log.Println("Insertando usuarios iniciales...")

	usuarios := []models.Usuario{
		{
			Nombre:   "Administrador",
			Apellido: "Sistema",
			Email:    "admin@municipio.gov.co",
			Username: "admin",
			Password: "admin123", // Se hasheará antes de guardar
			Rol:      "admin",
			Activo:   true,
		},
		{
			Nombre:   "Técnico",
			Apellido: "Soporte",
			Email:    "tecnico@municipio.gov.co",
			Username: "tecnico",
			Password: "tecnico123", // Se hasheará antes de guardar
			Rol:      "tecnico",
			Activo:   true,
		},
	}

	for _, usuario := range usuarios {
		// Verificar si ya existe el usuario
		var existing models.Usuario
		if err := s.DB.Where("username = ?", usuario.Username).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Hashear la contraseña antes de guardar
				if err := usuario.HashPassword(); err != nil {
					log.Printf("Error al hashear contraseña para usuario %s: %v", usuario.Username, err)
					return err
				}

				// No existe, crear nuevo
				if err := s.DB.Create(&usuario).Error; err != nil {
					log.Printf("Error al crear usuario %s: %v", usuario.Username, err)
					return err
				}
				log.Printf("Usuario '%s' con rol '%s' creado exitosamente", usuario.Username, usuario.Rol)
			} else {
				return err
			}
		} else {
			log.Printf("Usuario '%s' ya existe, omitiendo...", usuario.Username)
		}
	}

	return nil
}
