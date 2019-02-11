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

// OfMilliseconds creates new optional time.Duration containing provided duration in milliseconds
func OfMilliseconds(millis int) DurationMillis {
	d := time.Duration(int64(millis)) * time.Millisecond
	return DurationMillis{Duration: Duration{duration: &d}}
}

// OfMinutes creates new optional time.Duration containing provided duration in minutes
func OfMinutes(minutes int) DurationMinutes {
	d := time.Duration(int64(minutes)) * time.Minute
	return DurationMinutes{Duration: Duration{duration: &d}}
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (d Duration) FilterZero() Duration {
	if !d.IsPresent() || d.duration.Nanoseconds() == 0 {
		return emptyDuration
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
