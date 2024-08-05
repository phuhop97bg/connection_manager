package repository

import (
	"hrm-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type IPartnerRepository interface {
	GetAll() ([]model.Partner, error)
	GetByPartnerId(partnerId string) (*model.Partner, error)
	Create(*model.Partner) error
}

type partnerRepository struct {
	db     *gorm.DB
	dbName string
}

// Create implements IPartnerRepository.
func (p *partnerRepository) Create(do *model.Partner) error {
	err := p.db.Table("partners").Create(do).Error
	if err != nil {
		log.Printf("db: %s Error while creating partner:%+v, %v", p.dbName, do, err)
		return err
	}
	return nil
}

// GetAll implements IPartnerRepository.
func (p *partnerRepository) GetAll() ([]model.Partner, error) {
	var partners []model.Partner
	err := p.db.Table("partners").Find(&partners).Error
	if err != nil {
		log.Printf("db: %s Error while getting all partners: %v", p.dbName, err)
		return nil, err
	}
	return partners, nil
}

// GetByID implements IPartnerRepository.
func (p *partnerRepository) GetByPartnerId(id string) (*model.Partner, error) {
	var partner model.Partner
	err := p.db.Table("partners").Where("partner_id = ?", id).First(&partner).Error
	if err != nil {
		log.Printf("db: %+v Error while getting partner by id:%+v, %v", p.dbName, id, err)
		return nil, err
	}
	return &partner, nil
}

func GetPartnerRepo(db *gorm.DB, dbName string) IPartnerRepository {
	return &partnerRepository{
		db:     db,
		dbName: dbName,
	}
}
