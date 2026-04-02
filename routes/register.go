package routes

import (
	"net/http"
	"rest_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if err = event.RegisterForEvent(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Registered!"})
}

func cancelRegistration(context *gin.Context) {}
