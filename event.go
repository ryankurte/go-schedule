package scheduler

import (
	"time"

	"github.com/ryankurte/go-schedule/repeat"
)

// Event defines the interface that must be implemented by schedulable objects
type Event interface {
	GetID() string
	GetName() string
	GetDescription() string
	IsEnabled() bool
	IsCompleted() bool
	SetCompleted(bool)
	GetWhen() time.Time
	GetEnd() time.Time
	GetRepeat() repeat.Repeat
	GetLastExecution() time.Time
	SetLastExecution(time.Time)
	GetNextExecution() time.Time
	SetNextExecution(time.Time)
}
