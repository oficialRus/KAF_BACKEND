package database

import (
	"log"
	"restapi/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	DB, err = gorm.Open(sqlite.Open("restApi.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подкллючения к БД")
	}
	DB.AutoMigrate(&model.User{})

}
