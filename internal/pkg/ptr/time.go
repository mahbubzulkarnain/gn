package ptr

import "time"

// Time returns a pointer to the given time.Time
func Time(i time.Time) *time.Time {
	return &i
}
