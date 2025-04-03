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
	db.AutoMigrate(&model.User{}, &model.Chat{}, model.Message{})

}

/*go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/joho/godotenv
*/
