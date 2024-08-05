package repository

import (
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

type IPrivateDB interface {
	IDepartmentRepository
	IProfileRepository
}

type privateDB struct {
	IDepartmentRepository
	IProfileRepository
	db *gorm.DB
}

func (p *privateDB) CreateTables() error {
	sqlFile := "migrations/init.sql"
	sqlContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}
	queries := strings.Split(string(sqlContent), ";")

	// Thực thi từng câu lệnh SQL
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query != "" {
			err = p.db.Exec(query).Error
			if err != nil {
				log.Fatalf("Failed to execute query: %s, error: %v", query, err)
			}
		}
	}
	return err
}
