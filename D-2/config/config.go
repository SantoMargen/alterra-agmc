package config

import (
	"fmt"
	"log"
	"os"

	"D-2/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		dbuser, dbpass, dbhost, dbport, dbname, "utf8mb4", "true",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("success connect to DB")
	InitialMigration(DB)
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
	log.Println("Success auto migrate")
}
