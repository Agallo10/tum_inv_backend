package repositories

import (
	"tum_inv_backend/internal/domain/models"

	"gorm.io/gorm"
)

// PerifericoRepository define las operaciones del repositorio para Periferico
type PerifericoRepository interface {
	Create(periferico *models.Periferico) error
	FindByID(id uint) (*models.Periferico, error)
	Update(periferico *models.Periferico) error
	Delete(id uint) error
	FindAll() ([]models.Periferico, error)
	FindByEquipoID(equipoID uint) ([]models.Periferico, error)
	FindSinEquipo() ([]models.Periferico, error)
	AsignarEquipo(perifericoID uint, equipoID *uint) error
}

// perifericoRepository implementa PerifericoRepository
type perifericoRepository struct {
	db *gorm.DB
}

// NewPerifericoRepository crea una nueva instancia de PerifericoRepository
func NewPerifericoRepository(db *gorm.DB) PerifericoRepository {
	return &perifericoRepository{db: db}
}

// Create crea un nuevo periférico en la base de datos
func (r *perifericoRepository) Create(periferico *models.Periferico) error {
	return r.db.Create(periferico).Error
}

// FindByID busca un periférico por su ID
func (r *perifericoRepository) FindByID(id uint) (*models.Periferico, error) {
	var periferico models.Periferico
	err := r.db.First(&periferico, id).Error
	if err != nil {
		return nil, err
	}
	return &periferico, nil
}

// Update actualiza un periférico existente
func (r *perifericoRepository) Update(periferico *models.Periferico) error {
	return r.db.Save(periferico).Error
}

// Delete elimina un periférico por su ID
func (r *perifericoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Periferico{}, id).Error
}

// FindAll retorna todos los periféricos
func (r *perifericoRepository) FindAll() ([]models.Periferico, error) {
	var perifericos []models.Periferico
	err := r.db.Find(&perifericos).Error
	return perifericos, err
}

// FindByEquipoID retorna todos los periféricos asociados a un equipo
func (r *perifericoRepository) FindByEquipoID(equipoID uint) ([]models.Periferico, error) {
	var perifericos []models.Periferico
	err := r.db.Where("equipo_id = ?", equipoID).Find(&perifericos).Error
	return perifericos, err
}

// FindSinEquipo retorna todos los periféricos sin equipo asignado
func (r *perifericoRepository) FindSinEquipo() ([]models.Periferico, error) {
	var perifericos []models.Periferico
	err := r.db.Where("equipo_id IS NULL").Find(&perifericos).Error
	return perifericos, err
}

// AsignarEquipo actualiza solo el EquipoID de un periférico
func (r *perifericoRepository) AsignarEquipo(perifericoID uint, equipoID *uint) error {
	return r.db.Model(&models.Periferico{}).Where("id = ?", perifericoID).Update("equipo_id", equipoID).Error
}
