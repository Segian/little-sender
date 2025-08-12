package config

import (
	"github.com/Segian/little-sender/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	migrateErr := migrate()
	if err != nil {
		return err
	}
	if migrateErr != nil {
		return migrateErr
	}
	return nil
}

func migrate() error {
	DB.AutoMigrate(&model.BodyModel{})
	DB.AutoMigrate(&model.EndpointModel{})
	return nil
}
