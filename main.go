package main

import (
	"net/http"
	"rest_api/db"
	"rest_api/routes"

	"github.com/gin-gonic/gin"
)

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", homeRoute)
	routes.RegisterRoutes(server)

	server.Run(":4002")
}
