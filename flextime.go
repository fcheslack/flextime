//Package flexdate builds on the standard time pacakge match as much
//of a date as is specified, by attempting to parse starting with the
//longest string
package flextime

import (
	"time"
)

var (
	layouts = []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z07",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02T15:",
		"2006-01-02",
		"2006-01",
		"2006",
	}
)

//SetLayouts changes the accepted layouts used by Parse
func SetLayouts(newLayouts []string) {
	layouts = newLayouts
}

//Parse tries each layout in order until one parses without error.
//If none parse without error it returns the last error from time.Parse
func Parse(value string) (time.Time, error) {
	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t, err
		}
	}
	return t, err
}
