package scheduler

import (
	"log"
	"time"

	"github.com/ryankurte/go-schedule/repeat"
)

// Scheduler implements schedule based events
type Scheduler struct {
	storer   Storer
	tickRate time.Duration
	lastTick time.Time
	Out      chan Event
}

const (
	// EventBufferSize is the size of the output event buffer
	EventBufferSize = 1024
)

// NewScheduler creates a scheduler instance using the provided Storer
// startTime is used to filter previous events (ie. startTime of 0 will run and update any pending / not executed events)
// TODO: this will explode / run things as many times as they are scheduled in the past, not sure how to approach this.
// tickRate defines the update rate of the scheduler
func NewScheduler(storer Storer, startTime time.Time, tickRate time.Duration) *Scheduler {
	return &Scheduler{
		storer:   storer,
		tickRate: tickRate,
		lastTick: startTime,
		Out:      make(chan Event, EventBufferSize),
	}
}

// Schedule creates a scheduled event and saves it to the storer
func (s *Scheduler) Schedule(name, description string, when, end time.Time, repeat repeat.Repeat) (Event, error) {
	event, err := s.storer.AddEvent(name, description, when, end, when, repeat)
	return event, err
}

// Run evaluates scheduled tasks every tick
// This should be run as a goroutine
func (s *Scheduler) Run() {
	for {
		select {
		case <-time.After(s.tickRate):
			s.tick(time.Now())
		}
	}
}

func (s *Scheduler) tick(now time.Time) {
	events, err := s.storer.GetEventsFiltered(s.lastTick, now, false)
	if err != nil {
		log.Printf("[SCHEDULER] storage error fetching events (%s)", err)
		return
	}

	for _, event := range events {
		e, updated := s.evaluate(now, event)
		if updated {
			_, err = s.storer.UpdateEvent(e)
			if err != nil {
				log.Printf("[SCHEDULER] storage error updating event: %s (%s)", event.GetID(), err)
			}
		}

	}

	s.lastTick = now
}

func (s *Scheduler) evaluate(now time.Time, event Event) (Event, bool) {
	//log.Printf("Event: %+v", event)

	// Skip if event is completed or not enabled
	if event.IsCompleted() || !event.IsEnabled() {
		return event, false
	}

	// Skip if we are too early
	// ie. not yet scheduled, repeated event but not yet rescheduled
	if event.GetWhen().After(now) || (event.GetRepeat() != repeat.Never && event.GetNextExecution().After(now)) {
		return event, false
	}

	// Skip if after end time
	if !event.GetEnd().IsZero() && now.After(event.GetEnd()) {
		event.SetCompleted(true)
		return event, false
	}

	// Emit event
	s.Out <- event

	// Update run information
	thisRun := event.GetWhen()
	if event.GetNextExecution().After(thisRun) {
		thisRun = event.GetNextExecution()
	}

	// Update last execution and completed if never repeated
	event.SetLastExecution(thisRun)
	if event.GetRepeat() == repeat.Never {
		event.SetCompleted(true)
	}

	// Otherwise, reschedule
	next, err := repeat.Reschedule(thisRun, event.GetRepeat())
	if err != nil {
		log.Printf("[SCHEDULER] error rescheduling: %s", err)
		return event, false
	}
	event.SetNextExecution(next)

	return event, true
}
