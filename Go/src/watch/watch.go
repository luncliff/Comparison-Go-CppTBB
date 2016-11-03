package watch

import (
	"time"
)

type StopWatch struct {
	time.Time
}

func (watch *StopWatch) Pick() time.Duration {
	now := time.Now()
	return now.Sub(watch.Time)
}

func (watch *StopWatch) Reset() time.Duration {
	span := watch.Pick()
	watch.Time = time.Now()
	return span
}
