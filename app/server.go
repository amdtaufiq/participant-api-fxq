package app

import (
	"fmt"
	"log"
	"os"
	"participant-api/app/helper"
	"participant-api/config"

	"github.com/joho/godotenv"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	environment := helper.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		DB_USER := os.Getenv("DATABASE_USERNAME")
		DB_PASSWORD := os.Getenv("DATABASE_PASSWORD")
		DB_HOST := os.Getenv("DATABASE_HOST")
		DB_PORT := os.Getenv("DATABASE_PORT")
		DB_NAME := os.Getenv("DATABASE_NAME")

		config.Initialize(DB_USER, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME, environment)

	} else {

		DB_HOST := "localhost"
		DB_USER := "JCC"
		DB_PASSWORD := "JCC123"
		DB_NAME := "participant_db"
		DB_PORT := "5432"

		config.Initialize(DB_USER, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME, environment)
	}

}
