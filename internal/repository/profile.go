package repository

import (
	"hrm-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type IProfileRepository interface {
	GetAllProfile() ([]model.Profile, error)
	GetProfileByID(id int) (*model.Profile, error)
	CreateProfile(*model.Profile) error
}

type profileRepository struct {
	db     *gorm.DB
	dbName string
}

// Create implements IProfileRepository.
func (p *profileRepository) CreateProfile(do *model.Profile) error {
	err := p.db.Table("profiles").Create(do).Error
	if err != nil {
		log.Printf("db: %s Error while creating profile:%+v, %v", p.dbName, do, err)
		return err
	}
	return nil
}

// GetAll implements IProfileRepository.
func (p *profileRepository) GetAllProfile() ([]model.Profile, error) {
	var profiles []model.Profile
	err := p.db.Table("profiles").Find(&profiles).Error
	if err != nil {
		log.Printf("db: %s Error while getting all profiles: %v", p.dbName, err)
		return nil, err
	}
	return profiles, nil
}

// GetByID implements IProfileRepository.
func (p *profileRepository) GetProfileByID(id int) (*model.Profile, error) {
	var profile model.Profile
	err := p.db.Table("profiles").Where("id = ?", id).First(&profile).Error
	if err != nil {
		log.Printf("db: %s Error while getting profile by id:%d, %v", p.dbName, id, err)
		return nil, err
	}
	return &profile, nil
}

func GetProfileRepo(db *gorm.DB, dbName string) IProfileRepository {
	return &profileRepository{
		db:     db,
		dbName: dbName,
	}
}
