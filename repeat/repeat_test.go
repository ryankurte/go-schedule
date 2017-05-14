package repeat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	now := time.Now()

	t.Run("Never returns current time", func(t *testing.T) {
		next, err := Reschedule(now, Never)
		assert.Nil(t, err)
		assert.EqualValues(t, now, next)
	})

	t.Run("Reschedules daily", func(t *testing.T) {
		next, err := Reschedule(now, Daily)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 1), next)
	})

	t.Run("Reschedules weekly", func(t *testing.T) {
		next, err := Reschedule(now, Weekly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 7), next)
	})

	t.Run("Reschedules biweekly", func(t *testing.T) {
		next, err := Reschedule(now, Biweekly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 0, 14), next)

	})

	t.Run("Reschedules monthly", func(t *testing.T) {
		next, err := Reschedule(now, Monthly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 1, 0), next)

	})

	t.Run("Reschedules bimonthly", func(t *testing.T) {
		next, err := Reschedule(now, Bimonthly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 2, 0), next)

	})

	t.Run("Reschedules quarterly", func(t *testing.T) {
		next, err := Reschedule(now, Quarterly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(0, 3, 0), next)
	})

	t.Run("Reschedules yearly", func(t *testing.T) {
		next, err := Reschedule(now, Yearly)
		assert.Nil(t, err)
		assert.EqualValues(t, now.AddDate(1, 0, 0), next)
	})

	t.Run("Returns error for invalid repeat types", func(t *testing.T) {
		_, err := Reschedule(now, Repeat("YOLO"))
		assert.NotNil(t, err)
	})
}
