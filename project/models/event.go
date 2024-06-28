package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	Host        int
}

var events []Event = []Event{}

func (e Event) Save() {
	events = append(events, e)
}
