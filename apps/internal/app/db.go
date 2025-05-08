package app

import "gorm.io/gorm"
import "gorm.io/driver/postgres"

var DB *gorm.DB

func InitDB(config *Config) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
