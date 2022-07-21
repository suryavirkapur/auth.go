package database

import (
	"go-auth-server/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to DB.")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
