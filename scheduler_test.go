package scheduler

import (
	"testing"
	"time"

	"fmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ryankurte/go-schedule/helpers"
	"github.com/ryankurte/go-schedule/repeat"
)

type MockStorer struct {
	Events []*helpers.DefaultEvent
}

func (ms *MockStorer) AddEvent(name, description string, when, next time.Time, repeat repeat.Repeat) (Event, error) {
	event := helpers.DefaultEvent{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
		Enabled:     true,
		Completed:   false,
		When:        when,
		NextRun:     next,
		Repeat:      repeat,
	}
	ms.Events = append(ms.Events, &event)
	return &event, nil
}

func (ms *MockStorer) GetEvent(id string) (Event, error) {
	for _, e := range ms.Events {
		if e.GetID() == id {
			return e, nil
		}
	}
	return nil, nil
}

func (ms *MockStorer) UpdateEvent(event Event) (Event, error) {
	for i, e := range ms.Events {
		if e.GetID() == event.GetID() {
			ms.Events[i] = event.(*helpers.DefaultEvent)
			return event, nil
		}
	}
	return nil, fmt.Errorf("Event not found matching id: %s", event.GetID())
}

func (ms *MockStorer) GetEventsFiltered(start, end time.Time, getCompleted bool) ([]Event, error) {
	events := make([]Event, 0)

	for _, e := range ms.Events {
		if e.When.After(start) && e.When.Before(end) && (getCompleted || !e.Completed) {
			events = append(events, e)
		}
	}
	return events, nil
}

func TestScheduler(t *testing.T) {
	now := time.Now()

	baseEvent := helpers.DefaultEvent{
		Name:        "Test Event",
		Description: "Test Description",
		When:        now,
		Repeat:      repeat.Never,
		Enabled:     true,
	}

	storer := MockStorer{Events: make([]*helpers.DefaultEvent, 0)}
	scheduler := NewScheduler(&storer, now, time.Second)

	var event Event

	t.Run("Adds events", func(t *testing.T) {
		e, err := scheduler.Schedule(baseEvent.Name, baseEvent.Description, baseEvent.When, baseEvent.Repeat)
		assert.Nil(t, err)
		assert.NotNil(t, e)
		event = e
	})

	t.Run("Events are not executed until scheduled", func(t *testing.T) {
		e, done := scheduler.evaluate(now.Add(-time.Second), event)
		assert.False(t, done)
		assert.NotNil(t, e)
	})

	t.Run("RepeatNever events are executed once", func(t *testing.T) {
		e, done := scheduler.evaluate(now.AddDate(0, 1, 0), event)
		assert.True(t, done)
		assert.NotNil(t, e)

		select {
		case <-time.After(time.Second):
			t.Errorf("Timeout waiting for channel")
		case e := <-scheduler.out:
			assert.EqualValues(t, event.GetID(), e.GetID())
		}

		e, err := scheduler.storer.GetEvent(e.GetID())
		assert.Nil(t, err)
		assert.NotNil(t, e)
		assert.True(t, e.IsCompleted())

		e, done = scheduler.evaluate(now.AddDate(1, 0, 0), event)
		assert.False(t, done)
		assert.NotNil(t, e)
	})

	t.Run("RepeatDaily events are executed daily", func(t *testing.T) {
		dailyEvent := baseEvent
		dailyEvent.Repeat = repeat.Daily
		dailyEvent.When = dailyEvent.When.AddDate(0, 0, 1)

		// Should not run immediately
		e, done := scheduler.evaluate(now.AddDate(0, 0, 0), &dailyEvent)
		assert.False(t, done)
		assert.NotNil(t, e)

		// Should run on start date
		e, done = scheduler.evaluate(now.AddDate(0, 0, 1).Add(time.Second), e)
		assert.True(t, done)
		assert.NotNil(t, e)

		// Should not run again on the same date
		e, done = scheduler.evaluate(now.AddDate(0, 0, 1).Add(time.Second), e)
		assert.False(t, done)
		assert.NotNil(t, e)

		// Should run on next date
		e, done = scheduler.evaluate(now.AddDate(0, 0, 2).Add(time.Second), e)
		assert.True(t, done)
		assert.NotNil(t, e)

		// Should not run again on the same date
		e, done = scheduler.evaluate(now.AddDate(0, 0, 2).Add(time.Second), e)
		assert.False(t, done)
		assert.NotNil(t, e)

	})

}
