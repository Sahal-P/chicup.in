package db

import (
	"chicup/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " port=" + config.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
