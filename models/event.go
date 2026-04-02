package models

import (
	"rest_api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func (e *Event) DeleteEvent() error {
	query := `
	DELETE FROM events WHERE id=?`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return err
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, date_time, user_id)
	VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id=?`

	row := db.Db.QueryRow(query, id)

	var event Event

	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return Event{}, err
	}

	return event, nil
}

func (e *Event) UpdateEvent() error {
	query := `
	UPDATE events 
	SET name=?, description=?, location=?, date_time=? 
	WHERE id=?`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return err
}

func (e *Event) RegisterForEvent(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id) VALUES(?,?)`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
