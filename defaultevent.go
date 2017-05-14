package scheduler

import (
	"time"
)

// DefaultEvent implementation provides a standard implementation of the event interface
type DefaultEvent struct {
	ID          string
	Name        string
	Description string
	Enabled     bool
	Completed   bool
	When        time.Time
	Repeat      Repeat
	NextRun     time.Time
	LastRun     time.Time
}

func (de *DefaultEvent) GetID() string                   { return de.ID }
func (de *DefaultEvent) GetName() string                 { return de.Name }
func (de *DefaultEvent) GetDescription() string          { return de.Description }
func (de *DefaultEvent) IsEnabled() bool                 { return de.Enabled }
func (de *DefaultEvent) IsCompleted() bool               { return de.Completed }
func (de *DefaultEvent) SetCompleted(completed bool)     { de.Completed = completed }
func (de *DefaultEvent) GetWhen() time.Time              { return de.When }
func (de *DefaultEvent) GetRepeat() Repeat               { return de.Repeat }
func (de *DefaultEvent) GetLastExecution() time.Time     { return de.LastRun }
func (de *DefaultEvent) SetLastExecution(last time.Time) { de.LastRun = last }
func (de *DefaultEvent) GetNextExecution() time.Time     { return de.NextRun }
func (de *DefaultEvent) SetNextExecution(next time.Time) { de.NextRun = next }
