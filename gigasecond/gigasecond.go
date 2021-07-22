// Package gigasecond implement a method that add a gigasecond on any given time
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond take a time t and add a gigasecond to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
