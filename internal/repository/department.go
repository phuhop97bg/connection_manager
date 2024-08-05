package repository

import (
	"hrm-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	GetAll() ([]model.Department, error)
	GetByID(id int) (model.Department, error)
	Create(*model.Department) error
}

type departmentRepository struct {
	db     *gorm.DB
	dbName string
}

func GetDepartmentRepo(db *gorm.DB, dbName string) IDepartmentRepository {
	return &departmentRepository{
		db:     db,
		dbName: dbName,
	}
}

// Create implements IDepartmentRepository.
func (d *departmentRepository) Create(do *model.Department) error {
	err := d.db.Table("departments").Create(do).Error
	if err != nil {
		log.Printf("db: %s Error while creating department:%+v, %v", d.dbName, do, err)
		return err
	}
	return nil
}

// GetAll implements IDepartmentRepository.
func (d *departmentRepository) GetAll() ([]model.Department, error) {
	var departments []model.Department
	err := d.db.Table("departments").Find(&departments).Error
	if err != nil {
		log.Printf("db: %s Error while getting all departments: %v", d.dbName, err)
		return nil, err
	}
	return departments, nil
}

// GetByID implements IDepartmentRepository.
func (d *departmentRepository) GetByID(id int) (model.Department, error) {
	var department model.Department
	err := d.db.Table("departments").Where("id = ?", id).First(&department).Error
	if err != nil {
		log.Printf("db: %s Error while getting department by id:%d, %v", d.dbName, id, err)
		return model.Department{}, err
	}
	return department, nil
}
