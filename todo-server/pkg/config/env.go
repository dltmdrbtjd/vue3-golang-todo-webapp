package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUserName string
	DBPassword string
	DBPort     string
	DBName     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DBUserName = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
}
