package main

import (
	"log"
	"restapi/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("restApi.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подкллючения к БД")
	}
	db.AutoMigrate(&model.User{})

}
