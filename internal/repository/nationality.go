package repository

import (
	"hrm-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type INationalityRepository interface {
	GetAllNationality() ([]model.Nationality, error)
	GeNationalitytByID(id int) (*model.Nationality, error)
}

type NationalityRepository struct {
	db     *gorm.DB
	dbName string
}

func GetNationalityRepo(db *gorm.DB, dbName string) INationalityRepository {
	return &NationalityRepository{
		db:     db,
		dbName: dbName,
	}
}

// GetAll implements INationalityRepository.
func (n *NationalityRepository) GetAllNationality() ([]model.Nationality, error) {
	var nationalities []model.Nationality
	err := n.db.Table("nationalities").Find(&nationalities).Error
	if err != nil {
		log.Printf("db: %s Error while getting all nationalities: %v", n.dbName, err)
		return nil, err
	}
	return nationalities, nil
}

// GetByID implements INationalityRepository.
func (n *NationalityRepository) GeNationalitytByID(id int) (*model.Nationality, error) {
	var nationality model.Nationality
	err := n.db.Table("nationalities").Where("id = ?", id).First(&nationality).Error
	if err != nil {
		log.Printf("db: %s Error while getting nationality by id:%d, %v", n.dbName, id, err)
		return nil, err
	}
	return &nationality, nil
}
