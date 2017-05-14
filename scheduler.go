package scheduler

import (
	"log"
	"time"
)

// Scheduler implements schedule based events
type Scheduler struct {
	storer   Storer
	tickRate time.Duration
	lastTick time.Time
	out      chan Event
}

const (
	EventBufferSize = 1024
)

func NewScheduler(storer Storer, startTime time.Time, tickRate time.Duration) *Scheduler {
	return &Scheduler{
		storer:   storer,
		tickRate: tickRate,
		lastTick: startTime,
		out:      make(chan Event, EventBufferSize),
	}
}

func (s *Scheduler) Schedule(name, description string, when time.Time, repeat Repeat) (Event, error) {
	event, err := s.storer.AddEvent(name, description, when, when, repeat)
	return event, err
}

func (s *Scheduler) run() {
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
	// Skip if event is completed or not enabled
	if event.IsCompleted() || !event.IsEnabled() {
		return event, false
	}

	// Skip if we are too early
	// ie. not yet scheduled, repeated event but not yet rescheduled
	if event.GetWhen().After(now) || (event.GetRepeat() != RepeatNever && event.GetNextExecution().After(now)) {
		return event, false
	}

	// Emit event
	s.out <- event

	// Update run information
	thisRun := event.GetWhen()
	if event.GetNextExecution().After(thisRun) {
		thisRun = event.GetNextExecution()
	}

	event.SetLastExecution(thisRun)
	if event.GetRepeat() == RepeatNever {
		event.SetCompleted(true)
	}

	next, err := Reschedule(thisRun, event.GetRepeat())
	if err != nil {
		log.Printf("[SCHEDULER] error rescheduling: %s", err)
		return event, false
	}
	event.SetNextExecution(next)

	return event, true
}
