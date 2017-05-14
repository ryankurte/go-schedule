/**
 * go-schedule scheduling engine repeat module
 *
 * Copyright 2017 Ryan kurte
 */

package repeat

import (
	"fmt"
	"time"
)

// Repeat defines event repetitions
type Repeat string

const (
	// Never events are scheduled once
	Never Repeat = "never"
	// Daily events are repeated daily at the scheduled time
	Daily Repeat = "daily"
	// Weekly events are repeated weekly (every seven days)
	Weekly Repeat = "weekly"
	// Biweekly repeats every two weeks
	Biweekly Repeat = "biweekly"
	// Monthly events are repeated every month
	Monthly Repeat = "monthly"
	// Bimonthly repeats every two months
	Bimonthly Repeat = "bimonthly"
	// Quarterly events are repeated quarterly (every three months)
	Quarterly Repeat = "quarterly"
	// Yearly events are repeated every year
	Yearly Repeat = "yearly"
)

// Reschedule returns the next execution time for the event given the
// current execution time and the repetition type
func Reschedule(previous time.Time, repeat Repeat) (time.Time, error) {
	switch repeat {
	case Never:
		return previous, nil
	case Daily:
		return previous.AddDate(0, 0, 1), nil
	case Weekly:
		return previous.AddDate(0, 0, 7), nil
	case Biweekly:
		return previous.AddDate(0, 0, 7*2), nil
	case Monthly:
		return previous.AddDate(0, 1, 0), nil
	case Bimonthly:
		return previous.AddDate(0, 2, 0), nil
	case Quarterly:
		return previous.AddDate(0, 3, 0), nil
	case Yearly:
		return previous.AddDate(1, 0, 0), nil
	default:
		return previous, fmt.Errorf("Unsupported repeat value: %s", repeat)
	}
}
