package database

import (
	"gestor_api/src/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) *Database {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database := Database{db: db}
	database.migrate()
	return &database
}

func (d *Database) migrate() error {
	return d.db.AutoMigrate(
		&models.Celula{},
		&models.Encontro{},
		&models.MembroCelula{},
		&models.MembroCelulaEncontro{},
		&models.Culto{},
		&models.MetricasCulto{},
		&models.Turma{},
		&models.Aula{},
		&models.Pastor{},
		&models.Supervisor{},
		&models.Lider{},
		&models.Membro{},
		&models.Aluno{},
		&models.Rede{},
	)
}

// CELULAS
func (d *Database) GetCelulas() ([]models.Celula, error) {
	var celulas []models.Celula
	if err := d.db.Find(&celulas).Error; err != nil {
		return nil, err
	}
	return celulas, nil
}

func (d *Database) GetCelulaById(id uint) (models.Celula, error) {
	var celula models.Celula
	if err := d.db.First(&celula, id).Error; err != nil {
		return models.Celula{}, err
	}
	return celula, nil
}

func (d *Database) GetMembrosCelula(celulaId uint) ([]models.MembroCelula, error) {
	var membrosCelula []models.MembroCelula
	if err := d.db.Where("id_celula = ?", celulaId).Find(&membrosCelula).Error; err != nil {
		return nil, err
	}
	return membrosCelula, nil
}

func (d *Database) GetMembroCelulaById(id uint) (models.MembroCelula, error) {
	var membroCelula models.MembroCelula
	if err := d.db.First(&membroCelula, id).Error; err != nil {
		return models.MembroCelula{}, err
	}
	return membroCelula, nil
}

func (d *Database) AdicionarMembroCelula(membroCelula models.MembroCelula) (models.MembroCelula, error) {
	if err := d.db.Create(&membroCelula).Error; err != nil {
		return models.MembroCelula{}, err
	}
	return membroCelula, nil
}

func (d *Database) RemoverMembroCelula(id uint) error {
	if err := d.db.Delete(&models.MembroCelula{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateCelula(celula models.Celula) (models.Celula, error) {
	if err := d.db.Create(&celula).Error; err != nil {
		return models.Celula{}, err
	}
	return celula, nil
}

func (d *Database) UpdateCelula(celula models.Celula) (models.Celula, error) {
	if err := d.db.Save(&celula).Error; err != nil {
		return models.Celula{}, err
	}
	return celula, nil
}

func (d *Database) DeleteCelula(id uint) error {
	if err := d.db.Delete(&models.Celula{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ENCONTROS
func (d *Database) GetEncontros() ([]models.Encontro, error) {
	var encontros []models.Encontro
	if err := d.db.Find(&encontros).Error; err != nil {
		return nil, err
	}
	return encontros, nil
}

func (d *Database) GetEncontroByIdCelula(id uint) (models.Encontro, error) {
	var encontro models.Encontro
	if err := d.db.Where("id_celula = ?", id).First(&encontro).Error; err != nil {
		return models.Encontro{}, err
	}
	return encontro, nil
}

func (d *Database) CreateEncontro(encontro models.Encontro) (models.Encontro, error) {
	if err := d.db.Create(&encontro).Error; err != nil {
		return models.Encontro{}, err
	}
	return encontro, nil
}

func (d *Database) UpdateEncontro(encontro models.Encontro) (models.Encontro, error) {
	if err := d.db.Save(&encontro).Error; err != nil {
		return models.Encontro{}, err
	}
	return encontro, nil
}

func (d *Database) DeleteEncontro(id uint) error {
	if err := d.db.Delete(&models.Encontro{}, id).Error; err != nil {
		return err
	}
	return nil
}

// REDES
func (d *Database) GetRedes() ([]models.Rede, error) {
	var redes []models.Rede
	if err := d.db.Find(&redes).Error; err != nil {
		return nil, err
	}
	return redes, nil
}

func (d *Database) CreateRede(rede models.Rede) (models.Rede, error) {
	if err := d.db.Create(&rede).Error; err != nil {
		return models.Rede{}, err
	}
	return rede, nil
}

func (d *Database) UpdateRede(rede models.Rede) (models.Rede, error) {
	if err := d.db.Save(&rede).Error; err != nil {
		return models.Rede{}, err
	}
	return rede, nil
}

func (d *Database) DeleteRede(id uint) error {
	if err := d.db.Delete(&models.Rede{}, id).Error; err != nil {
		return err
	}
	return nil
}

// SUPERVISORES
func (d *Database) GetSupervisores() ([]models.Supervisor, error) {
	var supervisores []models.Supervisor
	if err := d.db.Find(&supervisores).Error; err != nil {
		return nil, err
	}
	return supervisores, nil
}

func (d *Database) CreateSupervisor(supervisor models.Supervisor) (models.Supervisor, error) {
	if err := d.db.Create(&supervisor).Error; err != nil {
		return models.Supervisor{}, err
	}
	return supervisor, nil
}

func (d *Database) UpdateSupervisor(supervisor models.Supervisor) (models.Supervisor, error) {
	if err := d.db.Save(&supervisor).Error; err != nil {
		return models.Supervisor{}, err
	}
	return supervisor, nil
}

func (d *Database) DeleteSupervisor(id uint) error {
	if err := d.db.Delete(&models.Supervisor{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetSupervisorById(id uint) (models.Supervisor, error) {
	var supervisor models.Supervisor
	if err := d.db.First(&supervisor, id).Error; err != nil {
		return models.Supervisor{}, err
	}
	return supervisor, nil
}
