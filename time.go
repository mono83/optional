package optional

import "time"

// Time represents optional time value
type Time struct {
	value    time.Time
	presents bool
}

func (t *Time) set(tt time.Time) {
	t.value = tt
	t.presents = true
}

// OfTime creates new optional time.Time containing provided value
func OfTime(t time.Time) Time {
	return Time{value: t, presents: true}
}

// OfUnixSeconds creates new optional time.Time containing provided unix timestamp in seconds
func OfUnixSeconds(sec int) TimeUnixSeconds {
	ts := time.Unix(int64(sec), 0).UTC()
	return TimeUnixSeconds{Time: Time{value: ts, presents: true}}
}

// OrNow returns value from optional or current time if optional is empty
func (t Time) OrNow() time.Time {
	return t.OrElse(time.Now())
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (t Time) FilterZero() Time {
	if !t.IsPresent() || t.value.IsZero() {
		return Time{}
	}

	return t
}

// TimeUnixSeconds is wrapper over Time to be used in JSON and SQL mappers
type TimeUnixSeconds struct {
	Time
}
