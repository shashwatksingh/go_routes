package utils

import (
	"os"
	"rest_api/config"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger initializes the logger with appropriate settings
func InitLogger() {
	Log = logrus.New()

	// Set output to stdout
	Log.SetOutput(os.Stdout)

	// Set log level based on environment
	env := config.GetEnv("ENV")
	if env == "production" {
		Log.SetLevel(logrus.InfoLevel)
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
	}

	Log.Info("Logger initialized successfully")
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if Log == nil {
		InitLogger()
	}
	return Log
}
