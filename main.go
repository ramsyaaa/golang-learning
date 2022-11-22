package main

import (
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/funding_campaign?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	
	userHandler := handler.NewUserHandler(userService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("/register/email-check", userHandler.ExistingEmail)
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/user/upload-avatar", userHandler.UploadAvatar)

	api.POST("/campaign/create", campaignHandler.CampaignInput)

	router.Run()

	// userInput := user.RegisterUserInput {}
	// userInput.Name = "Test Name"
	// userInput.Email = "email@test.com"
	// userInput.Occupation = "Programmer"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)
}