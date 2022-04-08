package main

import (
	"log"
	"participant-api/app"
	"participant-api/app/helper"
	"participant-api/docs"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	// for load godotenv
	environment := helper.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helper.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	app.Run()
}
