package bean

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB(config *Config) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DBUrl()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
