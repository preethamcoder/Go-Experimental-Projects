package models

import (
	"fmt"
	"time"
)

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

func GetAllEvents() []Event {
	for event := range events {
		fmt.Println(events[event].ID, events[event].Name, events[event].Location)
	}
	return events
}
