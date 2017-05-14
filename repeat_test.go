package scheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	now := time.Now()

	t.Run("Never returns current time", func(t *testing.T) {
		next, err := Reschedule(now, RepeatNever)
		assert.Nil(t, err)
		assert.EqualValues(t, now, next)
	})

	t.Run("Reschedules daily", func(t *testing.T) {
		next, err := Reschedule(now, RepeatDaily)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 1), next)
	})

	t.Run("Reschedules weekly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatWeekly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 7), next)
	})

	t.Run("Reschedules biweekly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatBiweekly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 14), next)

	})

	t.Run("Reschedules monthly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatMonthly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 1, 0), next)

	})

	t.Run("Reschedules bimonthly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatBimonthly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 2, 0), next)

	})

	t.Run("Reschedules quarterly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatQuarterly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 3, 0), next)
	})

	t.Run("Reschedules yearly", func(t *testing.T) {
		next, err := Reschedule(now, RepeatYearly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(1, 0, 0), next)
	})

	t.Run("Returns error for invalid repeat types", func(t *testing.T) {
		_, err := Reschedule(now, Repeat("YOLO"))
		assert.NotNil(t, err)
	})
}
