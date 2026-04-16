package routes

import (
	"net/http"
	"rest_api/models"
	"rest_api/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	logger := utils.GetLogger()
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		logger.WithError(err).Error("Failed to parse signup request")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err := user.Save()
	if err != nil {
		logger.WithError(err).WithField("email", user.Email).Error("Failed to save user")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	logger.WithField("email", user.Email).Info("User created successfully")
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func login(context *gin.Context) {
	logger := utils.GetLogger()
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		logger.WithError(err).Error("Failed to parse login request")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		logger.WithField("email", user.Email).Warn("Failed login attempt - invalid credentials")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthenticated!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		logger.WithError(err).WithField("email", user.Email).Error("Failed to generate token")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	logger.WithField("email", user.Email).Info("User logged in successfully")
	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})
}
