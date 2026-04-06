package db

import (
	"database/sql"
	"rest_api/utils"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	logger := utils.GetLogger()
	var err error
	
	logger.Info("Initializing database connection...")
	Db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
	
	logger.Info("Database connection established successfully")

	createTables()
}

func createTables() {
	logger := utils.GetLogger()
	
	logger.Info("Creating database tables if they don't exist...")
	
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		date_time DATETIME NOT NULL
	)`

	if _, err := Db.Exec(createUsersTable); err != nil {
		logger.WithError(err).Fatal("Failed to create users table")
	}
	logger.Debug("Users table ready")

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	if _, err := Db.Exec(createEventsTable); err != nil {
		logger.WithError(err).Fatal("Failed to create events table")
	}
	logger.Debug("Events table ready")

	registrationTables := `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id)
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	if _, err := Db.Exec(registrationTables); err != nil {
		logger.WithError(err).Fatal("Failed to create registrations table")
	}
	logger.Debug("Registrations table ready")
	
	logger.Info("All database tables created successfully")
}
