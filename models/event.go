package models

import "time"

type Event struct {
	ID int
	Name int
	Decription int
	Location string
	DateTime time.Time
	UserID int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}