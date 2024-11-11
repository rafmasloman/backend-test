package config

import (
	"backend/api/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(localhost:3306)/backend_technical_user"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
        fmt.Printf("Failed to connect to database: %v", err)
    }

	err = DB.AutoMigrate(&models.Customer{})
    fmt.Println("Database connected successfully!")
}