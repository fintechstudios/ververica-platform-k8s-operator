package scheduling

import "time"

type realClock struct{}

func (*realClock) Now() time.Time { return time.Now() }

func NewClock() Clock {
	return &realClock{}
}

// Clock knows how to get the current time.
// It can be used to fake out timing for testing.
type Clock interface {
	Now() time.Time
}
