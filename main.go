package main

import (
	"fmt"
	"net/http"
	"rest_api/db"
	"rest_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func getEvents(context *gin.Context)  {
	events, err:=models.GetAllEvents()
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context)  {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	event, err:=models.GetEventById(eventId)
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context)  {
	eventId, err :=strconv.ParseInt(context.Param("eventId"), 10, 64)
	var updatedEvent models.Event
	
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	existingEvent, err:=models.GetEventById(eventId)
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	if err := context.ShouldBindJSON(&updatedEvent); err!=nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse the request"})
		return
	}

	updatedEvent.ID = existingEvent.ID
	if err = updatedEvent.UpdateEvent(); err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context)  {
	eventId, err :=strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not parse the id"})
	}

	event, err := models.GetEventById(eventId)
	if err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
	}

	if err = event.DeleteEvent(); err!=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Internal Server Error "})
	}

	context.Status(http.StatusNoContent)
}

func createEvents(context *gin.Context)  {
	var event models.Event
	
	if err := context.ShouldBindJSON(&event); err!=nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse the request"})
		return
	}


	event.UserID=1
	err :=event.Save()
	if err!=nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save events"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Event created", "event": event})
}

func main()  {
	db.InitDB()
	server := gin.Default()

	server.GET("/", homeRoute)
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)

	server.POST("/events", createEvents)


	server.Run(":4002")
}