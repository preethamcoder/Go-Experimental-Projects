package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
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
