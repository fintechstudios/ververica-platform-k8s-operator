package utils

import "time"

// MustParseTime tries to parse a string into a time.Time and panics if there is an error
func MustParseTime(layout, str string) (t time.Time) {
	var err error
	if t, err = time.Parse(layout, str); err != nil {
		panic(err)
	}
	return t
}
