package db

import (
	"database/sql" 
	 _"github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB()  {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")

	if err!=nil {
		panic("Database not connected")
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTables()
}

func createTables()  {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)`

	if _, err := Db.Exec(createEventsTable); err!=nil {
		panic("Could not create events table!")
	}
}