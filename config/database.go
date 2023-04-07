package config

import (
	"fmt"

	"go-fiber-crud/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	helper.ErrorPanic(err)

	fmt.Println("Connected successfully to the database!")
	return db
}
