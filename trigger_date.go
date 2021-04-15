package agscheduler

import "time"

type DateTrigger struct {
	NextRunTime time.Time `json:"next_run_time"`
}

// if previous is empty, return NextRunTime directly
func (t *DateTrigger) GetNextRunTime(previous, now time.Time) time.Time {
	if !previous.Equal(MinDateTime) {
		return MinDateTime
	}
	return t.NextRunTime
}
