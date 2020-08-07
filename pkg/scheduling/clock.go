package scheduling

import "time"

type realClock struct{}

func (*realClock) Now() time.Time { return time.Now() }

// NewClock should always be used to create a NewClock for runtime code
func NewClock() Clock {
	return &realClock{}
}

type fixedClock struct {
	fixedTime *time.Time
}

func (f *fixedClock) Now() time.Time {
	return *f.fixedTime
}

// NewFixedClock creates a Clock that always returns the same time pointer.
// This allows tests to manipulate the current time.
func NewFixedClock(fixedTime *time.Time) Clock {
	if fixedTime == nil {
		panic("fixed time cannot be nil")
	}
	return &fixedClock{fixedTime}
}

// Clock knows how to get the current time.
// It can be used to fake out timing for testing.
type Clock interface {
	Now() time.Time
}
