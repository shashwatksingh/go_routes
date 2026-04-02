package routes

import (
	"rest_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	
	authenticated.GET("/events", getEvents)
	authenticated.GET("/events/:eventId", getEvent)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
