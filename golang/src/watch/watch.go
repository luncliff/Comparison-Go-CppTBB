// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author	: Park Dong Ha ( luncliff@gmail.com )
//
// 	Note	:
//		Simple stop watch to calculate elapsed time
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package watch

import (
	"time"
)

// StopWatch ...
//  	Simple stop watch class
type StopWatch struct {
	time.Time
}

// Pick ...
//		Return elapsed time from start time point
func (watch *StopWatch) Pick() time.Duration {
	now := time.Now()
	return now.Sub(watch.Time)
}

// Reset ...
//		Return elapsed time and reset the start point
func (watch *StopWatch) Reset() time.Duration {
	span := watch.Pick()
	watch.Time = time.Now()
	return span
}
