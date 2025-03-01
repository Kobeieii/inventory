package database

import (
	"inventory/internal/app/adapters/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&repositories.ProductModel{}); err != nil {
		return nil, err
	}
	return db, nil
}
