package database

import (
	"fmt"

	"github.com/anandawira/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:example@tcp(127.0.0.1:3306)/go-auth?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("cannot connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
