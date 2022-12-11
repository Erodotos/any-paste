package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"backend/config"
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// connectDb
func init() {

	var err error
	p := config.Config("DB_PORT")
	// fmt.Println( config.Config("DB_PORT"))
	// p := os.Getenv("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"))
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