package routes

import (
	"net/http"
	"rest_api/models"
	"rest_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	logger := utils.GetLogger()
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		logger.WithError(err).WithField("eventId", context.Param("eventId")).Error("Invalid event ID for registration")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to fetch event for registration")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if err = event.RegisterForEvent(userId); err != nil {
		logger.WithError(err).WithField("eventId", eventId).WithField("userId", userId).Error("Failed to register user for event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	logger.WithField("eventId", eventId).WithField("userId", userId).Info("User registered for event successfully")
	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	logger := utils.GetLogger()
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		logger.WithError(err).WithField("eventId", context.Param("eventId")).Error("Invalid event ID for cancellation")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		logger.WithError(err).WithField("eventId", eventId).WithField("userId", userId).Error("Failed to cancel registration")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for the user"})
		return
	}

	logger.WithField("eventId", eventId).WithField("userId", userId).Info("Registration cancelled successfully")
	context.Status(http.StatusNoContent)
}
