package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// connectDb
func InitializeDatabase() {

	var err error
	p := os.Getenv("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"))
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	autoMigrate()
}

func autoMigrate() {
	DB.AutoMigrate(&models.Post{})
}
