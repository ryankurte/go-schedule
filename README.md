# go-schedule

Database (or other storage) backed scheduling module (WIP).

[![Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/ryankurte/go-schedule)
[![GitHub tag](https://img.shields.io/github/tag/ryankurte/go-schedule.svg)](https://github.com/ryankurte/go-schedule)
[![Build Status](https://travis-ci.org/ryankurte/go-schedule.svg?branch=master)](https://travis-ci.org/ryankurte/go-schedule)

This is designed to allow user or application scheduling of events (one off or a variety of repetitions) that can be maintained in a datastore for consistency / coherency.


## Usage


### Creating the scheduler
```go
    s := NewScheduler(storer Storer, startTime time.Time, tickRate time.Duration)
    go s.Run()
```

### Adding an event
```go
e, err := s.Schedule(baseEvent.Name, baseEvent.Description, baseEvent.When, baseEvent.Repeat)
```

### Subscribing to events
```go
select {
    case e, ok := s.Out:
    if !ok {
        // Scheduler exited / channel closed
    }
    // Do things with event instance
}

```
---

If you have any questions, comments, or suggestions, feel free to open an issue or a pull request.



