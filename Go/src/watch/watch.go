// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File 	: watch.go
//  Author 	: Park  Dong Ha ( luncliff@gmail.com )
//  Updated : 2016/12/17
//
//  Note 	:
//		Simple stop watch to calculate elapsed time
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package watch

import (
	"time"
)

type StopWatch struct {
	time.Time
}

// Pick ...
//		Return elapsed time from start time point
func (watch *StopWatch) Pick() time.Duration {
	now := time.Now()
	return now.Sub(watch.Time)
}

// Pick ...
//		Return elapsed time and reset the start point
func (watch *StopWatch) Reset() time.Duration {
	span := watch.Pick()
	watch.Time = time.Now()
	return span
}
