package repository

import (
	"log"

	"gorm.io/gorm"
)

type ICommonDB interface {
	INationalityRepository
	IPartnerRepository
}

type commonDB struct {
	INationalityRepository
	IPartnerRepository
	db *gorm.DB
}

func (c *commonDB) CreateNewDB(partnerId string) error {
	err := c.db.Exec("CREATE DATABASE IF NOT EXISTS " + makeDatabaseName(partnerId)).Error
	if err != nil {
		log.Printf("db_common: Error while creating new database:%s, %v", partnerId, err)
		return err
	}
	return nil
}
