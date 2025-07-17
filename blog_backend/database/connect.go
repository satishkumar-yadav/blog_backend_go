package database

import (
	"log"

	"blog/config"
	"blog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.DSN

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Database connected successfully")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

}
