package Lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"laundry/Config"
)

var DB *gorm.DB

func initDB() {
	var err error = nil
	DB, err = gorm.Open(postgres.Open(Config.DatabaseConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
}
