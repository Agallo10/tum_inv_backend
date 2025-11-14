package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"

	"tum_inv_backend/internal/domain/models/dto"
)

// EquipoRepository define las operaciones del repositorio para Equipo
type EquipoRepository interface {
	Create(equipo *models.Equipo) error
	FindByID(id uint) (*models.Equipo, error)
	Update(equipo *models.Equipo) error
	Delete(id uint) error
	FindAll() ([]models.Equipo, error)
	FindByDependenciaID(dependenciaID uint) ([]models.Equipo, error)
	FindEquiUsuDepByID(id uint) (dto.EquipoConResponsableDTO, error)
	FindAllEquiposDetalle() ([]dto.EquipoConResponsableDTO, error)
}

// equipoRepository implementa EquipoRepository
type equipoRepository struct {
	db *gorm.DB
}

// NewEquipoRepository crea una nueva instancia de EquipoRepository
func NewEquipoRepository(db *gorm.DB) EquipoRepository {
	return &equipoRepository{db: db}
}

// Create crea un nuevo equipo en la base de datos
func (r *equipoRepository) Create(equipo *models.Equipo) error {
	return r.db.Create(equipo).Error
}

// FindByID busca un equipo por su ID
func (r *equipoRepository) FindByID(id uint) (*models.Equipo, error) {
	var equipo models.Equipo
	err := r.db.Preload("UsuarioResponsable").
		Preload("Perifericos").
		Preload("HardwareInterno").
		Preload("Software").
		Preload("ConfiguracionRed").
		Preload("UsuariosSistema").
		Preload("AccesosRemotos").
		Preload("Backups").
		Preload("Reportes").
		First(&equipo, id).Error
	if err != nil {
		return nil, err
	}
	return &equipo, nil
}

// Update actualiza un equipo existente
func (r *equipoRepository) Update(equipo *models.Equipo) error {
	return r.db.Save(equipo).Error
}

// Delete elimina un equipo por su ID
func (r *equipoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Equipo{}, id).Error
}

// FindAll retorna todos los equipos
func (r *equipoRepository) FindAll() ([]models.Equipo, error) {
	var equipos []models.Equipo
	err := r.db.Find(&equipos).Error
	return equipos, err
}

// FindByDependenciaID retorna todos los equipos de una dependencia
func (r *equipoRepository) FindByDependenciaID(dependenciaID uint) ([]models.Equipo, error) {
	var equipos []models.Equipo
	err := r.db.Raw(`
  SELECT e.*
  FROM usuario_responsables ur
  JOIN equipos e ON ur.id = e.usuario_responsable_id
  WHERE ur.dependencia_id = ?
`, dependenciaID).
		Preload("UsuarioResponsable").
		Preload("Perifericos").
		Preload("HardwareInterno").
		Preload("Software").
		Preload("ConfiguracionRed").
		Preload("UsuariosSistema").
		Preload("AccesosRemotos").
		Preload("Backups").
		Preload("Reportes").
		Preload("EstadoEquipo").
		Scan(&equipos).Error
	return equipos, err
}

func (r *equipoRepository) FindEquiUsuDepByID(equipoID uint) (dto.EquipoConResponsableDTO, error) {
	var equipo dto.EquipoConResponsableDTO
	err := r.db.Raw(`
        SELECT e.marca, e.modelo, e.observaciones_generales, 
		e.placa_inventario, e.serial, e.tipo_dispositivo, e.fecha_diligenciamiento,
		ur.nombres_apellidos, ur.cedula, d.ubicacion_oficina, es.nombre as Estado
        FROM equipos e 
        JOIN usuario_responsables ur ON ur.id = e.usuario_responsable_id
        JOIN dependencia d ON d.id = ur.dependencia_id
        JOIN estado_equipos es ON es.id = e.estado_equipo_id
        WHERE e.id = ?`, equipoID).Scan(&equipo).Error
	return equipo, err
}
func (r *equipoRepository) FindAllEquiposDetalle() ([]dto.EquipoConResponsableDTO, error) {
	var equipos []dto.EquipoConResponsableDTO
	err := r.db.Raw(`
        SELECT e.marca, e.modelo, e.observaciones_generales, 
		e.placa_inventario, e.serial, e.tipo_dispositivo, e.fecha_diligenciamiento,
		ur.nombres_apellidos, ur.cedula, d.ubicacion_oficina, es.nombre as Estado
        FROM equipos e 
        JOIN usuario_responsables ur ON ur.id = e.usuario_responsable_id
        JOIN dependencia d ON d.id = ur.dependencia_id
        JOIN estado_equipos es ON es.id = e.estado_equipo_id`).Scan(&equipos).Error
	return equipos, err
}

// // FindByDependenciaID retorna todos los equipos de una dependencia
// func (r *equipoRepository) FindEquiUsuDepByID(equipoID uint) (models.Equipo, error) {
// 	var equipo models.Equipo
// 	err := r.db.
// 		Preload("UsuarioResponsable.Dependencia").
// 		First(&equipo, equipoID).Error
// 	return equipo, err
// }

// // FindByDependenciaID retorna todos los equipos de una dependencia
// func (r *equipoRepository) FindByDependenciaID(dependenciaID uint) ([]models.Equipo, error) {
// 	var equipos []models.Equipo
// 	err := r.db.Where("dependencia_id = ?", dependenciaID).
// 		Preload("UsuarioResponsable").
// 		Preload("Perifericos").
// 		Preload("HardwareInterno").
// 		Preload("Software").
// 		Preload("ConfiguracionRed").
// 		Preload("UsuariosSistema").
// 		Preload("AccesosRemotos").
// 		Preload("Backups").
// 		Preload("Reportes").
// 		Find(&equipos).Error
// 	return equipos, err
// }
