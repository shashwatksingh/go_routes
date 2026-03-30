package main

import (
	"fmt"
	"net/http"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func getEvents(context *gin.Context)  {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvents(context *gin.Context)  {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err!=nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse the request"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message" : "Event created", "event": event})
}

func main()  {
	server := gin.Default()

	server.GET("/", homeRoute)
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)


	server.Run(":4002")
}