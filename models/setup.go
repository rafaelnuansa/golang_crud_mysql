package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Ambil variabel-variabel dari lingkungan
	hostname := os.Getenv("HOSTNAME")
	dbport := os.Getenv("DBPORT")
	dbusername := os.Getenv("DBUSERNAME")
	dbname := os.Getenv("DBNAME")
	dbpassword := os.Getenv("DBPASSWORD")

	// Buat string koneksi
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpassword, hostname, dbport, dbname)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Post{})
	DB = database
}
