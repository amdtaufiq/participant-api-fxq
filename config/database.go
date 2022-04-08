package config

import (
	"fmt"
	"participant-api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize(DbUser, DbPassword, DbPort, DbHost, DbName, environment string) {
	var err error
	var DB *gorm.DB
	var DBURL string

	if environment == "production" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", DbHost, DbPort, DbUser, DbPassword, DbName)
	} else {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	}
	DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to database")
		panic(err.Error())
	} else {
		fmt.Println("We are connected to the database")
	}

	// DB.AutoMigrate(
	// 	&entities.Participant{},
	// )

	routes.InitializeRoutes(DB)
}
