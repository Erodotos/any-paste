package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	// load .env file
	err := godotenv.Load(".env")
	fmt.Println("loading")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}