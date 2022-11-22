package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	//var inputEmail user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if(err != nil) { 
		
		errors := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if(err != nil) { 
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Bisa langsung check di Register API
	// emailCheck, err := h.userService.ExistingEmail(inputEmail)

	// if(err != nil) {
	// 	response := helper.APIResponse("Email already used", http.StatusBadRequest, "Error", emailCheck)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	authService := auth.NewService()
	
	token, err := authService.GenerateToken(newUser.ID, newUser.Name, newUser.Role)

	if(err!= nil) {
		response := helper.APIResponse("Generate token failed", http.StatusInternalServerError, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) { 
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if(err!= nil) { 
		errors := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if(err!= nil) { 
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authService := auth.NewService()
	
	token, err := authService.GenerateToken(loggedInUser.ID, loggedInUser.Name, loggedInUser.Role)

	if(err!= nil) {
		response := helper.APIResponse("Generate token failed", http.StatusInternalServerError, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse("Login Success", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ExistingEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if(err != nil) { 
		
		errors := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.ExistingEmail(input)

	if(err != nil) { 
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H {
		"is_available" : isEmailAvailable,
	}

	var metaMessage string

	if(isEmailAvailable) { 
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email is not available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar_file_name")

	if( err != nil ) {
		data := gin.H{"is_uploaded" : false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	path := fmt.Sprintf("upload/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if( err != nil ) {
		data := gin.H{"is_uploaded" : false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if( err != nil ) {
		data := gin.H{"is_uploaded" : false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success to upload avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}