package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB_LOGIN := os.Getenv("DB_LOGIN")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	dsn := DB_LOGIN + ":" + DB_PASS + "@tcp(" + DB_HOST + ")/web-app_check-in?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:root@tcp(localhost:3306)/web-app_check-in?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	//database.AutoMigrate(&Book{})
	database.AutoMigrate(&Users{})
	database.AutoMigrate(&Admins{})

	DB = database
}
