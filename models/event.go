package models

import "time"

type Event struct {
	ID int `json:"id"`
	Name int `json:"name"`
	Decription int `json:"description"`
	Location string `json:"location"`
	DateTime time.Time `json:"date_time"`
	UserID int `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event{
	return events
}