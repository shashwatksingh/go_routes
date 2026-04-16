package routes

import (
	"net/http"
	"rest_api/models"
	"rest_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	logger := utils.GetLogger()
	events, err := models.GetAllEvents()
	if err != nil {
		logger.WithError(err).Error("Failed to fetch events")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	logger.WithField("count", len(events)).Debug("Events fetched successfully")
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	logger := utils.GetLogger()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		logger.WithError(err).WithField("eventId", context.Param("eventId")).Error("Invalid event ID")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to fetch event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	logger.WithField("eventId", eventId).Debug("Event fetched successfully")
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	logger := utils.GetLogger()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	var updatedEvent models.Event

	if err != nil {
		logger.WithError(err).WithField("eventId", context.Param("eventId")).Error("Invalid event ID")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	existingEvent, err := models.GetEventById(eventId)
	if err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to fetch event for update")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	userId := context.GetInt64("userId")
	if existingEvent.UserID != userId {
		logger.WithField("eventId", eventId).WithField("userId", userId).Warn("Unauthorized update attempt")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		logger.WithError(err).Error("Failed to parse update request")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	updatedEvent.ID = existingEvent.ID
	updatedEvent.UserID = existingEvent.UserID
	if err = updatedEvent.UpdateEvent(); err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to update event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	logger.WithField("eventId", eventId).Info("Event updated successfully")
	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context) {
	logger := utils.GetLogger()
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		logger.WithError(err).WithField("eventId", context.Param("eventId")).Error("Invalid event ID")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to fetch event for deletion")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		logger.WithField("eventId", eventId).WithField("userId", userId).Warn("Unauthorized delete attempt")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	if err = event.DeleteEvent(); err != nil {
		logger.WithError(err).WithField("eventId", eventId).Error("Failed to delete event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	logger.WithField("eventId", eventId).Info("Event deleted successfully")
	context.Status(http.StatusNoContent)
}

func createEvents(context *gin.Context) {
	logger := utils.GetLogger()
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		logger.WithError(err).Error("Failed to parse create event request")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err := event.Save()
	if err != nil {
		logger.WithError(err).WithField("userId", userId).Error("Failed to save event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events"})
		return
	}

	logger.WithField("eventId", event.ID).WithField("userId", userId).Info("Event created successfully")
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
