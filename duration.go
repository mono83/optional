package optional

import "time"

// Duration represents optional duration value
type Duration struct {
	value    time.Duration
	presents bool
}

func (d *Duration) set(dd time.Duration) {
	d.value = dd
	d.presents = true
}

// OfDuration creates new optional time.Duration containing provided value
func OfDuration(d time.Duration) Duration {
	return Duration{value: d, presents: true}
}

// OfDurationRef creates new optional time.Duration containing provided value
func OfDurationRef(d *time.Duration) Duration {
	if d == nil {
		return Duration{}
	}

	return OfDuration(*d)
}

// OfSeconds creates new optional time.Duration containing provided duration in seconds
func OfSeconds(sec int) DurationSeconds {
	d := time.Duration(int64(sec)) * time.Second
	return DurationSeconds{Duration: Duration{value: d, presents: true}}
}

// OfSecondsRef creates new optional time.Duration containing provided duration in seconds
func OfSecondsRef(sec *int) DurationSeconds {
	if sec == nil {
		return DurationSeconds{}
	}

	return OfSeconds(*sec)
}

// OfMilliseconds creates new optional time.Duration containing provided duration in milliseconds
func OfMilliseconds(millis int) DurationMillis {
	d := time.Duration(int64(millis)) * time.Millisecond
	return DurationMillis{Duration: Duration{value: d, presents: true}}
}

// OfMillisecondsRef creates new optional time.Duration containing provided duration in milliseconds
func OfMillisecondsRef(millis *int) DurationMillis {
	if millis == nil {
		return DurationMillis{}
	}

	return OfMilliseconds(*millis)
}

// OfMinutes creates new optional time.Duration containing provided duration in minutes
func OfMinutes(minutes int) DurationMinutes {
	d := time.Duration(int64(minutes)) * time.Minute
	return DurationMinutes{Duration: Duration{value: d, presents: true}}
}

// OfMinutesRef creates new optional time.Duration containing provided duration in minutes
func OfMinutesRef(minutes *int) DurationMinutes {
	if minutes == nil {
		return DurationMinutes{}
	}

	return OfMinutes(*minutes)
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (d Duration) FilterZero() Duration {
	if !d.IsPresent() || d.value.Nanoseconds() == 0 {
		return Duration{}
	}

	return d
}

// DurationSeconds is wrapper over Duration to be used in JSON and SQL mappers
type DurationSeconds struct {
	Duration
}

// DurationMillis is wrapper over Duration to be used in JSON and SQL mappers
type DurationMillis struct {
	Duration
}

// DurationMinutes is wrapper over Duration to be used in JSON and SQL mappers
type DurationMinutes struct {
	Duration
}
