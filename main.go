package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{"nessage": "You are here!"})
}

func main()  {
	server := gin.Default()

	server.GET("/", getEvents)

	server.Run(":4002")
}