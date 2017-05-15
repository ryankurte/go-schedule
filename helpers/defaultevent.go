package helpers

import (
	"time"

	"github.com/ryankurte/go-schedule/repeat"
)

// DefaultEvent implementation provides a standard implementation of the event interface
type DefaultEvent struct {
	ID          string
	Name        string
	Description string
	Enabled     bool
	Completed   bool
	When        time.Time
	End         time.Time
	Repeat      repeat.Repeat
	NextRun     time.Time
	LastRun     time.Time
}

// GetID fetches the event ID
func (de *DefaultEvent) GetID() string { return de.ID }

// GetName fetches the event name
func (de *DefaultEvent) GetName() string { return de.Name }

// GetDescription fetches the event description
func (de *DefaultEvent) GetDescription() string { return de.Description }

// IsEnabled checks if an event is enabled
func (de *DefaultEvent) IsEnabled() bool { return de.Enabled }

// IsCompleted checks if an event is completed
func (de *DefaultEvent) IsCompleted() bool { return de.Completed }

// SetCompleted sets the event completed flag
func (de *DefaultEvent) SetCompleted(completed bool) { de.Completed = completed }

// GetWhen fetches the event time
func (de *DefaultEvent) GetWhen() time.Time { return de.When }

// GetEnd fetches the event end time
func (de *DefaultEvent) GetEnd() time.Time { return de.End }

// GetRepeat fetches the event repeat policy
func (de *DefaultEvent) GetRepeat() repeat.Repeat { return de.Repeat }

// GetLastExecution fetches the last execution time
func (de *DefaultEvent) GetLastExecution() time.Time { return de.LastRun }

// SetLastExecution sets the last execution time
func (de *DefaultEvent) SetLastExecution(last time.Time) { de.LastRun = last }

// GetNextExecution fetches the next execution time
func (de *DefaultEvent) GetNextExecution() time.Time { return de.NextRun }

// SetNextExecution sets the next execution time
func (de *DefaultEvent) SetNextExecution(next time.Time) { de.NextRun = next }
