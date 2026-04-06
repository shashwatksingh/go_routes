package main

import (
	"net/http"
	"rest_api/config"
	"rest_api/db"
	"rest_api/middlewares"
	"rest_api/routes"
	"rest_api/utils"

	"github.com/gin-gonic/gin"
)

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func main() {
	// Load environment variables
	config.LoadDotEnv()
	
	// Initialize logger
	utils.InitLogger()
	logger := utils.GetLogger()
	
	logger.Info("Starting application...")
	
	// Initialize database
	db.InitDB()
	logger.Info("Database initialized")
	
	// Create Gin server without default middleware
	server := gin.New()
	
	// Add recovery middleware
	server.Use(gin.Recovery())
	
	// Add custom logger middleware
	server.Use(middlewares.LoggerMiddleware())
	
	// Register routes
	server.GET("/", homeRoute)
	routes.RegisterRoutes(server)
	
	port := config.GetEnv("PORT")
	logger.WithField("port", port).Info("Server starting")
	
	if err := server.Run(":" + port); err != nil {
		logger.WithError(err).Fatal("Failed to start server")
	}
}
