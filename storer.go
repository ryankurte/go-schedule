package scheduler

import (
	"time"
)

type Storer interface {
	AddEvent(name, description string, when, next time.Time, string Repeat) (Event, error)
	GetEvent(id string) (Event, error)
	UpdateEvent(event Event) (Event, error)
	GetEventsFiltered(start, end time.Time, completed bool) ([]Event, error)
}
