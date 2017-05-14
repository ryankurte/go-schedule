/**
 * go-schedule scheduling engine repeat module
 *
 * Copyright 2017 Ryan kurte
 */

package scheduler

import (
	"fmt"
	"time"
)

// Repeat defines event repetitions
type Repeat string

const (
	// RepeatNever events are scheduled once
	RepeatNever Repeat = "never"
	// RepeatDaily events are repeated daily at the scheduled time
	RepeatDaily Repeat = "daily"
	// RepeatWeekly events are repeated weekly (every seven days)
	RepeatWeekly Repeat = "weekly"
	// RepeatBiweekly repeats every two weeks
	RepeatBiweekly Repeat = "biweekly"
	// RepeatMonthly events are repeated every month
	RepeatMonthly Repeat = "monthly"
	// RepeatBimonthly repeats every two months
	RepeatBimonthly Repeat = "bimonthly"
	// RepeatQuarterly events are repeated quarterly (every three months)
	RepeatQuarterly Repeat = "quarterly"
	// RepeatYearly events are repeated every year
	RepeatYearly Repeat = "yearly"
)

// Reschedule returns the next execution time for the event given the
// current execution time and the repetition type
func Reschedule(previous time.Time, repeat Repeat) (time.Time, error) {
	switch repeat {
	case RepeatNever:
		return previous, nil
	case RepeatDaily:
		return previous.AddDate(0, 0, 1), nil
	case RepeatWeekly:
		return previous.AddDate(0, 0, 7), nil
	case RepeatBiweekly:
		return previous.AddDate(0, 0, 7*2), nil
	case RepeatMonthly:
		return previous.AddDate(0, 1, 0), nil
	case RepeatBimonthly:
		return previous.AddDate(0, 2, 0), nil
	case RepeatQuarterly:
		return previous.AddDate(0, 3, 0), nil
	case RepeatYearly:
		return previous.AddDate(1, 0, 0), nil
	default:
		return previous, fmt.Errorf("Unsupported repeat value: %s", repeat)
	}
}
