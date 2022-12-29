package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	//dsn := "host=localhost user=son password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(sqlite.Open("sonhs.db"), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Person{})
	if err != nil {
		return
	}
	DB = db
}
