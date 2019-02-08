package optional

import "time"

var (
	emptyDuration        = Duration{}
	emptyDurationSeconds = DurationSeconds{}
)

// Duration represents optional duration value
type Duration struct {
	duration *time.Duration
}

// OfDuration creates new optional time.Duration containing provided value
func OfDuration(d time.Duration) Duration {
	return Duration{duration: &d}
}

// OfSeconds creates new optional time.Duration containing provided duration in seconds
func OfSeconds(sec int) DurationSeconds {
	d := time.Duration(int64(sec)) * time.Second
	return DurationSeconds{Duration: Duration{duration: &d}}
}

// DurationSeconds is wrapper over Duration to be used in JSON and SQL mappers
type DurationSeconds struct {
	Duration
}
