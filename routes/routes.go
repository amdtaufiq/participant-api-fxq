package routes

import (
	"participant-api/app/controllers"
	"participant-api/app/repositories"
	"participant-api/app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitializeRoutes(DB *gorm.DB) {

	htmlToPdfService := services.HtmlToPDFService()
	participantRepository := repositories.ParticipantRepository(DB)
	participantService := services.ParticipantService(participantRepository)
	participantController := controllers.ParticipantController(participantService, htmlToPdfService)

	//gin config
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	//participant
	api.GET("/participant", participantController.GetParticipants)
	api.GET("/participant/:id", participantController.GetParticipant)
	api.POST("/participant", participantController.CreateParticipant)
	api.PUT("/participant/:id", participantController.UpdateParticipant)
	api.DELETE("/participant/:id", participantController.DeleteParticipant)
	api.GET("/participant/name-tag/:id", participantController.GetPrintNameTag)
	api.GET("/participant/certificate/:id", participantController.GetPrintCertificate)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
